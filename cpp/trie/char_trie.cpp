#include "char_trie.h"
#include <sstream>
#include <list>
#include <vector>
#include <pthread.h>

void Trie::Insert(const std::string& word) {
        if(word.empty()) return;
	std::unique_lock lock(wr_lock_);

	std::shared_ptr<TreeNode> node = root_;
	std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator it;
	for(const char& c: word) {
		const std::string c_str(1, c);
		it = node->child_.find(c_str);
		if(it == node->child_.end()) {
			node = node->child_.emplace(c_str, std::make_shared<TreeNode>(c_str)).first->second;	
		} else {
			node = it->second;	
		}
	}
	
        node->is_word_ = true;
        return;
}

bool Trie::Search(const std::string& word) {
	if(word.empty()) return false;
	std::shared_lock lock(wr_lock_);

        std::shared_ptr<TreeNode> node = root_;
	std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator it;
	for(const char& c: word) {
		const std::string c_str(1, c);
		it = node->child_.find(c_str);
		if(it == node->child_.end()) return false;	
		node = it->second;
	}
	return node->is_word_;
}

bool Trie::StartWith(const std::string& word) {
	if(word.empty()) return false;
	std::shared_lock lock(wr_lock_);

	std::shared_ptr<TreeNode> node = root_;
	std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator it;
	for(const char& c: word) {
                const std::string c_str(1, c);
		it = node->child_.find(c_str);
                if(it == node->child_.end()) return false;
                node = it->second;
        }
        return true;
}

bool Trie::PrefixInclude(const std::string& word) {
	if(word.empty()) return false;
	std::shared_lock lock(wr_lock_);

        std::shared_ptr<TreeNode> node = root_;
        std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator it;
	for(const char& c: word) {
                const std::string c_str(1, c);
		it = node->child_.find(c_str);
                if(it == node->child_.end()) return false;
                node = it->second;
		if(node->is_word_) return true;
        }
        return node->is_word_;
}

bool Trie::SubInclude(const std::string& word) {
	if(word.empty()) return false;

	for(int i = 0; i < word.size(); i++) {
		const std::string sub_str = word.substr(i, word.size() - i);
		//std::cout<<"sub_str = " << sub_str << '\n';
		if(PrefixInclude(sub_str))return true;
	}
	return false;
}

void Trie::PrintLayered() {
        std::shared_ptr<TreeNode> node = root_;

	std::vector<std::vector<std::shared_ptr<TreeNode>>> node_layer;
	std::list<std::shared_ptr<TreeNode>> node_list;
	node_list.push_back(root_);
	std::ostringstream oss;
	while(!node_list.empty()) {
		node_layer.push_back(std::vector<std::shared_ptr<TreeNode>>{});
		int size = node_list.size();
		for(int i = 0; i < size; i++) {
			const std::shared_ptr<TreeNode> & sptr = node_list.front();
			node_layer.back().push_back(sptr);
			for(const auto p : sptr->child_) {
				node_list.push_back(p.second);
			}
			node_list.pop_front();
			
		}
	}
	for(int i = 0; i < node_layer.size(); i++) {
		std::vector<std::shared_ptr<TreeNode>>& layer = node_layer[i];
		for(int j = 0; j < layer.size(); j++) {
                        const std::shared_ptr<TreeNode> & sptr = layer[j];
			oss << "(" << sptr->char_ <<", "<< sptr->is_word_ << ")";
			if(j == layer.size() - 1) oss << "\n";
			else oss << " ";
                }
	}
	std::cout << oss.str();
	return;
}

void Trie::PrintTree() {
        std::shared_ptr<TreeNode> node = root_;

	PrintTrie(root_, "", true);
}

void Trie::PrintTrie(std::shared_ptr<TreeNode> node, std::string start, bool last_child) {
	if(node == NULL) return;
	if(node != root_) {
		std::cout << start + "└---";			
	}
	if(node->is_word_) std::cout << "(" << node->char_ <<","<< node->is_word_ << ")\n";
	else std::cout << "(" << node->char_ <<")\n";
	int i = 0;
	for(auto const p : node->child_) {
		std::string space = "    ";
		if(last_child == false) space = "|   ";
		if(i != node->child_.size() - 1)PrintTrie(p.second, start + space, false);
		else PrintTrie(p.second, start + space, true);
		i++;
	}
	return;
}

template<typename data_type>
std::shared_ptr<TreeDataNode<data_type>> DataTrie<data_type>::Insert(const std::string& word) {
	if(word.empty()) return NULL;

        std::shared_ptr<TreeDataNode<data_type>> node = root_;
        typename std::unordered_map<std::string, std::shared_ptr<TreeDataNode<data_type>>>::iterator it;
        for(const char& c: word) {
                const std::string c_str(1, c);
                it = node->child_.find(c_str);
                if(it == node->child_.end()) {
                        node = node->child_.emplace(c_str, std::make_shared<TreeDataNode<data_type>>(c_str)).first->second;
                } else {
                        node = it->second;
                }
        }
        node->is_word_ = true;
        return node;
}

