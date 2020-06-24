package request

type Page struct {
	Page int `form:"page" json:"page"` //当前页码
	PageSize int `form:"page_size" json:"page_size" `//分页大小
	OrderFile []OrderFile
}
type OrderFile struct {
	Field string `json:"field" form:"field"`
	Sort string `json:"sort" form:"sort"`
}
func PageLimitOffset( page int ,pagzeSize int ) int {
	return (page-1)*pagzeSize
}