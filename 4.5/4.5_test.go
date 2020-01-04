package main_test

import (
	"golang-exercises/4.5"
	"reflect"
	"testing"
)

func TestDeleteByAppend(t *testing.T) {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	got := main.DeleteByAppend(x)
	want := []int{42, 43, 44, 48, 49, 50, 51}

	if !reflect.DeepEqual(got, want) {
		t.Error("error!")
	}
}
