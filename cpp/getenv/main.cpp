#include <stdlib.h>
#include <string>
#include <iostream>
int main()
{
	const char* pfb = std::getenv("PFB_NAME");
	std::string data;
	//if(pfb != NULL)data = pfb;
        data = pfb;
	std::cout << "PFB_NAME = " << data << '\n';

}
