#include "item.h"

#include <unistd.h>
#include <map>
#include <string>
#include <vector>
#include <unordered_map>
#include <memory>
#include <chrono>
#include <thread>

struct ImprQuery {
    std::string query_raw;
    std::string query_norm;
    std::string cate_l1;
    float cate_l1_score;
    std::string cate_l2;
    float cate_l2_score;
};

struct NerInfo {
    std::string text;    // 单个ner结果实体文本
    std::string type;    // 单个ner结果实体类型, 单个text可能是多个类型["哈利波特"], 如"1;2;3"
    double score;        // 单个ner结果实体置信度, 置信度>0.5时, 表示结果可用.
    std::string source;  // 当前ner的来源: query, pure_caption, hashtag, ocr, unknown. 不输入的话就是unknown.

    NerInfo() {}
    NerInfo(const NerInfo& ni) {
        text = ni.text;
        type = ni.type;
        score = ni.score;
        source = ni.source;
    }
};

struct HisSearchQuery {
    std::string query;
    std::string tokens;
    std::string norm;
    std::vector<NerInfo> ners;
    std::vector<float> embedding;
    uint64_t timestamp = 0UL;  // 13 位时间戳
    float query_freshness = 1.0F;
    std::vector<std::shared_ptr<ImprQuery>> features;
}; 

struct HashTag {
    std::string hashtag;
    std::string tokens;  // 分词
    std::string norm;  // 词干
    std::vector<NerInfo> ners;
    std::vector<float> embedding;
    std::vector<std::shared_ptr<ImprQuery>> features;
    double l1_chisquare_weight;
    double chisquare_sum_weight;
    HashTag() {}
    HashTag(const HashTag& ht) {
        hashtag = ht.hashtag;
        tokens = ht.tokens;
        norm = ht.norm;
        ners = ht.ners;
        embedding = ht.embedding;
        features = ht.features;
        l1_chisquare_weight = ht.l1_chisquare_weight;
        chisquare_sum_weight = ht.chisquare_sum_weight;
    }
};

struct Qanchor {
    double weight;
    std::string text;
    std::string norm;
    std::string tokens;
    std::vector<float> embedding;
};
struct Ner {
    std::string text;
    std::string tokens;
    std::string norm;
    NerInfo ner_info;
    std::vector<float> embedding;
    std::vector<std::shared_ptr<ImprQuery>> features;
};

struct FeedVideo {
    uint64_t video_id;
    std::vector<HashTag> hashtags;
    std::vector<Ner> ners;
    std::vector<Qanchor> qanchors;
    std::string label_l1_name;
    float video_freshness = 1.0F;
};

class Context: public Item{
public:
        Context() {
                //std::cout << "Context constructor call\n";
        }
        virtual ~Context() {
                //std::cout << "Context destructor call\n";
        }
private:
        int t;
        std::map<std::string, std::string> mp;
        std::vector<std::shared_ptr<HisSearchQuery>> history_queries;  //历史搜索query
        std::shared_ptr<FeedVideo> refer_video;
        std::vector<std::shared_ptr<FeedVideo>> longview_video_list;  // finished video list
        std::vector<std::shared_ptr<FeedVideo>> like_video_list;
        std::vector<std::shared_ptr<FeedVideo>> forward_video_list;  // share video list
        std::vector<std::shared_ptr<FeedVideo>> follow_video_list;

        std::vector<std::string> viewed_queries;
        std::map<std::string, bool> viewed_query_map;

        //std::vector<std::shared_ptr<ImprQuery>> item_word_impr_queries;
        std::map<std::string, std::shared_ptr<ImprQuery>> item_word_impr_queries;
        std::vector<std::string> exposured_queries;
        std::unordered_map<std::string, float> user_cate_l1_score_map;
};

static int N = 99999;
void test1_func() {
	for(int i = 0; i < N; i++) {
		Context context;
		//std::cout<<"tid = " <<std::this_thread::get_id() << '\n';
	}	
}
void test2_func() {
	for(int i = 0; i < N; i++) {
		Item* it = new Context;
		delete it;
		//std::cout<<"tid = " <<std::this_thread::get_id() << '\n';
	}	
}
int main() {
	std::cout << "sizeof(Context) = " << sizeof(Context) << '\n';
	std::vector<std::thread> t_v;
	std::vector<std::thread> t_v2;
	for(int i = 0; i < 5; i++) {
		std::thread t(test1_func);
		t_v.push_back(std::move(t));

		std::thread t2(test2_func);
		t_v2.push_back(std::move(t2));
	}
	auto start = std::chrono::steady_clock::now();
	test1_func();
	for(auto& t : t_v) {
		t.join();
	}
	auto end = std::chrono::steady_clock::now();
	std::cout << "test1_func elapsed time: " << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << "ms\n";
	start = std::chrono::steady_clock::now();
	//test2_func();
	for(auto& t : t_v2) {
		//t.join();
	}
	end = std::chrono::steady_clock::now();
	std::cout << "test2_func elapsed time: " << std::chrono::duration_cast<std::chrono::milliseconds>(end-start).count() << "ms\n";
	std::cout << "end main\n";
	
}
