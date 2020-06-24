package idea

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"sync"

	//"io/ioutil"
	"net/http"
)

type Idea struct {
	One string
	OneLength int
	Two string
	TwoLength int
	Three string
	ThreeLength int

}

func GetIdeaOne(code *Idea, wg *sync.WaitGroup)   {
	req ,_:=http.Get("http://www.lookdiv.com/code")
	doc,_:=goquery.NewDocumentFromResponse(req)
	text := doc.Find("textarea").Text()
	code.One=text
	code.OneLength= len([]byte(text))
	defer wg.Done()
	return
}
func IdeaList(code *Idea, wg *sync.WaitGroup)  {
	req ,_:=http.Get("https://shimo.im/docs/XvW3WpHgHdRHVXgV/read")
	doc,_:=goquery.NewDocumentFromResponse(req)
	doc.Find(".ql-sheet-cell p").Each(func(i int, selection *goquery.Selection) {
		if i==0 {
			code.Two = selection.Text()
			code.TwoLength= len([]byte(selection.Text()))
		}
		if i==1 {
			code.Three=selection.Text()
			code.ThreeLength= len([]byte(selection.Text()))
		}
	})
	defer wg.Done()

}
func GetIdeaTwo(code *Idea, wg *sync.WaitGroup)   {
	req,_:=http.Get("http://www.yanjie.site/jjj2.txt")
	byte,_:=ioutil.ReadAll(req.Body)
	codeString:=string(byte)
	code.Two=codeString
	defer wg.Done()
	return
}

func GetAllIdea()  *Idea  {
	var wg sync.WaitGroup
	var idea Idea
	wg.Add(2)
	//go GetIdeaTwo(&idea,&wg)
	go GetIdeaOne(&idea,&wg)
	go IdeaList(&idea,&wg)
	wg.Wait()
	return &idea

}
