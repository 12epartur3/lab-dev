#include <iostream>
#include <exception>


struct Item {
        Item() {
                std::cout << "Item constructor call\n";
        }
        virtual ~Item() {
                std::cout << "Item destructor call\n";
        }
        void* operator new(std::size_t size) {
                std::cout << "Item operator new call, size = " << size << "\n";
                return malloc(size);

        }
        void operator delete(void* ptr) {
                std::cout << "Item operator delete call\n";
                free(ptr);
        }
};

void foo() {
	std::exception e;
	throw e;
	std::cout << "end foo\n";
}
void mumble() {
	Item* it_ptr;
	Item it;
	//try {
		foo();
	//} catch(std::exception& e) {
	//	std::cout <<"exception is " << e.what() << '\n';
	//}
	std::cout << "end mumble\n";
}
int main() {
	mumble();
	try {
	} catch(std::exception& e) {
		std::cout <<"exception is " << e.what() << '\n';
	}
	std::cout << "end main\n";
}
