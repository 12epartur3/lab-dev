#include <thread>
#include <atomic>
#include <cassert>
#include <string>
#include <iostream>
using namespace std;

std::atomic<std::string*> ptr;
int data;

void producer()
{
    std::string* p  = new std::string("Hello");
    data = 42;
    //ptr.store(p, std::memory_order_release);
    ptr.store(p, std::memory_order_relaxed);
}

void consumer()
{
    std::string* p2;
    while (!(p2 = ptr.load(std::memory_order_relaxed)));
    assert(data == 42); // never fires
    assert(*p2 == "Hello"); // never fires
}

int main()
{
    while (1) {
            std::thread t2(consumer);
            std::thread t1(producer);
            t1.join(); t2.join();
	    //std::cout << "\n";
    }
}
