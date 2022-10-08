#ifndef _CHAR_TRIE_H_
#define _CHAR_TRIE_H_

#include <unordered_map>
#include <string>
#include <iostream>
#include <memory>

struct TreeNode {
        TreeNode() = default;
        TreeNode(const char& c) {
                is_word_ = false;
                char_ = c;
        }
        TreeNode(const std::string& word) {
		//std::cout <<"word constructor call\n";
                is_word_ = false;
                char_ = word;
        }
	~TreeNode() {
		//std::cout <<"destructor call\n";
	}
        bool operator==(const TreeNode& tree_node) const {
                return char_ == tree_node.char_;
        }
        bool operator<(const TreeNode& tree_node) const {
                return char_ < tree_node.char_;
        }

        std::unordered_map<std::string, std::shared_ptr<TreeNode>> child;  //Save each character`s child
        std::string char_;  // Save each character 
        bool is_word_;      // If @char_ is end of a word
};


class Trie {
public:
        Trie() {
		root_ = std::make_shared<TreeNode>();
		root_->char_ = "root";
        }
        void Insert(const std::string& word);         // Build the trie by @word
        bool Search(const std::string& word);         // Some word in the tire perfect match with @word
        bool StartWith(const std::string& word);      // Some word in the tire start match with @word
        bool PrefixInclude(const std::string& word);  // Some word in the tire match with @word`s prefix or entire @word
        bool SubInclude(const std::string& word);     // Some word in the tire match with @word and @word`s subword 
	void PrintLevel();
	void Print();
private:
        std::shared_ptr<TreeNode> root_;
	void PrintTrie(std::shared_ptr<TreeNode> node, std::string start, bool last_child);
};

#endif
