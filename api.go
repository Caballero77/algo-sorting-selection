package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/selection", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/selection", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
}

func sort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		index := i
		for j := i; j < length; j++ {
			if array[index] > array[j] {
				index = j
			}
		}
		buf := array[index]
		array[index] = array[i]
		array[i] = buf
	}
	return array
}
