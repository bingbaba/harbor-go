package harbor

import (
	"fmt"
	"net/url"
)

type PageOption struct {
	Page     int32
	PageSize int32
}

func (opt PageOption) Urls() (v url.Values) {
	v = make(map[string][]string)
	if opt.Page > 1 {
		v.Set("page", fmt.Sprintf("%d", opt.Page))
	}
	if opt.PageSize > 0 {
		v.Set("page_size", fmt.Sprintf("%d", opt.PageSize))
	}
	return v
}
