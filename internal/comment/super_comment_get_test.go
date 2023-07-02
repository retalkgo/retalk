package comment

import (
	"retalk/internal/core"
	"retalk/internal/query"
	"testing"
)

func TestSuperCommentGet(t *testing.T) {
	core.InitCore()
	rawData, _ := query.Comment.Find()
	_, err := SuperCommentGet(rawData)
	if err != nil {
		t.Errorf(err.Error())
	}
}