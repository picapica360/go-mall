package snowflake

import (
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
)

func TestNewID(t *testing.T) {
	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())

	time.Sleep(10 * time.Millisecond)

	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())
	t.Logf("snowflake id:%d\n", NewID())
}

func TestSnowflakeId(t *testing.T) {
	node2, err := snowflake.NewNode(1)
	if err != nil {
		return
	}

	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
	t.Logf("snowflake1 id:%d\n", node2.Generate().Int64())
}
