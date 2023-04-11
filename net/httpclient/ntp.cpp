#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

#include <iostream>

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

int main() {
  TimeCost tc;
  while (1) {
    std::cout << "current ms = " << tc.CurrentMs() << "\nclock ms = " << tc.CurrentMsClock() << "\n";

    usleep(20 * 1000);
  }
}