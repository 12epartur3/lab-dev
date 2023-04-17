#include <iostream>
#include <atomic>
#include <thread>
#include <condition_variable>
#include <mutex>
#include <random>

std::mutex m;

class WaitGroup {
    public:
	WaitGroup(int n) : nleft_(n) {}
	void Done() const {
		nleft_--;
		c_.notify_one();
	}
	void Wait() const {
		std::unique_lock<std::mutex> l(m_);
		while (nleft_ > 0) {
			c_.wait(l);	
			std::cout << "wait weak up nleft = " << nleft_ << "\n";
		}
		//c_.wait(m_, [this]() {return this->nleft_ <= 0;});
	}
    private:
	mutable std::mutex m_;
	mutable std::condition_variable c_;
	mutable std::atomic<int> nleft_;
};

void runner(const WaitGroup& wg, int s) {
	m.lock();
	std::cout << "i am thread " << std::this_thread::get_id() << ", sleep for " << s << " seconds\n" << std::flush;
	m.unlock();
	
	std::this_thread::sleep_for(std::chrono::seconds(s));
	wg.Done();
	m.lock();
	std::cout << "i am thread " << std::this_thread::get_id() << ", sleep end\n" << std::flush;
	m.unlock();
}
int main() {
	std::cout << "start\n";
	int N = 5;
	WaitGroup wg(N);
	std::default_random_engine generator;
	std::uniform_int_distribution<int> distribution(3, 9);
	for (int i = 0; i < N; i++) {
		int dice_roll = distribution(generator);
		std::thread t(runner, std::ref(wg), dice_roll);
		t.detach();
	}
	std::cout << "wait\n";
	auto start = std::chrono::system_clock::now();
	wg.Wait();
	auto end = std::chrono::system_clock::now();
	std::chrono::duration<double> elapsed_seconds = end - start;
	std::cout << "wait end, for " << elapsed_seconds.count() << " s\n";
}
