package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find a word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, value string) {
	d[word] = value
}
