#include <vector>
#include <iostream>

std::vector<int>& return_vector(void)
{
    std::vector<int> tmp {1,2,3,4,5};
    //return std::move(tmp);
    return tmp;
}

int main()
{
	std::vector<int>& rval_ref = return_vector();
	for(int& item : rval_ref)
	{
		std::cout<< item << '\n';
	}
	
}
