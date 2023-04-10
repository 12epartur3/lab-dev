#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

#include <atomic>
#include <iostream>

#include "httplib.h"

const int kGuessTime = 600; // guess times
const int kThreadNum = 20;  // concurrency number
const float kDiff = 8;      // to discard result
const int kSleepTimeMs = 10 * 1000;  // wait interval
const int kStartTimeMs = 20;  // start to get Gaussian
const int kEndTimeMs = 650;   // stop to get Gaussian and prepare to commit answer
const int64_t kOffsetTimeMs = 210; // clock offset with server

const std::string kIp = "59.110.159.71";
const int kPort = 16541;

static pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;
static pthread_cond_t cond = PTHREAD_COND_INITIALIZER;
static bool run;
static std::atomic<int> stopNum;

// use for time cost count
class TimeCost {
 public:
  TimeCost() { gettimeofday(&tm_start_, NULL); }

  int64_t GetElapsedUs() const {
    struct timeval tm_end;
    gettimeofday(&tm_end, NULL);
    return (tm_end.tv_sec - tm_start_.tv_sec) * 1000 * 1000 +
           (tm_end.tv_usec - tm_start_.tv_usec);
  }

  int64_t GetElapsedMs() const { return GetElapsedUs() / 1000; }

  void Reset() { gettimeofday(&tm_start_, NULL); }

  static int64_t CurrentUs() {
    struct timeval tm_end;
    gettimeofday(&tm_end, NULL);
    return tm_end.tv_sec * 1000 * 1000 + tm_end.tv_usec;
  }

  static int64_t CurrentMs() {
    struct timeval tm_end;
    gettimeofday(&tm_end, NULL);
    return tm_end.tv_sec * 1000 + tm_end.tv_usec / 1000;
  }

  static int64_t CurrentMsClock() { return CurrentMs() % 1000; }

  static int64_t CurrentMsClockWithOffset(int64_t offset) {
    return (CurrentMs() + offset) % 1000;
  }

 private:
  struct timeval tm_start_;
};

// send request and commit answer
class GuessMean {
 public:
  explicit GuessMean(httplib::Client& cli) : cli_(&cli) {}

  bool getGaussian(float& r) {
    // TimeCost tc;
    // std::cout << __func__ << " curMsClock = " << tc.CurrentMsClockWithOffset(kOffsetTimeMs)
    //           << '\n';
    auto res = cli_->Get("/");
    if (!res) {
      std::cout << __func__ << " error code: " << res.error() << '\n';
      return false;
    }
    if (res->status != 200) {
      std::cout << " error http code = " << res->status << '\n';
      return false;
    }
    // std::cout << __func__ << " cost = " << tc.GetElapsedMs() << "ms\n";

    // std::cout << "body = " << res->body << '\n';
    // std::string mean = res->body;
    // std::cout << __func__ << " mean = " << mean
    //           << " curMsClock = " << tc.CurrentMsClockWithOffset(kOffsetTimeMs)
    //           << '\n';
    r = std::stof(res->body);
    return true;
  }

  bool guess(const float& g) {
    // TimeCost tc;
    auto res = cli_->Get("/submit?guess=" + std::to_string(g));
    if (!res) {
      std::cout << __func__ << " error code: " << res.error() << '\n';
      return false;
    }
    if (res->status != 200) {
      std::cout << " error http code = " << res->status << '\n';
      return false;
    }
    // std::cout << __func__ << " cost = " << tc.GetElapsedMs() << "ms\n";
    // std::cout << __func__ << " body = " << res->body << '\n';
    std::cout << res->body << '\n';
    return true;
  }

private:
  httplib::Client* cli_;
};

struct tls {
  httplib::Client* c;
  float sum;
  int num;
};

