package response

import "github.com/vinoMamba/lazy-doc-end/storage"

type ProjectList struct {
	Total int               `json:"total"`
	List  []storage.Project `json:"list"`
}

type Project struct {
}
