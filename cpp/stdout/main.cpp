#include <iostream>
#include <unistd.h>
#include <stdio.h>


int main() {
	for(int i = 0; i < 9999; i++) {
		sleep(1);
		//std::cout<<"cout " << i << std::endl;
		std::cout<<"cout " << i << '\n';
		printf("printf %d\n", i);
		fflush(stdout);
		fprintf(stdout, "stdout fprintf %d\n", i);
		//stderr no buffer
		fprintf(stderr, "stderr fprintf %d\n", i);

	}

}
