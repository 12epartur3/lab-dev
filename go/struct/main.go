package main

import(
	"fmt"
	"github.com/my/project/subdir"
	"encoding/json"
	"time"
)



type ItemInfo struct {
	Title             string
	RelevanceModelData example.RelevanceModelDataRaw
}

type DebugInfo struct {
	DebugLevel int `json:"debug_level"`
	Rs         struct {
		Timestamp int64 `json:"timestamp"`
	} `json:"rs"`
}

type RelevanceModelData struct {
	RelRawScore float64
	RelLxScore  float64
	RelAddScore float64
}

func convertDebugInfo(timeStamp int64) (debugInfoStr string) {
	debugInfo := &DebugInfo{}
	debugInfo.DebugLevel = 1
	debugInfo.Rs.Timestamp = timeStamp
	debugInfoByte, err := json.Marshal(debugInfo)
	if err == nil {
		debugInfoStr = string(debugInfoByte)
	}
	return
}
type MartSearch struct{}
var _ S = MartSearch{}
func main() {
	itemInfo := &ItemInfo{}
	//itemInfo.RelevanceModelData = &RelevanceModelData{}
	itemInfo.Title = "yuanye";
	itemInfo.RelevanceModelData.RelRawScore = 1.25
	fmt.Printf("itemInfo.Title = %v\n", itemInfo.Title)
	fmt.Printf("itemInfo.RelevanceModelData.RelRawScore = %v\n", itemInfo.RelevanceModelData.RelRawScore)
	fmt.Printf("DebugInfo = %s\n", convertDebugInfo(time.Now().Unix()))
}
