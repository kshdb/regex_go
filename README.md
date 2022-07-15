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
