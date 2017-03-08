package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKeywords_New(t *testing.T) {

}

func TestKeywords_Get(t *testing.T) {

	k := &Keywords{}
	k = k.New()
	res, err := k.Get("lobster")
	t.Log(res)
	assert.Equal(t, Value{}, res, "should be equal objects")
	assert.EqualValues(t, Value{Code: 19}, res, "should be equal values")
	assert.Equal(t, nil, err, "error should ne nil")

}
