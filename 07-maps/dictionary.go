package maps

import (
	"errors"
)

const (
	ErrNotFound         = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word exists")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(key, def string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		d[key] = def
	case err == nil:
		return ErrWordExists
	default:
		return nil
	}
	return nil
}

func (d Dictionary) Update(key, def string) error {
	_, err := d.Search(key)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[key] = def
	default:
		return nil
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
