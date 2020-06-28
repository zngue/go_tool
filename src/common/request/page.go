package request

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
func PageLimitOffset( page int ,pagzeSize int ) int {
	return (page-1)*pagzeSize
}
func (p *Page) PageLimitOffset ()  int {
	return  (p.Page-1)*p.PageSize
}