package main

import (
	"strings"

	"github.com/gramework/gramework"
)

var counters = map[string]uint{}

func main() {
	app := gramework.New()

	app.GET("/*any", func(ctx *gramework.Context) interface{} {
		url := ctx.UserValue("any").(string)
		if url != "/" {
			url = strings.TrimRight(url, "/")
		}
		counters[url]++
		//ctx.Writeln(url)
		return counters
	})

	app.ListenAndServe()

}
