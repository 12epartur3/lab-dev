#include <iostream>
#include <stdio.h>
using namespace std;

class A {
public:
	virtual void f1() {
		cout << "A::f1()" << endl;
	}
	virtual void f2() {
		cout << "A::f2()" << endl;
	}
	void f3() {
		cout << "A::f2()" << endl;
	}
	int a;
};

class B : public A {
public:
	virtual void f1() {
		cout << "B::f1()" << endl;
	}
	virtual void f2() {
		cout << "B::f2()" << endl;
	}
	void f3() {
		cout << "B::f2()" << endl;
	}
	int b;
};

class C : public B {
public:
	//virtual void f1() {
	//	cout << "C::f1()" << endl;
	//}
	//virtual void ff1() {
	//	cout << "C::ff1()" << endl;
	//}
	//virtual void ff2() {
	//	cout << "C::ff2()" << endl;
	//}
	void f3() {
		cout << "B::f2()" << endl;
	}
	int c;
};
typedef void(*VTable)();

void print(A& a) {
	int64_t vptr = *(int64_t*)&a;
	printf("vptr = %p\n", vptr);
	VTable vtb = (VTable)*(int64_t*)*(int64_t*)&a;
	int  i = 0;
	//while (vtb != NULL) {
	while (i < 2) {
		cout << "Function:" << ++i << endl;
		cout << "vtb = " << *vtb << endl;
		cout << "------->";
		vtb();
		vtb = (VTable)*((int64_t*)*(int64_t*)&a + i);
	}
}

int main(int argc,char* argv[])
{
	A a;
	B b;
	C c;
	//print(a);
	print(b);
	print(c);
}

