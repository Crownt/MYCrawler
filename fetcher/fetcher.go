package fetcher

import (
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io/ioutil"
	"net/http"
	//	"golang.org/x/text/encoding/simplifiedchinese"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	//	"io"
	"log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,
			fmt.Errorf("fff>>>: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	//获得网页编码格式
	bufReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bufReader)

	//使用第三方库将其他编码转化为utf-8编码
	//gopm get -v -g golang.org/x/text  使用其中的encoding/simplifiedchinese/gbk.go
	utf8Reader := transform.NewReader(bufReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader) //直接return返回的 ［］byte, err
}

//自动检查html页面中的编码格式
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	//Peek()用于返回输入流的下n个字节，而不会移动读取位置。返回的[]byte只在下一次调用读取操作前合法。保证下次还可读取这n个字节
	//如果Peek返回的切片长度比n小，它也会返会一个错误说明原因。如果n比缓冲尺寸还大，返回的错误将是ErrBufferFull。
	bytes, err := r.Peek(1024)
	if err != nil {
		//peek不出来，但还是要保证文件可读,返回utf8即可
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}

	//gopm get -v -g golang.org/x/net/html 第三方库，可以自动检查html页面中的编码
	//从传入的1024个字节进行编码判断，返回编码名称，说明和判定结果
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
