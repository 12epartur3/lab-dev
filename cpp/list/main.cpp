#include <list>
#include <iostream>

int main()
{
	std::list<int> my_list = {1, 2, 3, 4};
	std::list<int>::iterator it = std::begin(my_list);
	do
	{
		std::cout<< *it << '\n';
		++it;
	}
	while(it != std::end(my_list));
	it = std::begin(my_list);
	my_list.erase(it);
	std::cout<< *it << '\n';
	
	
}
