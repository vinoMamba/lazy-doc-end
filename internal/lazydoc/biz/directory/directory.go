package directory

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"github.com/vinoMamba/lazydoc/internal/lazydoc/store"
	"github.com/vinoMamba/lazydoc/internal/pkg/errno"
	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"github.com/vinoMamba/lazydoc/internal/pkg/model"
	"github.com/vinoMamba/lazydoc/pkg/request"
	"github.com/vinoMamba/lazydoc/pkg/token"
)

type DirBiz interface {
	CreateDirBiz(c context.Context, req *request.CreateDirRequest) error
	UpdateDirBiz(c context.Context, id int64, req *request.UpdateDirRequest) error
	DeleteDirBiz(c context.Context, id int64) error
	GetDirListBiz(c context.Context) ([]*request.DirListRequest, error)
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

	dirM.ParentId = cast.ToInt64(req.ParentId)

	if err := d.ds.Dirs().CreateDir(c, &dirM); err != nil {
		return err
	}
	return nil
}

func (d *dirBiz) UpdateDirBiz(c context.Context, id int64, req *request.UpdateDirRequest) error {

	current, err := d.ds.Dirs().GetDirById(c, id)

	if err != nil {
		return errno.ErrDirNotFound
	}

	token := c.Value(known.XUserInfoKey).(*token.TokenInfo)

	_ = copier.Copy(&current, req)
	current.UpdatedBy = token.ID

	if err := d.ds.Dirs().UpdateDir(c, current); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (d *dirBiz) DeleteDirBiz(c context.Context, id int64) error {
	current, err := d.ds.Dirs().GetDirById(c, id)

	if err != nil {
		return errno.ErrDirNotFound
	}

	current.IsDeleted = 1

	if err := d.ds.Dirs().UpdateDir(c, current); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (d *dirBiz) GetDirListBiz(c context.Context) ([]*request.DirListRequest, error) {
	token := c.Value(known.XUserInfoKey).(*token.TokenInfo)

	var dirList []*request.DirListRequest

	l1, err := d.ds.Dirs().GetDirListByUserId(c, token.ID)
	if err != nil {
		return nil, err
	}
	l2, err := d.ds.Dirs().GetDirListByStatus(c, 0)
	if err != nil {
		return nil, err
	}

	for _, v := range l1 {
		var dir request.DirListRequest
		if v.IsPublic == 1 {
			_ = copier.Copy(&dir, v)
			dir.Id = cast.ToString(v.ID)
			dir.ParentId = cast.ToString(v.ParentId)
			dirList = append(dirList, &dir)
		}
	}
	for _, v := range l2 {
		var dir request.DirListRequest
		_ = copier.Copy(&dir, v)
		dir.Id = cast.ToString(v.ID)
		dir.ParentId = cast.ToString(v.ParentId)
		dirList = append(dirList, &dir)
	}

	return dirList, nil
}
