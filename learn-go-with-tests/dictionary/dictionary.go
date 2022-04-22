package main

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	return d[key], nil
}
