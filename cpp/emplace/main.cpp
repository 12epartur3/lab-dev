#include <vector>
#include <string>
#include <iostream>


int main() {
	std::string s = "1234";
	std::vector<std::string> v;
	std::cout <<"s = " << s <<'\n';
	v.emplace_back(s);
	std::cout <<"s = " << s <<'\n';
}
