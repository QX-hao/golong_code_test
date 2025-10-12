package main

import (
	"fmt"
	"strings"
	"unicode"
	// "reflect"
)

// 包含关系
func test_contains() {
	var str1 string = "hello world"
	// 字符串分割，将字符串按空格分割成多个子字符串
	str2 := strings.Split(str1, " ")
	fmt.Println(str2)
	fmt.Println("str2的长度为:", len(str2))

	fmt.Println("-----------------")
	fmt.Println("contains--包含关系")
	//检查包含关系
	//contains判断字符串是否包含指定的子字符串
	fmt.Println("pan qi hao是否包含qi:", strings.Contains("pan qi hao", "qi"))

	//contiansany判断字符串是否包含指定的任意一个字符
	fmt.Println("panqihao是否包含aeiou中的任意一个字符:", strings.ContainsAny("panqihao", "aeiou"))

	//containsrune判断字符串是否包含指定的unicode字符
	fmt.Println("panqihao是否包含q:", strings.ContainsRune("panqihao", 'q'))
}

// 前后缀判断
func test_prefix_suffix() {
	fmt.Println("-----------------")
	fmt.Println("prefix/suffix--前缀/后缀判断")
	fmt.Println("panqihao是否以q开头:", strings.HasPrefix("panqihao", "q"))
	fmt.Println("panqihao是否以hao结尾:", strings.HasSuffix("panqihao", "hao"))
}

// 字符位置查找
func test_index() {
	fmt.Println("-----------------")
	fmt.Println("index/lastindex/IndexAny/LastIndexAny/IndexRune--字符位置查找")

	c := "i am xiao pan"
	//index查找第一次出现位置
	fmt.Println("am第一次出现的位置:", strings.Index(c, "am"))
	//lastindex查找最后一次出现位置
	fmt.Println("xiao最后一次出现的位置:", strings.LastIndex(c, "xiao"))
	//IndexAny查找第一次出现任何指定字符的位置
	fmt.Println("aeiou第一次出现的位置:", strings.IndexAny(c, "aeiou"))
	//LastIndexAny查找最后一次出现任何指定字符的位置
	fmt.Println("aeiou最后一次出现的位置:", strings.LastIndexAny(c, "aeiou"))
	//IndexRune查找第一次出现指定unicode字符的位置
	fmt.Println("o第一次出现的位置:", strings.IndexRune(c, 'o'))
	fmt.Println(c)
}

// 子串出现次数
func test_count() {
	fmt.Println("-----------------")
	fmt.Println("count--计数比较")
	c := "i am xiao pan"
	fmt.Println("i出现的次数:", strings.Count(c, "i"))
	//注意空串的情况
	fmt.Println("there is a girl--空串出现的次数:", strings.Count("there is a girl", ""))
}

// 计数比较
func test_compare() {
	fmt.Println("-----------------")
	fmt.Println("compare--比较")

	//比较ASCALL码值

	// strings.Compare 函数按字符逐个比较两个字符串，返回值规则如下：
	// - 返回 -1 表示第一个字符串小于第二个字符串
	// - 返回 0 表示两个字符串相等
	// - 返回 1 表示第一个字符串大于第二个字符串

	// 示例 1: 比较 "abc" 和 "abe"
	// 比较过程：
	// 字符 'a' == 'a' -> 继续比较下一个字符
	// 字符 'b' == 'b' -> 继续比较下一个字符
	// 字符 'c' < 'e' -> 得出结论 "abc" < "abe"，返回 -1
	// strings.Compare("abc", "abe")
	fmt.Println("abc和abe的比较结果为:", strings.Compare("abc", "abe"))

	// 示例 2: 比较 "abcd" 和 "abe"
	// 比较过程：
	// 前两个字符相同，第三个字符 'c' < 'e'，所以 "abcd" < "abe"，返回 -1
	// 注意：字符串比较不会因为长度更长而认为更大，比较到不同字符时就得出结果。
	fmt.Println("abcd和abe的比较结果为:", strings.Compare("abcd", "abe"))

	// 示例 3: 比较 "abijk" 和 "abe"
	// 比较过程：
	// 前两个字符相同，第三个字符 'i' > 'e'，所以 "abijk" > "abe"，返回 1
	fmt.Println("abijk和abe的比较结果为:", strings.Compare("abijk", "abe"))

	fmt.Println("abe和abe的比较结果为:", strings.Compare("abe", "abe"))

	//不区分大小写比较
	fmt.Println("EqualFold--不区分大小写比较")
	fmt.Println(strings.EqualFold("Go", "go"))       // true
	fmt.Println(strings.EqualFold("Hello", "HELLO")) // true
	fmt.Println(strings.EqualFold("Go", "Python"))   // false
}

