package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("charlie", "3")
	store.Set("alpha", "1")
	store.Set("bravo", "2")

	got := store.Keys()
	want := []string{"alpha","bravo","charlie"}

	if !reflect.DeepEqual(got,want){
		t.Errorf("Keys() = %v, want %v", got,want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	// store.Set("a","100")

	got := store.Keys()
	want := []string{}

	if !reflect.DeepEqual(got,want){
		t.Errorf("Keys() on empty store = %v, want empty slice",got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	//Make a store
	store := NewStore()

	//set hello -> world
	store.Set("Hello", "world")

	//call get("hello") assert if it returns ("world", true)
	if val, ok := store.Get("Hello"); !ok || val != "world"{
		t.Error("Get() failed")
	}

	//call get("missing") assert if it returns ("", false)
	if val, ok := store.Get("missing"); ok || val != ""{
		t.Error("Get() failed")
	}

}