package engine

import (
	"crownt.org/crawler/fetcher"
	//	"crownt.org/crawler/zhenai/parser"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	//engine需要维护一个Request的队列
	var Q []Request
	for _, seed := range seeds {
		Q = append(Q, seed)
	}

	for len(Q) > 0 {
		r := Q[0]
		Q = Q[1:]

		//对每个request的url页面进行fetch
		body, err := fetcher.Fetch(r.Url)
		log.Printf("fetch url:%s", r.Url)
		if err != nil {
			log.Printf("Fetch error: fetch url %s: %v", r.Url, err)
			continue
			//return
		}
		//对fetch到的内容进行parser
		results := r.ParserFunc(body)
		Q = append(Q, results.Requests...)
		for _, dataItem := range results.DataItem {
			fmt.Printf("data: %v\n", dataItem)
		}
		//		for _, request := range results.Requests {
		//			Q = append(Q, request)
		//		}
	}
}
