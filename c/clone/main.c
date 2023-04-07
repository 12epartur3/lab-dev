#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <stdlib.h>
#include <pthread.h>

int N = 1;
pthread_attr_t attr;
pthread_key_t key;

void *t_f(void *argv) {
	pthread_setspecific(key, argv);
	int *i = pthread_getspecific(key);
	printf("i am thread %ld, argv = %d\n", pthread_self(), *i);

	*i = *i * 2;

	return i;
}

int main() {
	pthread_attr_init(&attr);
	pthread_key_create(&key, NULL);
	pthread_t t_id[N];
	for (int i = 0; i < N; i++) {
		int *argv = (int *)malloc(sizeof(int));
		*argv = i;
		pthread_create(&t_id[i], NULL, t_f, argv);
	}
	int *argv = (int *)malloc(sizeof(int));
	*argv = 999;
	pthread_setspecific(key, argv);
	argv = NULL;
	argv = pthread_getspecific(key);
	printf("i am main thread %ld, argv = %d\n", pthread_self(), *argv);
	void *res = NULL;
	for (int i = 0; i < N; i++) {
		pthread_join(t_id[i], &res);	
		if (res) {
			printf("thread %ld, res = %d\n", t_id[i], *((int *)res));
		}
	}
}
