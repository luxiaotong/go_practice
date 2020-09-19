package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectElastic(t *testing.T) {
	InitElastic()
}

func TestSearch(t *testing.T) {
	Search(10, 1)
}

func TestIndex(t *testing.T) {
	comment := Comment{
		PKID:        1000,
		UserURL:     "datassets.cn",
		CommentTime: "2020-06-18",
		RatingNum:   "user-stars allstar20 rating",
		Content:     "还不错哦",
		UserName:    "Shannon",
		VoteCount:   100,
	}
	assert.Nil(t, Index(comment))
}

func TestGet(t *testing.T) {
	assert.True(t, Get("1000"))
}

func TestUpdate(t *testing.T) {
	assert.Nil(t, Update("1000", 101))
}

func TestDelete(t *testing.T) {
	assert.Nil(t, Delete("1000"))
}
