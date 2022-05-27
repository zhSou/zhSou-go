package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/zhSou/zhSou-go/core/model"
	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/initialize"
)

func main() {
	err := initialize.Init()
	defer global.Jieba.Free()
	if err != nil {
		panic(err)
	}
	s := "珈黛"
	res := search(s)
	fmt.Printf("%+v\n", res)
}

func search(text string) []string {
	res := make([]string, 0)
	cut := global.Jieba.CutAll(text)
	for _, v := range cut {
		docIds, err := global.Index.GetAll(v)
		if err != nil {
			if errors.Is(err, model.ErrIndexNotFound) {
				continue
			}
			panic(err)

		}
		for _, docId := range *docIds {
			doc, err := global.Doc.Get(docId)
			if err != nil {
				if errors.Is(err, model.ErrDocNotFound) {
					continue
				}
				panic(err)
			}
			res = append(res, doc)
		}

	}
	return res
}
