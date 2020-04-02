package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

// NewID snowflake
// +--------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
// +--------------------------------------------------------------------------+
//
func NewID(node int64) (id int64, err error) {
	node2, err := snowflake.NewNode(node)
	if err != nil {
		return
	}

	id = node2.Generate().Int64()

	return
}
