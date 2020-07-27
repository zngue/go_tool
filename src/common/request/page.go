package request

import "github.com/jinzhu/gorm"

type Page struct {
	Page int `form:"page" json:"page"` //当前页码
	IsCount int `form:"is_count" json:"is_count"`
	PageSize int `form:"page_size" json:"page_size" `//分页大小
	OrderFile []OrderFile `json:"order_file" form:"order_file"`
}
type OrderFile struct {
	Field string `json:"field" form:"field"`
	Sort string `json:"sort" form:"sort"`
}
type ListCount func() error
func (p *Page) PageLimitOffset ()  int {
	return  (p.Page-1)*p.PageSize
}

func (p *Page) OnlyPage (db *gorm.DB) *gorm.DB {
	if p.PageSize>0 &&p.Page>0 {
		db = db.Offset(p.PageLimitOffset()).Limit(p.PageSize)
	}
	return db
}
func (p *Page) OnlyCount(count ListCount) (bool,error) {
	var typeInfo bool
	if count!=nil && p.IsCount!=0 {
		err:=count()
		if  err!=nil{
			return  typeInfo,err
		}
		if p.IsCount==2 {
			return typeInfo,nil
		}
	}
	return typeInfo,nil
}

