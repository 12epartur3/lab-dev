#include <iostream>
#include <string>

struct Test {
	Test() {
		std::cout << "constructor call\n";
	}
	virtual ~Test() {
		std::cout << "destructor call\n";
	}
	int t;
};

struct BigTest: public Test {

	BigTest() {
		std::cout << "big constructor call\n";
	}
	virtual ~BigTest() {
		std::cout << "big destructor call\n";
	}
	int b;
};

Test func() {
	//Test t;
	return Test();
}
int main() {
	//Test *p1 = new Test[5];
	//delete p1;
	//Test *p2 = new BigTest[3];
	//std::cout << "delete\n";
	//delete [] p2;
	//Test t = func();
	std::cout << "end main\n";
	
}
