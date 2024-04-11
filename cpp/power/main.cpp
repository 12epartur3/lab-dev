#include <iostream>

int nPowerOf2(uint64_t integer) {
    int n = 0;
    while (integer > 1) {
        integer >>= 1;
        n++;
    }
    return n;
}

int main() {
	std::cout << "nPowerOf2(8) = " << nPowerOf2(8) << "\n";
}
