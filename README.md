# regex_go
golang与Per5和.net兼容的正则处理包

#### 比较regexp和regex_go

|Category |regexp |regex_go|
| --- | --- | --- |
|Catastrophic backtracking possible	|no, constant execution time guarantees	|yes, if your pattern is at risk you can use the <font color="#ff0000">re.MatchTimeout</font> field|
|Python-style capture groups <font color="#ff0000">(?P<name>re)</font>|	yes	|no (yes in RE2 compat mode)
|.NET-style capture groups <font color="#ff0000">(?<name>re)</font> or <font color="#ff0000">(?'name're)</font>|	no|	yes|
|comments <font color="#ff0000">(?#comment)</font>|	no|	yes|
|branch numbering reset <font color="#ff0000">(?\|a\|b)</font>	|no	|no|
|possessive match <font color="#ff0000">(?>re)</font> |	no|	yes|
|positive lookahead <font color="#ff0000">(?=re)</font>|	no|	yes|
|negative lookahead <font color="#ff0000">(?!re)</font>|	no|	yes|
|positive lookbehind <font color="#ff0000">(?<=re)</font>|	no|	yes|
|negative lookbehind <font color="#ff0000">(?<!re)</font>|	no|	yes|
|back reference <font color="#ff0000">\1</font>|	no|	yes|
|named back reference <font color="#ff0000">\k'name'</font>|	no|	yes|
|named ascii character class <font color="#ff0000">[[:foo:]]</font>|	yes|	no (yes in RE2 compat mode)|
|conditionals <font color="#ff0000">(?(expr)yes\|no)</font>|	no|	yes|

#### 使用例子：
```go
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
hooudcnwtn.com 
hooeeoooeeeeeen.com 
www.oo234r44.ss.com 
hoo666666.com
	`
	reg1 := regex.MustCompile(`(?<=o{2})([2-6a-z\.]+)(?=\.)`, regex.RE2)
	//sss, _ := reg1.FindStringMatch(_srt1) //获取单个
	sss := regexp2FindAllString(reg1, _srt1) //获取批量
	fmt.Println("执行的结果是", sss)

}
/*
获取批量
*/
func regexp2FindAllString(re *regex.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

```