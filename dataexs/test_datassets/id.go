package testdatassets

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node   *snowflake.Node
	idOnce sync.Once
)

func GenID() int64 {
	idOnce.Do(func() {
		var err error
		node, err = snowflake.NewNode(int64(1))
		if err != nil {
			panic(err)
		}
	})
	return node.Generate().Int64()
}
