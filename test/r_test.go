package test

import (
	"archive/zip"
	"fmt"
	"sync"
	"testing"
)

type HttpResponseMap map[string]interface{}


func (HttpResponseMap)Abc()  {

	fmt.Println("3366")
}
func TestMp(t *testing.T) {

	url :="http://idea.medeming.com/jets/images/jihuoma.zip"
	rc, _ := zip.OpenReader(url)
	fmt.Println(rc.File)



}
func abc(key string,abc HttpResponseMap,wg *sync.WaitGroup)  {

	defer  wg.Done()
	abc[key]="hangar"


}
