package main

import "errors"

type Dictionary map[string]string

// map is *a pointer* to a runtime.hmap structure
// so copied receivers of a map can manipulate the values unlike struct

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
