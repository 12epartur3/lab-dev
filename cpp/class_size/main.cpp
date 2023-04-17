#include <iostream>


class X {
	void virtual func() {}
};


class Y : virtual public X {
};

class Z : virtual public X {
};

class A : public Y, Z {
};


int main() {
	std::cout << "sizeof(X) = " << sizeof(X) << "\n";
	std::cout << "sizeof(Y) = " << sizeof(Y) << "\n";
	std::cout << "sizeof(Z) = " << sizeof(Z) << "\n";
	std::cout << "sizeof(A) = " << sizeof(A) << "\n";
}
