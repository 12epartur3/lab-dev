#include <iostream>


class A
{
	public:
		A(int& value):_value(value)
		{
			//_value = value;
		}
		int& _value;

};


int main()
{
	int value = 5;
	A a(value);
	std::cout<< "value = " << a._value << '\n';

}
