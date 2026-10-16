package main

import (
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a","42")
	s.Set("b","50")
	s.Del("a")

	val , err := s.Get("a")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(val)

}
