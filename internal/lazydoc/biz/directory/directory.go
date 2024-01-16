package directory

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"github.com/vinoMamba/lazydoc/pkg/request"
	"github.com/vinoMamba/lazydoc/pkg/token"
)

type DirBiz interface {
	CreateDirBiz(c context.Context, req *request.CreateDirRequest) error
}

type dirBiz struct {
	ds store.IStore
}

var _ DirBiz = (*dirBiz)(nil)

func New(ds store.IStore) *dirBiz {
	return &dirBiz{ds}
}

func (d *dirBiz) CreateDirBiz(c context.Context, req *request.CreateDirRequest) error {
	token := c.Value(known.XUserInfoKey).(*token.TokenInfo)

	var dirM model.DirectoryM

	dirM.CreatedBy = token.ID
	dirM.UpdatedBy = token.ID

	_ = copier.Copy(&dirM, req)

	if err := d.ds.Dirs().CreateDir(c, &dirM); err != nil {
		return err
	}
	return nil
}
