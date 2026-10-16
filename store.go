package main

import (
	"errors"
	"sort"
)

type Store struct {
	data map[string]string
}

var ErrKeyDoesNotExist = errors.New("Key does not exist")

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

func NewStore() *Store {
	return &Store{
		make(map[string]string),
	}
}