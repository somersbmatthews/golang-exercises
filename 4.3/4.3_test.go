package main_test

import (
	"testing"
	"golang-exercises/4.3"
	"reflect"
)

func TestPrint(t *testing.T){

	got_1, got_2, got_3, got_4, got_5 := main.Slicing()
	want_1, want_2, want_3, want_4, want_5 := []int{42, 43, 44, 45, 46}, []int{47, 48, 49, 50, 51}, []int{44, 45, 46, 47, 48}, []int{43, 44, 45, 46, 47}, []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	if (!reflect.DeepEqual(got_1, want_1)) {
		t.Error("problem with first set")	
	}

	if (!reflect.DeepEqual(got_2, want_2)) {
		t.Error("problem with second set")
	}

	if (!reflect.DeepEqual(got_3, want_3)) {
		t.Error("problem with third set")
	}
	
	if (!reflect.DeepEqual(got_4, want_4)) {
		t.Error("problem with four set")
	}

	if (!reflect.DeepEqual(got_5, want_5)) {
		t.Error("problem with five set")
	}


}