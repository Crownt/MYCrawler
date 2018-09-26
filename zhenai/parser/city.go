package parser

import (
	"crownt.org/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

//通过正则表达式对city页面进行解析
func ParserCity(content []byte) engine.ParseResult {
	result := engine.ParseResult{}

	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	for _, match := range matches {
		//？？？？？？
		name := string(match[2])
		result.DataItem = append(result.DataItem, string(match[2]))
		result.Requests = append(result.Requests, engine.Request{string(match[1]), func(content []byte) engine.ParseResult { return ParserProfile(content, name) }})
	}
	return result
}
