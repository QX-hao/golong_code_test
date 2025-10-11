package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "hello world"
	// 字符串分割，将字符串按空格分割成多个子字符串
	str2 := strings.Split(str1," ")
	fmt.Println(str2)
	fmt.Println("str2的长度为:",len(str2))


	fmt.Println("-----------------")
	fmt.Println("contains--包含关系")
	//检查包含关系
	//contains判断字符串是否包含指定的子字符串
	fmt.Println("pan qi hao是否包含qi:",strings.Contains("pan qi hao","qi"))

	//contiansany判断字符串是否包含指定的任意一个字符
	fmt.Println("panqihao是否包含aeiou中的任意一个字符:",strings.ContainsAny("panqihao","aeiou"))

	//containsrune判断字符串是否包含指定的unicode字符
	fmt.Println("panqihao是否包含q:",strings.ContainsRune("panqihao",'q'))

	//前后缀判断
	fmt.Println("-----------------")
	fmt.Println("prefix/suffix--前缀/后缀判断")
	fmt.Println("panqihao是否以q开头:",strings.HasPrefix("panqihao","q"))
	fmt.Println("panqihao是否以hao结尾:",strings.HasSuffix("panqihao","hao"))
	
	//字符位置查找
	fmt.Println("-----------------")
	fmt.Println("index/lastindex/IndexAny/LastIndexAny/IndexRune--字符位置查找")
	
	c := "i am xiao pan"
	//index查找第一次出现位置
	fmt.Println("am第一次出现的位置:",strings.Index(c,"am"))
	//lastindex查找最后一次出现位置
	fmt.Println("xiao最后一次出现的位置:",strings.LastIndex(c,"xiao"))
	//IndexAny查找第一次出现任何指定字符的位置
	fmt.Println("aeiou第一次出现的位置:",strings.IndexAny(c,"aeiou"))
	//LastIndexAny查找最后一次出现任何指定字符的位置
	fmt.Println("aeiou最后一次出现的位置:",strings.LastIndexAny(c,"aeiou"))
	//IndexRune查找第一次出现指定unicode字符的位置
	fmt.Println("o第一次出现的位置:",strings.IndexRune(c,'o'))

	//计数比较
	fmt.Println("i出现的次数:",strings.Count(c,"i"))


	

	fmt.Println(c)



}