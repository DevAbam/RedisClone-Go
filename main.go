package main

import (
	"fmt"
	"sort"
)

type Store struct {
	data map[string]string
}

//returns every key in the store orted alphabeticaly
func (s *Store) Keys() []string{
	keys := make([]string, 0, len(s.data))
	for key := range s.data{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, bool) {
	val , ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, val string) {
	s.data[key] = val
}

func (s *Store) Del(key string){
	delete(s.data, key)
}

func NewStore() *Store {
	return &Store{
		make(map[string]string),
	}
}

func main() {
	s := NewStore()
	s.Set("a","42")
	s.Set("b","50")
	s.Del("a")
	_, exists := s.Get("a")
	fmt.Println(exists)

}
