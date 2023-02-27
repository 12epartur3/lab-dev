#ifndef _CHAR_TRIE_H_
#define _CHAR_TRIE_H_

#include <unordered_map>
#include <string>
#include <iostream>
#include <memory>
#include <shared_mutex>

struct TreeNode {
        TreeNode() = default;
        TreeNode(const char& c) {
		std::cout <<"char constructor call\n";
                char_ = c;
                is_word_ = false;
        }
        TreeNode(const std::string& word): char_(word){
		//std::cout <<"word constructor call\n";
                is_word_ = false;
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

        std::string char_;  // Save each character 
        bool is_word_;      // If @char_ is end of a word
        std::unordered_map<std::string, std::shared_ptr<TreeNode>> child_;  //Save each character`s child
};

template <typename data_type>
struct TreeDataNode {
	TreeDataNode() = default;
	TreeDataNode(const char& c) {
		std::cout <<"data char constructor call\n";
                char_ = c;
		is_word_ = false;
		data_ptr_ = NULL;
	}
	TreeDataNode(const std::string& word): char_(word) {
		is_word_ = false;
		data_ptr_ = NULL;
	}

	void set_data(std::shared_ptr<data_type> data_ptr) {
		data_ptr_ = data_ptr;	
		return;
	}
	std::shared_ptr<data_type> get_data() {
		return data_ptr_;
	}
	~TreeDataNode() {
	}
	// Save each character
        std::string char_;

	// If @char_ is end of a word
        bool is_word_;

	// Save each character`s child
        std::unordered_map<std::string, std::shared_ptr<TreeDataNode<data_type>>> child_;
private:	
	// User data
	std::shared_ptr<data_type> data_ptr_;
};

/*
Thread Safe
*/
class Trie {
public:
        Trie() {
		root_ = std::make_shared<TreeNode>();
		root_->char_ = "root";
        }
        void Insert(const std::string& word);         // Build the trie by @word
        bool Search(const std::string& word);         // Some word in the trie perfect match with @word
        bool StartWith(const std::string& word);      // @Word is some word`s prefix in the trie
        bool PrefixInclude(const std::string& word);  // Some word in the trie match with @word`s prefix or entrie @word
        bool SubInclude(const std::string& word);     // Some word in the trie match with @word and @word`s subword 
	void PrintLayered();			      // Print the trie	in layered
	void PrintTree();			      // Print the trie in a human-readable format	
private:
        std::shared_ptr<TreeNode> root_;
	std::shared_mutex wr_lock_;
	void PrintTrie(std::shared_ptr<TreeNode> node, std::string start, bool last_child);
};

/*
Not Thread Safe
*/
template<typename data_type>
class DataTrie {
public:
	DataTrie() {
		root_ = std::make_shared<TreeDataNode<data_type>>();
		root_->char_ = "root";
	}

	// Build the trie by @word
	// @Insert not thread safe
	// Do not @Search,@PrefixInclude and @SubInclude when @Insert
	std::shared_ptr<TreeDataNode<data_type>> Insert(const std::string& word);         

	// Some word in the trie perfect match with @word
        std::shared_ptr<const TreeDataNode<data_type>> Search(const std::string& word);         

	// Some word in the trie match with @word`s prefix or entrie @word
        std::shared_ptr<const TreeDataNode<data_type>> PrefixInclude(const std::string& word);

	// Some word in the trie match with @word and @word`s subword
        std::shared_ptr<const TreeDataNode<data_type>> SubInclude(const std::string& word); 

	// Print the trie in a human-readable format
	void PrintTree();
private:
	std::shared_ptr<TreeDataNode<data_type>> root_;
	void PrintTrie(std::shared_ptr<TreeDataNode<data_type>> node, std::string start, bool last_child);
};

#endif