void* runner(void* argv) {
  tls* t = static_cast<tls*>(argv);
  GuessMean gm(*(t->c));
  TimeCost tc;
  float prev = 0;
  while (1) {
    pthread_mutex_lock(&mtx);
    while (run == false) {
      ++stopNum;
      prev = 0;
      pthread_cond_wait(&cond, &mtx); // wait to start
    }
    pthread_mutex_unlock(&mtx);
    // printf("i am thread %ld, run at %ldms, seq = %d\n", pthread_self(),
    //        tc.CurrentMsClockWithOffset(kOffsetTimeMs), t->num);
    float v;
    bool ret = gm.getGaussian(v);
    if (!ret) continue;
    float diff = abs(v - prev);
    if (prev != 0 && diff >= kDiff) {
      printf("i am thread %ld, run at %ldms, seq = %d, diff = %f\n", pthread_self(),
             tc.CurrentMsClockWithOffset(kOffsetTimeMs), t->num, diff);
    }
    prev = v;
    t->sum = t->sum + v;
    t->num++;
  }
  return NULL;
}

int main(void) {
  httplib::Client cli(kIp, kPort);
  cli.set_connection_timeout(0, 500000);  // 500 milliseconds
  cli.set_read_timeout(0, 500000);        // 500 milliseconds
  cli.set_write_timeout(0, 500000);       // 500 milliseconds
  cli.set_keep_alive(true);
  GuessMean gm(cli);

  run = false;
  pthread_t t_id[kThreadNum];
  std::vector<tls*> t;
  t.resize(kThreadNum);
  TimeCost tc;
  printf("start %ldms\n", tc.CurrentMs());
  
  for (int i = 0; i < kThreadNum; i++) {
    t[i] = new tls();
    t[i]->c = new httplib::Client(kIp, kPort);
    t[i]->c->set_connection_timeout(0, 500000);  // 500 milliseconds
    t[i]->c->set_read_timeout(0, 500000);        // 500 milliseconds
    t[i]->c->set_write_timeout(0, 500000);       // 500 milliseconds
    t[i]->c->set_keep_alive(true);
    t[i]->sum = 0;
    t[i]->num = 0;
  }


  for (int i = 0; i < kThreadNum; i++) {
    pthread_create(&t_id[i], NULL, runner, static_cast<void*>(t[i]));
  }
  sleep(1); // wait thread init
  stopNum.store(0);
  int curMsClock;

  int guessCount = 0;
  while (1) {
    curMsClock = tc.CurrentMsClockWithOffset(kOffsetTimeMs);
    // printf("main thread %ld, curMsClock = %d\n", pthread_self(), curMsClock);
    if (curMsClock <= kStartTimeMs && !run) {  // start getGaussian
      printf("main thread %ld, curMsClock = %d, start getGaussian\n", pthread_self(),
             curMsClock);
      pthread_mutex_lock(&mtx);
      run = !run;
      pthread_cond_broadcast(&cond);
      pthread_mutex_unlock(&mtx);
    } else if (curMsClock >= kEndTimeMs && run) {  // stop getGaussian
      printf("main thread %ld, curMsClock = %d, stop getGaussian\n", pthread_self(),
             curMsClock);
      guessCount++;
      pthread_mutex_lock(&mtx);
      run = !run;
      pthread_mutex_unlock(&mtx);
      while (stopNum.load() < kThreadNum) {
        usleep(kSleepTimeMs / 2);
        // printf("main thread %ld, curMsClock = %d, wait, stopNum = %d\n",
        //        pthread_self(), curMsClock, stopNum.load());
      }
      stopNum.store(0);
      float sum = 0;
      int num = 0;
      float guess = 0;
      for (int i = 0; i < t.size(); i++) {
        sum = sum + t[i]->sum;
        num = num + t[i]->num;
        t[i]->sum = 0;
        t[i]->num = 0;
      }
      if (num == 0) {
        continue;
      }
      guess = sum / num;
      printf("main thread %ld, curMsClock = %d, getGaussian num = %d, sum = %f, guess = %f\n",
          pthread_self(), curMsClock, num, sum, guess);
    
      TimeCost tc;
      gm.guess(guess);  // commit answer
      curMsClock = tc.CurrentMsClockWithOffset(kOffsetTimeMs);
      printf("main thread %ld, curMsClock = %d, guess cost = %ldms\n",
             pthread_self(), curMsClock, tc.GetElapsedMs());
    }
    if (guessCount >= kGuessTime) break;
    fflush(stdout);
    usleep(kSleepTimeMs);
  }

  // release resource
  for (int i = 0; i < t.size(); i++) {
    delete t[i]->c;
    delete t[i];
  }
  return 0;
}

