package comment

import (
	"testing"

	"github.com/retalkgo/retalk/internal/core"
	"github.com/retalkgo/retalk/internal/query"
)

func TestSuperCommentGet(t *testing.T) {
	core.InitCore()
	rawData, _ := query.Comment.Find()
	_, err := SuperCommentGet(rawData)
	if err != nil {
		t.Errorf(err.Error())
	}
}
