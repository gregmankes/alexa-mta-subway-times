package mta

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetFeedData(t *testing.T) {
	apiKey := "12d3b9fc4e8c2506f0c22d866e1275ce"
	lineName := "l"
	duractions, _ := GetFeedData(apiKey, lineName, "MORGAN AV", "north")
	spew.Dump(duractions)
	t.Fail()
}
