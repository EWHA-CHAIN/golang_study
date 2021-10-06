package dictEx

import "errors"

// Dictionary type
type Dictionary map[string]string // Dictionary는 단지 alias일 뿐(변경 가능함)

// Declare customizing err value
var errNotFound = errors.New("no such word")
var errWordExists = errors.New("word already exists")

// Search for a word
func (dict Dictionary) Search(word string) (string, error) {
	value, exists := dict[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (dict Dictionary) AddWord(word, def string) error {
	_, err := dict.Search(word)
	switch err {
	case errNotFound:
		dict[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
