package invertedindex

import (
	"encoding/gob"
	"github.com/zhSou/zhSou-go/util/algorithm/set"
	"io"
	"sort"
)

type invertedIndex struct {
	Data map[string][]int
}

func NewInvertedIndex() *invertedIndex {
	return &invertedIndex{
		Data: make(map[string][]int),
	}
}

func LoadInvertedIndexFromDisk(r io.Reader) (*invertedIndex, error) {
	ii := invertedIndex{}
	err := gob.NewDecoder(r).Decode(&ii)
	if err != nil {
		return nil, err
	}
	return &ii, nil
}

func (i *invertedIndex) SaveToDisk(w io.Writer) error {
	err := gob.NewEncoder(w).Encode(*i)
	if err != nil {
		return err
	}
	return nil
}

func (i *invertedIndex) Add(word string, id int) {
	i.Data[word] = append(i.Data[word], id)
}

func (i *invertedIndex) AddWords(words []string, id int) {
	for _, word := range set.Deduplication[string](words) {
		i.Data[word] = append(i.Data[word], id)
	}
}

func (i *invertedIndex) Get(word string) []int {
	return i.Data[word]
}

func (i *invertedIndex) Sort() {
	for _, ids := range i.Data {
		sort.Ints(ids)
	}
}
