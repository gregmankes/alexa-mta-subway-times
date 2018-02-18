package mta

import (
	"fmt"
	"strings"
)

type Feed int

const (
	OneThroughSixAndSFeed Feed = 1
	ACEFeed               Feed = 26
	NQRWFeed              Feed = 16
	BDFMFeed              Feed = 21
	LFeed                 Feed = 2
	GFeed                 Feed = 31
	JZFeed                Feed = 36
	UndefinedFeed         Feed = 0
)

var feedMap = map[string]Feed{
	"1": OneThroughSixAndSFeed,
	"2": OneThroughSixAndSFeed,
	"3": OneThroughSixAndSFeed,
	"4": OneThroughSixAndSFeed,
	"5": OneThroughSixAndSFeed,
	"6": OneThroughSixAndSFeed,
	"S": OneThroughSixAndSFeed,
	"A": ACEFeed,
	"C": ACEFeed,
	"E": ACEFeed,
	"N": NQRWFeed,
	"Q": NQRWFeed,
	"R": NQRWFeed,
	"W": NQRWFeed,
	"L": LFeed,
	"G": GFeed,
	"J": JZFeed,
	"Z": JZFeed,
}

func getFeed(lineName string) Feed {
	return feedMap[strings.ToUpper(lineName)]
}

func GetFeedData(apiKey, lineName, stop string) error {
	feedID := getFeed(lineName)
	if feedID == UndefinedFeed {
		return newError(fmt.Sprintf("Feed for line name %s is undefined", lineName), FeedUndefinedErrorType)
	}
	return nil
}
