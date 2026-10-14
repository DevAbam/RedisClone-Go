package main

import (
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a","42")
	s.Set("b","50")
	s.Del("a")
	_, exists := s.Get("a")
	fmt.Println(exists)

}
