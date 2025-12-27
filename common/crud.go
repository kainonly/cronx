package common

import "fmt"

type FindDto struct {
	PageSize int64    `header:"x-pagesize" vd:"omitempty,min=0,max=1000"`
	Page     int64    `header:"x-page" vd:"omitempty,min=0"`
	Q        string   `query:"q,omitempty"`
	Sort     []string `query:"sort,omitempty" vd:"omitempty,dive,sort"`
}

func (x *FindDto) GetPageSize() int {
	if x.PageSize == 0 {
		x.PageSize = 1000
	}
	return int(x.PageSize)
}

func (x *FindDto) GetOffset() int {
	return int(x.Page) * int(x.PageSize)
}

func (x *FindDto) GetKeyword() string {
	return fmt.Sprintf(`%%%s%%`, x.Q)
}
