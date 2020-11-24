package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound  = errors.New("Didn't find the value")
	ErrKeyExists = errors.New("Key already exists")
)

func (d Dictionary) Search(k string) (string, error) {
	val, ok := d[k]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(k string, v string) error {
	_, err := d.Search(k)
	if err == ErrNotFound {
		d[k] = v
		return nil
	}
	return ErrKeyExists
}

func (d Dictionary) Update(k string, v string) error {
	_, ok := d[k]
	if !ok {
		return ErrNotFound
	}
	d[k] = v
	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
