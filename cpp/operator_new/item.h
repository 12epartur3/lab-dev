#ifndef _ITEM_H_
#define _ITEM_H_

#include <iostream>
#include <string>
#include <list>
#include <mutex>
#include <shared_mutex>

class Item {
public:
	Item() {
		//std::cout << "Item constructor call\n";
	}
	virtual ~Item() {
		//std::cout << "Item destructor call\n";
	}
	void* operator new(std::size_t size);
	void operator delete(void* ptr); 
private:
	static std::shared_mutex list_lock_;
	static std::list<void*> item_list_;
};

std::shared_mutex Item::list_lock_;
std::list<void*> Item::item_list_;

void* Item::operator new(std::size_t size) {
	void* ptr = NULL;
	list_lock_.lock_shared();
	if(item_list_.empty()) {
		//std::cout << "Item operator new call malloc, item_list_ size = " << item_list_.size() << "\n";
		list_lock_.unlock();
		ptr = malloc(size);
		return ptr;
	}
	list_lock_.unlock();

	std::unique_lock lock(list_lock_);
	//std::cout << "Item operator new call pop_front, item_list_ size = " << item_list_.size() << "\n";
	ptr = item_list_.front();
	item_list_.pop_front();
	return ptr; 
}

void Item::operator delete(void* ptr) {
	list_lock_.lock_shared();
	if(item_list_.size() > 10) {
		//std::cout << "Item operator delete call free, item_list_ size = " << item_list_.size() << "\n";
		list_lock_.unlock();
		free(ptr);
	}
	list_lock_.unlock();

	std::unique_lock lock(list_lock_);
	item_list_.push_back(ptr);	
	//std::cout << "Item operator delete call push_back, item_list_ size = " << item_list_.size() << "\n";
	return;
}

#endif
