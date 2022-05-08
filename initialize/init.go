package initialize

import "github.com/pkg/errors"

//
// Init
//  @Description:
//
func Init() error {
	InitJieba()

	err := InitDoc()
	if err != nil {
		return errors.WithMessage(err, "InitDoc failed")
	}
	err = InitIndex()
	if err != nil {
		return errors.WithMessage(err, "InitIndex failed")
	}
	return nil
}
