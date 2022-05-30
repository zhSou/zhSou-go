package dict

import (
	"encoding/gob"
	"io"
)

// Dict 分词到分词id的映射
type Dict struct {
	Data map[string]int
}

func NewDict() *Dict {
	return &Dict{
		Data: make(map[string]int),
	}
}

func Load(r io.Reader) (*Dict, error) {
	var dict Dict
	err := gob.NewDecoder(r).Decode(&dict)
	if err != nil {
		return nil, err
	}
	return &dict, nil
}

func (d *Dict) Save(w io.Writer) error {
	err := gob.NewEncoder(w).Encode(*d)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dict) Put(word string) int {
	if id := d.Get(word); id != -1 {
		return id
	}
	wordId := len(d.Data)
	d.Data[word] = wordId
	return wordId
}

func (d *Dict) Get(word string) int {
	i, ok := d.Data[word]
	if !ok {
		return -1
	}
	return i
}

func (d *Dict) Len() int {
	return len(d.Data)
}
