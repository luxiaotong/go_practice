package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGInitElastic(t *testing.T) {
	assert.Nil(t, GInitElastic())
}

func TestGSearch(t *testing.T) {
	assert.Nil(t, GSearch())
}

func TestGIndex(t *testing.T) {
	comment := Comment{
		PKID:        1001,
		UserURL:     "datassets.cn",
		CommentTime: "2020-06-19",
		RatingNum:   "user-stars allstar20 rating",
		Content:     "测试测试测试测试测试",
		UserName:    "Shannon",
		VoteCount:   100,
	}
	assert.Nil(t, GIndex(comment))
}

func TestGGet(t *testing.T) {
	assert.Nil(t, GGet("1001"))
}

func TestGUpdate(t *testing.T) {
	comment := Comment{
		PKID:        1001,
		UserURL:     "datassets.cn",
		CommentTime: "2020-06-19",
		RatingNum:   "user-stars allstar20 rating",
		Content:     "测试测试测试测试测试",
		UserName:    "Shannon",
		VoteCount:   88,
	}
	assert.Nil(t, GUpdate(comment))
}

func TestGDelete(t *testing.T) {
	assert.Nil(t, GDelete("1001"))
}