// 大小写转换
func caseDemo() {
	fmt.Println("-----------------")
	fmt.Println("ToUpper/ToLower--转换为大小写")

	fmt.Println(strings.ToUpper("Panqihao")) // GOLANG
	fmt.Println(strings.ToLower("PANQIHAO")) // golang
}

// 字符串的修剪
func TrimDemo() {
	fmt.Println("-----------------")
	fmt.Println("字符串的修剪")
	fmt.Println("TrimSpace 去除前后空白字符", strings.TrimSpace("  hello world  ")) // panqihao

	fmt.Println("Trim去除前后指定字符", strings.Trim("!!hello!!world!!", "!"))

	// TrimLeft 去除左侧指定字符
	fmt.Println(strings.TrimLeft("!!!hello!!!", "!")) // "hello!!!"

	// TrimRight 去除右侧指定字符
	fmt.Println(strings.TrimRight("!!!hello!!!", "!")) // "!!!hello"

	// TrimPrefix 去除前缀
	fmt.Println(strings.TrimPrefix("hello world", "hello ")) // "world"

	// TrimSuffix 去除后缀
	fmt.Println(strings.TrimSuffix("hello world", " world")) // "hello"

	//自定义规则
	fmt.Println(strings.TrimFunc(
		"i am pan qi hao 123123", func(r rune) bool {
			return unicode.IsDigit(r)
		}))
}

// 字符串的替换
func replaceDemo() {
	fmt.Println("-----------------")
	fmt.Println("字符串的替换")
	fmt.Println("Replace 替换指定子字符串", strings.Replace("hello world", "world", "Golang", 1)) // hello Golang
	// -1 表示替换所有匹配项
	fmt.Println("Replace 替换所有匹配项", strings.Replace("hello world", "l", "L", -1)) // heLLo worLd

	// 替换所有匹配项---go-1.12+
	fmt.Println("ReplaceAll 替换所有匹配项", strings.ReplaceAll("hello world", "l", "L")) // heLLo worLd

	// Map 替换
	// fmt.Println(
	// 	strings.Map()
	// )
}

// 字符串的切割和连接
func splitAndJoinDemo() {
	fmt.Println("-----------------")
	fmt.Println("字符串的切割")

	// Split 切割字符串
	fmt.Println("Split 切割字符串", strings.Split("a,b,c", ",")) // [a b c]
	// SplitN 切割字符串，指定切割次数
	fmt.Println("SplitN 切割字符串", strings.SplitN("a,b,c", ",", 2)) // [a b,c]
	//SplitAfter 切割字符串，保留分隔符
	fmt.Println("SplitAfter 切割字符串", strings.SplitAfter("a,b,c", ",")) // [a, b, c]
	//Fields 切割字符串，自动去除空白字符
	fmt.Println("Fields 切割字符串", strings.Fields("a b c")) // [a b c]
	//FieldsFunc 切割字符串，根据自定义函数判断分隔符
	// fmt.Println("FieldsFunc 切割字符串", strings.FieldsFunc("a,b,c", func(r rune) bool {
	// 	return r == ','
	// })) // [a b c]

	fmt.Println("-----------------")
	fmt.Println("字符串的连接")
	// Join 连接字符串
	fmt.Println("Join 连接字符串", strings.Join([]string{"a", "b", "c"}, ":")) // a,b,c
	// abc := "sdfghj"
	// fmt.Println("abc的数据类型", reflect.TypeOf(abc))

	// var bbb = strings.Split("i am panqihao", " ")
	// fmt.Println("bbb的数据类型", reflect.TypeOf(bbb))
}

// 重复填充
func repeatDemo() {
	fmt.Println("-----------------")
	fmt.Println("重复填充")
	fmt.Println("Repeat 重复填充", strings.Repeat("abc", 3)) // abcabcabc
}

// 字符串遍历
func rangeDemo() {
	fmt.Println("-----------------")
	fmt.Println("字符串遍历")
	for i, v := range "hello" {
		fmt.Println("索引:", i, "值:", string(v))
	}
}

func main() {
	//包含关系
	test_contains()

	//前后缀判断
	test_prefix_suffix()

	//字符位置查找
	test_index()

	//子串出现次数
	test_count()

	//字符串的比较
	test_compare()

	//大小写转换
	caseDemo()

	//字符串的修剪
	TrimDemo()

	//字符串的替换
	replaceDemo()

	//字符串的切割和连接
	splitAndJoinDemo()

	//重复填充
	repeatDemo()

	//字符串遍历
	rangeDemo()
}
