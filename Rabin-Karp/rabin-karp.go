package main

import (
	"github.com/Chain-Zhang/pinyin"
	"log"
)

type HashUtils struct {
}

func (hashUtils *HashUtils) removeAllSpace(str *string) string {
	ret := ""
	for _, v := range *str {
		if v != ' ' {
			ret += string(v)
		}
	}
	return ret
}

func (hashUtils *HashUtils) chineseToChar(s string) string {
	str, err := pinyin.New(s).Convert()
	if err != nil {
		// 错误处理
		log.Panicln("中文转英文字符串出错")
	} else {
		hashUtils.removeAllSpace(&str)
	}
	return str
}

func (hashUtils *HashUtils) rabinKarp(s string) int {
	base := 131
	p := []int{1}
	hash := []int{0}
	for i := 0; i < len(s); i++ {
		hash = append(hash, hash[len(hash)-1]*base+int(s[i]-'a')+1)
		p = append(p, p[len(p)-1]*base)
	}
	return hash[len(hash)-1] - hash[0]
}

//func main()  {
//
//	str, err := pinyin.New("  我是中国人").Convert()
//	if err != nil {
//		// 错误处理
//	}else{
//		fmt.Println(str)
//	}
//	fmt.Println(removeAllSpace(&str))
//}

//func main() {
//	fmt.Println(rabinKarp("ab"))
//}
