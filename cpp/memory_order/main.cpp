#include <pthread.h>
#include <stdio.h>
#include <atomic>

int x, y;
int a, b;
pthread_rwlock_t wr_lock; 
// x = 1, a = 0, y = 1, b = 1
// x = 1, y = 1, a = 1, b = 1
// y = 1, b = 0, x = 1, a = 1
// y = 1, x = 1, a = 1, b = 1
void* t_f1(void* arg) {
	x = 1;
	std::atomic_thread_fence(std::memory_order_release);
	//std::atomic_thread_fence(std::memory_order_acquire);
	a = y;
	return NULL;
}
void* t_f2(void* arg) {
	y = 1;
	std::atomic_thread_fence(std::memory_order_release);
	//std::atomic_thread_fence(std::memory_order_acquire);
	b = x;
	return NULL;
}

int main() {
	unsigned long c = 0;
	while (1) {
		x = 0, y = 0;
		a = 0, b = 0;
		int r = pthread_rwlock_init(&wr_lock, NULL);
		if (r != 0) {
			printf("init lock, r = %d\n", r);
			return 0;
		}
		pthread_t tid1, tid2;
		pthread_create(&tid1, NULL, t_f1, NULL);
		pthread_create(&tid2, NULL, t_f2, NULL);
		pthread_rwlock_destroy(&wr_lock);
		if (r != 0) {
			printf("destroy lock, r = %d\n", r);
			return 0;
		}
		pthread_join(tid1, NULL);
		pthread_join(tid2, NULL);
		if (a == 0 && b == 0) {
			printf("re order, c = %lu\n", c);
			return 0;
		}
		c++;
	}
}
