#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>

int main(int argc, char *argv[], char *envp[]) {
	printf("hello, i am exec, pid = %d\n", getpid());
	int i = 0;
	char *p = argv[i];
	while (p != NULL) {
		printf("argv = %s\n", p);
		p = argv[++i];
	}
	i = 0;
	p = envp[i];
	while (p != NULL) {
                //printf("envp = %s\n", p);
                p = envp[++i];
        }
	printf("exec end\n");
}
