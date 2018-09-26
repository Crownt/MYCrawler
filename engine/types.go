package engine

import ()

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	DataItem []interface{}
	Requests []Request
}

func NilParserFunc([]byte) ParseResult {
	return ParseResult{}
}
