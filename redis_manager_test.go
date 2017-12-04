package main

import (
	"testing"
)

func TestGetStringValue(t *testing.T) {
	test_key := "marathon"
	test_value := "sfai"
	redisMgr, err := NewRedisManager("0.0.0.0", "6379", 0, "")
	if err != nil {
		t.Error(err)
	}
	err = redisMgr.SetStringValue(test_key, test_value)
	if err != nil {
		t.Error(err)
		return
	}
	value, err := redisMgr.GetStringValue(test_key)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
