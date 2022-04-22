package main

type Dictionary map[string]string

func Search(dict map[string]string, key string) string {
	return dict[key]
}
