package dictEx

import "errors"

// Dictionary type
type Dictionary map[string]string // Dictionary는 단지 alias일 뿐(변경 가능함)

// Declare customizing err value
var (
	errNotFound   = errors.New("no such word")
	errCantUpdate = errors.New("cant update non-existing word")
	errWordExists = errors.New("word already exists")
)

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

func (dict Dictionary) UpdateWord(word, def string) error {
	_, err := dict.Search(word)
	if err != nil {
		return errCantUpdate
	} else { // 기존에 존재하는 word에 대해서만 업데이트 가능
		dict[word] = def
	}
	return nil
}

func (dict Dictionary) DeleteWord(word string) {
	delete(dict, word) // builtin.go에 정의되어 있음
}
