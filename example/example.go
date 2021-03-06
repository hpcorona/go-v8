package main

import (
	"fmt"
	"github.com/hpcorona/go-v8/v8"
)

func main() {
	v8ctx := v8.NewContext()

	// setup console.log()
	v8ctx.Eval(`
	this.console = { "log": function(args) { _console_log(args) }}`)
	v8ctx.AddFunc("_console_log", func(args... interface{}) interface{} {
		for _, arg := range args {
			fmt.Printf("%v ", arg)
		}
		fmt.Println()
		return ""
	})

	ret, _ := v8ctx.Eval(`
	var a = 1;
	var b = 'B'
	a += 2;
	a;
	`)
	println(int(ret.(float64))) // 3

	v8ctx.Eval(`console.log(a + '年' + b + '組 金八先生！')`)  // 3b
	v8ctx.Eval(`console.log("Hello World, こんにちわ世界")`) // john
}
