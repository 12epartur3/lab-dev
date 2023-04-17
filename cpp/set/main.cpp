#include <set>
#include <iostream>

void travel_set(const std::set<int>& set)
{

	for(const int& x : set)
	{
		std::cout<< "set = " << x << '\n';
	}
}
int main()
{
	std::set<int> myset;
	travel_set(myset);
	
}
