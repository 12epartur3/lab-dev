#include <stdio.h>
#include <iostream>

struct flexible_t{
    int a;
    //double b;
    char c[0];
}; 

int main() {
	flexible_t* f = (flexible_t*)malloc(sizeof(flexible_t) + (100) * sizeof(char));
	std::cout << sizeof(*f) << '\n';
	f->a = 10;
	f->c[0] = '1';
	//flexible_t f;
	//f.a = 10;
	std::cout << "f.a = " << f->a <<'\n';
	//f.c[0] = '1';
	std::cout << "f.c[0] = " << f->c[0] <<'\n';
	std::cout << sizeof(*f) << '\n';
}
