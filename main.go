package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"syscall/js"
)

var (
	global js.Value
)

func init() {
	global = js.Global()
}

func main() {
	global.Set("getData", js.FuncOf(getData))

	select {}
}

func getData(_ js.Value, args []js.Value) interface{} {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		r, _ := http.Get("https://rickandmortyapi.com/api/character/")
		defer r.Body.Close()

		characters, _ := ioutil.ReadAll(r.Body)

		fmt.Println(fmt.Sprintf("%s", characters))
	}()
	wg.Wait()

	return nil
}
