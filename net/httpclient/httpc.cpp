#include <iostream>
#include <sys/time.h>
#include <stdio.h>

#include "httplib.h"

const int N = 3;
const std::string ip = "59.110.159.71";
const int port = 16543;

static pthread_mutex_t mtx = PTHREAD_MUTEX_INITIALIZER;
static pthread_cond_t cond = PTHREAD_COND_INITIALIZER;
static bool run;

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

 private:
  struct timeval tm_start_;
};

class GuessMean {
public:
  GuessMean(httplib::Client& cli) : cli_(&cli) {}
  bool getGaussian(float& r) {
    TimeCost tc;
    auto res = cli_->Get("/");
    if (!res) {
      std::cout << __func__ << "error code: " << res.error() << '\n';
      return false;
    }
    if (res->status != 200) {
      std::cout << "error http code = " << res->status << '\n';
      return false;
    }
    std::cout << __func__ << " cost = " << tc.GetElapsedUs() << "us\n";

    // std::cout << "body = " << res->body << '\n';
    std::string mean = res->body;
    std::cout << __func__ << " mean = " << mean << '\n';
    r = std::stof(mean);
    return true;
  }

  bool guess(const float& g) {
    TimeCost tc;
    auto res = cli_->Get("/submit?guess=" + std::to_string(g));
    if (!res) {
      std::cout << __func__ << "error code: " << res.error() << '\n';
      return false;
    }
    if (res->status != 200) {
      std::cout << "error http code = " << res->status << '\n';
      return false;
    }
    std::cout << __func__ << " cost = " << tc.GetElapsedUs() << "us\n";
    std::cout << __func__ << " body = " << res->body << '\n';
    return true;
  }

private:
  httplib::Client* cli_;
};

void* runner(void* argv) {
  httplib::Client* cli = static_cast<httplib::Client*>(argv);
  //httplib::Client cli(ip, port);
  GuessMean gm(*cli);
  std::pair<bool, float>* p = new std::pair<bool, float>(false, 0);
  TimeCost tc;
  // pthread_mutex_lock(&mtx);
  while (run == false) {  
    pthread_cond_wait(&cond, &mtx);
  }
  pthread_mutex_unlock(&mtx);
  printf("i am thread %ld, run at %ld\n", pthread_self(), tc.CurrentUs());
  p->first = gm.getGaussian(p->second);
  //printf("i am thread %ld, ok = %d, v = %f\n", pthread_self(), p->first, p->second);
  return p;
}

float mean(const std::vector<std::pair<bool, float>*>& r) {
  float sum = 0;
  float count = 0;
  for (auto p : r) {
    if (p->first == false) continue;
    count++;
    sum = sum + p->second;
  }
  return sum / count;
}


int main(void) {
  httplib::Client cli(ip, port);
  GuessMean gm(cli);


  pthread_t t_id[N];
  std::vector<httplib::Client*> c;
  std::vector<std::pair<bool, float>*> r;
  c.resize(N);
  r.resize(N);
  TimeCost tc, tcAll;
  printf("start %ldms\n", tc.CurrentMs());
  
  for (int i = 0; i < N; i++) {
    c[i] = new httplib::Client(ip, port);
  }

  for (int i = 0; i < N; i++) {
    pthread_mutex_lock(&mtx);
    pthread_create(&t_id[i], NULL, runner, static_cast<void*>(c[i]));
  }
  pthread_mutex_lock(&mtx);
  sleep(1);
  run = true;
  pthread_cond_broadcast(&cond);
  pthread_mutex_unlock(&mtx);

  printf("cost create %ldus\n", tc.GetElapsedUs());
  tc.Reset();
  void* ret = NULL;
  for (int i = 0; i < N; i++) {
    ret = NULL;
    pthread_join(t_id[i], &ret);
    if (!ret) {
      continue;
    }
    std::pair<bool, float>* p = static_cast<std::pair<bool, float>*>(ret);
    printf("thread %ld, ok = %d, v = %f\n", t_id[i], p->first, p->second);
    r[i] = p;
  }
  printf("cost runner %ldus\n", tc.GetElapsedUs());
  tc.Reset();
  float m = mean(r);
  printf("mean = %f\n", m);
  printf("cost mean %ldus\n", tc.GetElapsedUs());
  tc.Reset();
  gm.guess(m);
  printf("cost guess %ldus\n", tc.GetElapsedUs());
  printf("end %ldms, cost = %ldus\n", tc.CurrentMs(), tcAll.GetElapsedUs());
}
