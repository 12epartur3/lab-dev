#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>


int main(int argc, char *argv[],char *envp[]) {
	printf("hello, this is main, pid = %d\n", getpid());
	int r = execve("./a.exe", argv, envp);
	printf("main end, exe r = %d\n", r);
}
