#include <iostream>
#include <memory>


int main() {
	int a = 10;
	const int* p = &a;
	std::shared_ptr<const int> s_p = std::make_shared<int>(10);
	//*p = 11;
	//*s_p = 11;
	int* const p1 = &a;
	const std::shared_ptr<int> s_p1 = std::make_shared<int>(10);
	int b = 10;
	//p1 = &b;
	//s_p1 = std::make_shared<int>(11);
}
