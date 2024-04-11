#include <atomic>
#include <iostream>

void callAdd() {
	static std::atomic<uint64_t> count = 0;
	std::cout <<"count = " << count++ << "\n";
}
int main() {
	for (int i = 0; i < 10; i++) callAdd();
}
