#include <iostream>

class A {
  public:
    A() {
      std::cout << "[C] constructor fired." << std::endl;
    }

    A(const A &a) {
      std::cout << "[C] copying constructor fired." << std::endl;
      x = a.x;
    }

    A(A &&a) {
      std::cout << "[C] moving copying constructor fired." << std::endl;
      x = a.x;
    }

    ~A() {
      std::cout << "[C] destructor fired." << std::endl;
    }
    int x;
};

A getTempA() {
  A a;
  a.x = 99999;
  //return a;
  return std::move(a);
}
 
int main(int argc, char **argv) {
  A a = getTempA();
  std::cout<< "x = "<< a.x<< '\n';
  return 0;
}
