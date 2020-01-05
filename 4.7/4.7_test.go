package main_test

import (
	"golang-exercises/4.7"
	"reflect"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	got := main.MakeSlice()
	want := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	if !reflect.DeepEqual(got, want) {
		t.Error("error!")
	}
}
