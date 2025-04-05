package maps

const (
	ErrKeyNotFound = DictionaryErr("Key not found in map")
	ErrKeyExists   = DictionaryErr("Key already exists")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	defination, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = val
	case nil:
		return ErrKeyExists
	}
	//default behaviour not sure shoudl we use
	// Also wondering about strong pattern matching ?
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)
	// if err != nil {
	// 	return ErrKeyNotFound
	// }
	switch err {
	case ErrKeyNotFound:
		return err
	case nil:
		d[key] = val
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFound:
		return err
	case nil:
		delete(d, key)
	}
	return nil
}
