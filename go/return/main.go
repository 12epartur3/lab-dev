package main
import(
	"fmt"
)


type DashboardStatKey struct {
	Keyword         string
	KeywordType     string
	CategoryCluster string
}
func (k DashboardStatKey) WithKeywordType(keywordType string) DashboardStatKey {
	k.KeywordType = keywordType
	return k
}

func (k DashboardStatKey) WithCategoryCluster(categoryCluster string) DashboardStatKey {
	k.CategoryCluster = categoryCluster
	return k
}

func main() {
	key := DashboardStatKey{
		Keyword: "keyyyy",
	}
	fmt.Printf("key = %+v\n", key)
	k1 := key.WithKeywordType("t1")
	fmt.Printf("key = %+v\n", key)
	fmt.Printf("k1 = %+v\n", k1)
}
