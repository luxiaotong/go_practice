package sms

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildWWSE(t *testing.T) {
	buildWSSE()
}

func TestGenCode(t *testing.T) {
	assert.NotEqual(t, GenCode(), GenCode())
}

func TestSendCode(t *testing.T) {
	t.Skip("skip send message")
	err := SendCode("+86185XXXXXXXX", GenCode())
	fmt.Println("err: ", err)
}

func TestSetCode(t *testing.T) {
	assert.Nil(t, SetCode("+8618500000000", 111111))
	assert.Nil(t, SetCode("+8618500000001", 111112))
}

func TestCheckCode(t *testing.T) {
	assert.Nil(t, CheckCode("+8618500000000", 111111))
	assert.NotNil(t, CheckCode("+8618500000001", 111113))
}

func TestCheckCode_Expired(t *testing.T) {
	t.Skip("skip wait 60 secs")
	time.Sleep(60 * time.Second)
	assert.NotNil(t, CheckCode("+8618500000000", 111111))
}
