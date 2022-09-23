#include <set>
#include <string>
#include <iostream>

struct data {
	int position;
	std::string val;
	bool operator<(const data& d) const {
		return position < d.position;
	}
};

class method {
public:
	void print_data(const data& d) {
		std::cout << "data = " << d.position << ":" <<d.val << '\n';
	}
};


int main() {
	std::set<data> data_set;
	data d1;
	d1.position = 1;
	d1.val = "v1";
	data d2;
	d2.position = 2;
	d2.val = "v2";
	data_set.insert(std::move(d1));
	data_set.insert(std::move(d2));
	std::set<data>::iterator it = data_set.begin();
	method m;
	m.print_data(*it);
}

