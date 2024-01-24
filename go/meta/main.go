package main
import(
	"context"
	"fmt"
	"git.garena.com/shopee/deep/search-traffic-meta/go/pkg/traffic"
)


fun main() {
	ctx := context.Background()
	var err error
	ctx, err = traffic.WithCollection(ctx)
	if err != nil {
		fmt.Printf("WithGlobalSearch error = %v\n", err)
	}
}
