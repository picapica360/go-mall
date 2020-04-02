package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var _node *snowflake.Node

func init() {
	build(1)
}

// NewID snowflake
// +--------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
// +--------------------------------------------------------------------------+
//
func NewID() int64 {
	return _node.Generate().Int64()
}

func build(node int64) {
	node2, err := snowflake.NewNode(node)
	if err != nil {

	}
	_node = node2
}
