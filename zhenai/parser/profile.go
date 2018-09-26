package parser

import (
	"crownt.org/crawler/engine"
	"crownt.org/crawler/model"
	"regexp"
	"strconv"
)

//在外面对正则表达式进行预先编译，提高代码效率
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(` <td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(` <td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(` <td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(` <td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(` <td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(` <td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`  <td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

//由于在城市解析器中已经获得了人的名字，这里的解析器直接多传入一个名字参数，
func ParserProfile(content []byte, name string) engine.ParseResult {
	result := engine.ParseResult{}
	profile := model.Profile{}

	profile.Name = name
	age, err := strconv.Atoi(extractString(content, ageRe))
	if err == nil {
		profile.Age = age
	}

	gender := extractString(content, genderRe)
	profile.Gender = gender

	height, err := strconv.Atoi(extractString(content, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(content, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	income := extractString(content, incomeRe)
	profile.Income = income

	education := extractString(content, educationRe)
	profile.Education = education

	occupation := extractString(content, occupationRe)
	profile.Occupation = occupation

	hokou := extractString(content, hokouRe)
	profile.Hokou = hokou

	xingzuo := extractString(content, xingzuoRe)
	profile.Xinzuo = xingzuo

	house := extractString(content, houseRe)
	profile.House = house

	car := extractString(content, carRe)
	profile.Car = car

	result.DataItem = append(result.DataItem, profile)
	return result
}

//通过不同的正则表达式，对内容进行抽取
func extractString(content []byte, re *regexp.Regexp) string {
	//FindSubmatch()　只查找第一个匹配的内容
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
