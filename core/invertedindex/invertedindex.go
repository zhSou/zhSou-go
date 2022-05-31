package invertedindex

import (
	"encoding/gob"
	"io"
	"sort"
	"sync"

	"github.com/zhSou/zhSou-go/core/dict"
	"github.com/zhSou/zhSou-go/util/algorithm/set"
)

type InvertedIndex struct {
	mutex sync.RWMutex
	Data  map[int][]int
	dict  *dict.Dict
}

func NewInvertedIndex(dict *dict.Dict) *InvertedIndex {
	return &InvertedIndex{
		Data: make(map[int][]int),
		dict: dict,
	}
}

func LoadInvertedIndexFromDisk(r io.Reader, dict *dict.Dict) (*InvertedIndex, error) {
	ii := InvertedIndex{}
	err := gob.NewDecoder(r).Decode(&ii)
	if err != nil {
		return nil, err
	}
	ii.dict = dict
	return &ii, nil
}

func (i *InvertedIndex) SaveToDisk(w io.Writer) error {
	err := gob.NewEncoder(w).Encode(*i)
	if err != nil {
		return err
	}
	return nil
}

func (i *InvertedIndex) Add(word string, id int) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	wordId := i.dict.Put(word)
	i.Data[wordId] = append(i.Data[wordId], id)
}

func (i *InvertedIndex) AddWords(words []string, id int) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	for _, word := range set.Deduplication[string](words) {
		wordId := i.dict.Put(word)
		i.Data[wordId] = append(i.Data[wordId], id)
	}
}

func (i *InvertedIndex) Get(word string) []int {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	wordId := i.dict.Put(word)
	return i.Data[wordId]
}

func (i *InvertedIndex) Sort() {
	for _, ids := range i.Data {
		sort.Ints(ids)
	}
}
