package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		make(map[string]string),
	}
}

func main() {
	fmt.Println("GoKV")
}
