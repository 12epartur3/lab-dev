#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>

int main() {
        pid_t p = fork();
        if (p == 0) {
                printf("i am child\n");
        } else {
		printf("i am parent, child pid = %d\n", p);
	}
}
