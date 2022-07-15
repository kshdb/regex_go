package main

import (
	"fmt"
	regex "github.com/kshdb/regex_go"
)

func main() {
	//fmt.Println("测试")
	_str := `[{"AAA":"1111","BBB":"222","CCC":""333 444"","DDD":"555"},{"AAA":"1111","BBB":"222","CCC":""333 444"","DDD":"555"},{"AAA":"1111","BBB":"222","CCC":""333 444"","DDD":"555"},{"AAA":"1111","BBB":""333 444"","CCC":"555","DDD":""}]`
	reg := regex.MustCompile(`\"\"`, regex.RE2)
	m, _err := reg.Replace(_str, `"`, 0, -1)
	fmt.Println("执行的结果是", m, _err)

	_srt1 := `
hooudunren.com 
hooeeoooeeeeeen.com 
www.oo234r44.ss.com 
hoo666666.com
	`
	reg1 := regex.MustCompile(`(?<=o{2})([2-6a-z\.]+)(?=\.)`, regex.RE2)
	//sss, _ := reg1.FindStringMatch(_srt1) //获取单个
	sss := regexp2FindAllString(reg1, _srt1) //获取批量
	fmt.Println("执行的结果是", sss)

}

func regexp2FindAllString(re *regex.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
