#include<iostream>
using namespace std;
void ConstTest1() {
	const volatile int a = 5;
	int *p;
	p = const_cast<int*>(&a);
	(*p)++;
	cout << a << endl;
	cout << *p << endl;
	cout << *(&a) << endl;
}
int main()
{
	ConstTest1();
	double a = 9.12345999;
	//int b = static_cast<int>(a);
	cout<< "b = "<< a<< '\n';
	return 0;
}
