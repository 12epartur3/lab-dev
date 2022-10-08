#include <iostream>
#include <unordered_set>
#include <set>
#include <string>
#include <memory>
#include <functional>
struct TreeNode {
	TreeNode() = default;
	TreeNode(const char& c) {
		is_word = false;
		c_ = c;
	}
	bool operator==(const TreeNode& tree_node) const {
		return c_ == tree_node.c_;
	}
	bool operator<(const TreeNode& tree_node) const {
		return c_ < tree_node.c_;
	}
	std::set<TreeNode> node_set;
	char c_; // Save each character
	bool is_word; // If @c_ is end of a word
};

struct TreeNodeHash {
	size_t operator()(const TreeNode& tree_node) const {
		return std::hash<char>()(tree_node.c_);
	}
};


class Trie {
public:
	Trie() {
	}
	void Insert(std::string& str);
	bool Search(std::string& str);
private:
	TreeNode root_;
};

void Trie::Insert(std::string& str) {
	if(str.empty()) return;
	TreeNode* node  = &root_;
	std::set<TreeNode>::iterator it;
	for(const char& c : str) {
		it = node->node_set.find(c);
		if(it == node->node_set.end()) {
			//node = const_cast<TreeNode*>(&(*node->node_set.insert(c).first));
			//node->node_set.insert(c);
		} else {
			node = const_cast<TreeNode*>(&*it); 
		}
	}
	node->is_word = true;
	return;
}

int main() {
	std::set<TreeNode> node_set;
	node_set.insert('a');
	node_set.insert(TreeNode('b'));
	std::cout << node_set.count('a') << '\n';
	std::cout << node_set.count('b') << '\n';
	std::cout << node_set.count(TreeNode('a')) << '\n';
	std::cout << node_set.count(TreeNode('b')) << '\n';
	return 0;	
}

