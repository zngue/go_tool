package helper

import (
	"errors"
	"github.com/zngue/go_tool/src/fun/array"
	"github.com/zngue/go_tool/src/fun/zng_str"
	"sync"
)
var (
	AddNum = 1
	ListNum =2
	UpdateNum=3
	DeleteNum=4
	StatusNum=5
	DetailNum=6
	ListMoreNum=7
)
type Model struct {
	TypeExt string
	IsGroupRun bool
	Add *Helper
	AddMore []*Helper
	List *Helper
	ListMore []*Helper
	Update *Helper
	Delete *Helper
	Status *Helper
	Detail  *Helper
	ResponseErr *ResponseErr
	Err error
}
type ResponseErr struct {
	AddErr error
	ListErr error
	DeleteErr error
	StatusErr error
	UpdateErr error
	AllErr bool
}

func ListMore(helper ...*Helper) []*Helper  {
	return helper
}
func (m *Model) Run() *Model  {
	if m.TypeExt=="" {
		m.Err=errors.New("typpe ext is should")
		return m
	}
	if m.IsGroupRun {
		return m.GroupRun()
	}else {
		return  m.GoRun()
	}
}
func (m *Model) GroupRun() *Model {
	var wg sync.WaitGroup

	idsArr := zng_str.IDStringToSlice(m.TypeExt,",")
	if &m.Add!=nil && array.IntIsArray(AddNum,idsArr) {
		wg.Add(1)
		go m.Add.GoRun(&wg)
	}
	if &m.List!=nil  && array.IntIsArray(ListNum,idsArr) {
		wg.Add(1)
		go m.List.GoRun(&wg)
	}
	if len(m.ListMore)>0 && array.IntIsArray(ListMoreNum,idsArr) {
		for _,fn:=range m.ListMore{
			wg.Add(1)
			go fn.GoRun(&wg)
		}
	}
	if &m.Status!=nil &&  array.IntIsArray(StatusNum,idsArr) {
		wg.Add(1)
		go m.Status.GoRun(&wg)
	}
	if &m.Update!=nil &&  array.IntIsArray(UpdateNum,idsArr) {
		wg.Add(1)
		go m.Update.GoRun(&wg)
	}
	if &m.Delete!=nil && array.IntIsArray(DeleteNum,idsArr) {
		wg.Add(1)
		go m.Delete.GoRun(&wg)
	}
	if &m.Detail!=nil && array.IntIsArray(DetailNum,idsArr)  {
		wg.Add(1)
		go m.Detail.GoRun(&wg)
	}
	wg.Wait()
	return m
}
func (m *Model) GoRun() *Model {
	idsArr := zng_str.IDStringToSlice(m.TypeExt,",")
	if &m.Add!=nil && array.IntIsArray(AddNum,idsArr) {
		 m.Add.Run()
	}
	if &m.List!=nil && array.IntIsArray(ListNum,idsArr) {
		 m.List.Run()
	}
	if &m.Status!=nil && array.IntIsArray(StatusNum,idsArr)  {
		 m.Status.Run()
	}
	if &m.Update!=nil && array.IntIsArray(UpdateNum,idsArr)  {
		  m.Update.Run()
	}
	if &m.Delete!=nil && array.IntIsArray(DetailNum,idsArr)  {
		 m.Delete.Run()
	}
	if &m.Detail!=nil && array.IntIsArray(DetailNum,idsArr)  {
		m.Detail.Run()
	}
	return m
}
