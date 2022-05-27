package initialize

import (
	"github.com/pkg/errors"

	"github.com/zhSou/zhSou-go/core/model"
	"github.com/zhSou/zhSou-go/global"
)

func InitIndex() error {
	global.Index = model.NewIndex()

	keys, values, err := global.Doc.GetAll()
	if err != nil {
		return errors.Wrap(err, "get all doc failed")
	}
	for i := 0; i < len(keys); i++ {
		words := global.Jieba.CutAll(values[i])
		for j := 0; j < len(words); j++ {
			err := global.Index.Add(words[j], keys[i])
			if err != nil && !errors.Is(err, model.ErrIndexExists) {
				return errors.Wrapf(err, "add index failed, word: %s, doc: %s", words[j], values[i])
			}
		}
	}
	return nil
}
