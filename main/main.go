package main

import (
	"crownt.org/crawler/engine"
	//	"crownt.org/crawler/fetcher"
	"crownt.org/crawler/zhenai/parser"
	//	"fmt"
)

func main() {
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})

	//	content, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//
	//	fmt.Printf("%s", content)
}
