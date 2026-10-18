package main

import (
	"errors"
	"sort"
)

type Store struct {
	data map[string]string
	maxSize int
}

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("Key is mandatory")

// returns every key in the store orted alphabeticaly
func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {

	if key == ""{
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Set(key, val string) {
	s.data[key] = val
}

func (s *Store) Del(key string) {
	delete(s.data, key)
}

func NewStore(maxSize int) *Store {
	return &Store{
		make(map[string]string),
		 maxSize,
	}
}