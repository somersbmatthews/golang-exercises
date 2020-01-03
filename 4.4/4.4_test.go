package main_test

import (
	"golang-exercises/4.4"
	"reflect"
	"testing"
)

func TestAppendSlice(t *testing.T) {
	x := []int{42, 43, 44, 45, 46, 47, 48, 50, 51}
	got := main.AppendSlice(x)
	want := []int{42, 43, 44, 45, 46, 47, 48, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}

	if !reflect.DeepEqual(got, want) {
		t.Error("error!")
	}
}
