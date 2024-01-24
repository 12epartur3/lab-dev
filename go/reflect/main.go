package main

import(
	"fmt"
	//"os"
	"reflect"
	//"io"
	"encoding/json"
)

var rank = `{"instance_id":"6079401272","scores":{"boosting_score":0.019720852206709306,"ctr_score":0.001862536882981658,"cvr_cvr_score":0.043600648641586304,"cvr_score":0.043600648641586304,"dsrm_ctr_score":0.001862536882981658,"dsrm_cvr_score":0.09560093283653259,"dsrm_l2p_score":0.0001780602615326643,"final_score":0.019720852206709306,"merge_score":0.019720852206709306,"query_item_embedding_score":0.019720852206709306,"rel_add_score":0.099999,"rel_lx_score":0.0,"rel_qcp_all_score":0.019720852206709306,"rel_qcp_el_score":0.019720852206709306,"rel_raw_score":0.0},"summary":{"rank":16,"score":0.019720852206709306}}`

type RankInfo struct {
	InstanceID string  `json:"instance_id,omitempty"`
	Scores     Scores  `json:"scores,omitempty"`
	Summary    Summary `json:"summary,omitempty"`
}

type Scores struct {
	BoostingScore           float64 `json:"boosting_score,omitempty"`
	CtrScore                float64 `json:"ctr_score,omitempty"`
	CvrCvrScore             float64 `json:"cvr_cvr_score,omitempty"`
	CvrScore                float64 `json:"cvr_score,omitempty"`
	DsrmCtrScore            float64 `json:"dsrm_ctr_score,omitempty"`
	DsrmCvrScore            float64 `json:"dsrm_cvr_score,omitempty"`
	DsrmL2PScore            float64 `json:"dsrm_l2p_score,omitempty"`
	FinalScore              float64 `json:"final_score,omitempty"`
	MergeScore              float64 `json:"merge_score,omitempty"`
	QueryItemEmbeddingScore float64 `json:"query_item_embedding_score,omitempty"`
	RelAddScore             float32 `json:"rel_add_score,omitempty"`
	RelLxScore              float32 `json:"rel_lx_score,omitempty"`
	RelQcpAllScore          float64 `json:"rel_qcp_all_score,omitempty"`
	RelQcpElScore           float64 `json:"rel_qcp_el_score,omitempty"`
	RelRawScore             float32 `json:"rel_raw_score,omitempty"`
}

type Summary struct {
	Rank  int     `json:"rank,omitempty"`
	Score float64 `json:"score,omitempty"`
}

type DebugInfoCategoryMap map[string] map[int64] interface{}

func main() {
	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w))
	debugInfoCategoryMap := make(DebugInfoCategoryMap)
	var ItemID int64
	ItemID = 123
	if rank != "" {
		debugInfoCategoryMap["rank"] = make(map[int64]interface{})
		rankInfo := &RankInfo{}
		if err := json.Unmarshal([]byte(rank), rankInfo); err == nil {
			debugInfoCategoryMap["rank"][ItemID] = rankInfo
			fmt.Printf("ItemID = %d\n", ItemID)
		}
	}
	var id int64
	id = 123
	if _, ok := debugInfoCategoryMap["rank"]; ok {
		if info, ok := debugInfoCategoryMap["rank"][id]; ok {
			if val, ok := info.(*RankInfo); ok {
				fmt.Printf("ItemID = %d, RelAddScore = %f, RelLxScore = %f, RelRawScore = %f\n", id, val.Scores.RelAddScore, val.Scores.RelLxScore, val.Scores.RelRawScore)
			} else {
				fmt.Printf("reflect.TypeOf(info) = %v\n", reflect.TypeOf(info))
			}
		}
	}
}

