package parser

import (
	"crownt.org/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//通过正则表达式，找出城市的url
func ParseCityList(content []byte) engine.ParseResult {
	result := engine.ParseResult{}

	//编译正则表达式
	re := regexp.MustCompile(cityListRe)
	//用编译后的正则表达式进行匹配
	matches := re.FindAllSubmatch(content, -1)

	limit := 10
	for _, match := range matches {
		result.Requests = append(result.Requests,
			engine.Request{Url: string(match[1]),
				ParserFunc: ParserCity,
			})
		result.DataItem = append(result.DataItem, string(match[2]))

		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
