package goapp

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPut(t *testing.T){
	t.Run("single put", func(t *testing.T){
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

	})

}

func TestMakeId(t *testing.T){
	t.Run("Different input different result", func(t *testing.T){
		first := "https://google.com"
		second := "https://github.com"
		firstResult := makeId(first)
		secondResult := makeId(second)
		if firstResult == secondResult{
			t.Errorf("expect different output but got %s and %s", firstResult, secondResult)
		}
	})
	t.Run("Same input same result", func(t *testing.T){
		input := "https://github.com"
		firstResult := makeId(input)
		secondResult := makeId(input)
		if firstResult != secondResult {
			t.Errorf("expect same result from input %s but got %s and %s", input, firstResult, secondResult)
		}
	})
}