template<typename data_type>
std::shared_ptr<const TreeDataNode<data_type>> DataTrie<data_type>::Search(const std::string& word) {
	if(word.empty()) return NULL;

        std::shared_ptr<TreeDataNode<data_type>> node = root_;
        typename std::unordered_map<std::string, std::shared_ptr<TreeDataNode<data_type>>>::iterator it;
        for(const char& c: word) {
                const std::string c_str(1, c);
                it = node->child_.find(c_str);
                if(it == node->child_.end()) return NULL;
                node = it->second;
        }
        return node->is_word_ ? node:NULL;
}

template<typename data_type>
std::shared_ptr<const TreeDataNode<data_type>> DataTrie<data_type>::PrefixInclude(const std::string& word) {
	if(word.empty()) return false;

        std::shared_ptr<TreeDataNode<data_type>> node = root_;
        typename std::unordered_map<std::string, std::shared_ptr<TreeDataNode<data_type>>>::iterator it;
	for(const char& c: word) {
                const std::string c_str(1, c);
		it = node->child_.find(c_str);
                if(it == node->child_.end()) return NULL;
                node = it->second;
		if(node->is_word_) return node;
        }
        return node->is_word_ ? node:NULL;
}

template<typename data_type>
std::shared_ptr<const TreeDataNode<data_type>> DataTrie<data_type>::SubInclude(const std::string& word) {
	if(word.empty()) return NULL;

	std::shared_ptr<TreeDataNode<data_type>> node = NULL;
	for(int i = 0; i < word.size(); i++) {
		const std::string sub_str = word.substr(i, word.size() - i);
		//std::cout<<"sub_str = " << sub_str << '\n';
		node = PrefixInclude(sub_str);
		if(node) return node;
	}
	return NULL;
}

template<typename data_type>
void DataTrie<data_type>::PrintTree() {
        std::shared_ptr<TreeDataNode<data_type>> node = root_;

	PrintTrie(root_, "", true);
}

template<typename data_type>
void DataTrie<data_type>::PrintTrie(std::shared_ptr<TreeDataNode<data_type>> node, std::string start, bool last_child) {
	if(node == NULL) return;
	if(node != root_) {
		std::cout << start + "└---";			
	}
	if(node->is_word_) std::cout << "(" << node->char_ <<","<< node->is_word_ << ")\n";
	else std::cout << "(" << node->char_ <<")\n";
	int i = 0;
	for(auto const p : node->child_) {
		std::string space = "    ";
		if(last_child == false) space = "|   ";
		if(i != node->child_.size() - 1)PrintTrie(p.second, start + space, false);
		else PrintTrie(p.second, start + space, true);
		i++;
	}
	return;
}

int main() {
	std::unordered_map<std::string, std::shared_ptr<TreeNode>> test;
	std::pair<std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator, bool> p1 = test.emplace("myname", std::make_shared<TreeNode>("Yuanye2"));	
	std::pair<std::unordered_map<std::string, std::shared_ptr<TreeNode>>::iterator, bool> p2 = test.insert({"myname", std::make_shared<TreeNode>("Yuanye")});	
	p1.first->second->char_ = "Yuanye3";
	//std::cout<<"p1.second = " << p1.second <<" p1.first = " << p1.first->second->char_ << '\n';
	//std::cout<<"p2.second = " << p2.second <<" p2.first = " << p2.first->second->char_ <<'\n';
	std::cout << std::boolalpha;
	Trie T;
	T.Insert("Jay");
	T.Insert("Yuanye");
	T.Insert("Abner");
	T.Insert("Jack");
	T.Insert("Bartholomew");
	T.Insert("Albert");
	T.Insert("Albin");
	T.Insert("Ben");
	T.Insert("Bryan");
	T.Insert("Bradford");
	T.PrintTree();
	T.PrintLayered();
	std::cout << "T.Search(Jay) =  "<< T.Search("Jay") << '\n';
	std::cout << "T.Search(Yuanye) =  "<< T.Search("Yuanye") << '\n';
	std::cout << "T.Search(Yuanye2) =  "<< T.Search("Yuanye2") << '\n';

	std::cout << "T.StartWith(Jac) =  "<< T.StartWith("Jac") << '\n';
	std::cout << "T.StartWith(yu) =  "<< T.StartWith("yu") << '\n';
	std::cout << "T.StartWith(Yu) =  "<< T.StartWith("Yu") << '\n';

	std::cout << "T.PrefixInclude(Benjamin) =  "<< T.PrefixInclude("Benjamin") << '\n';
	std::cout << "T.PrefixInclude(Albin) =  "<< T.PrefixInclude("Albin") << '\n';
	std::cout << "T.PrefixInclude(Yu) =  "<< T.PrefixInclude("Yu") << '\n';

	std::cout << "T.SubInclude(Super Yuanye) =  "<< T.SubInclude("Super Yuanye") << '\n';
	std::cout << "T.SubInclude(Super Benman) =  "<< T.SubInclude("Super Benman") << '\n';
}

