package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrUnknownWord        = DictionaryErr("Unknown word requested")
	ErrWordAlreadyDefined = DictionaryErr("Word already defined, cannot add again")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	if !found {
		return "", ErrUnknownWord
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, found := d[key]
	if found {
		return ErrWordAlreadyDefined
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, found := d[key]
	if !found {
		return ErrUnknownWord
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, found := d[key]
	if !found {
		return ErrUnknownWord
	}

	delete(d, key)
	return nil
}
