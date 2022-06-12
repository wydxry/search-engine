package Segment

//package main
import (
	"github.com/yanyiwu/gojieba"
)

func Segmenter(sentences [][]string) [][]string {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	//var useHmm = true

	var resWords []string
	ret := [][]string{}
	//var sentence = "万里长城万里长"

	for i := 0; i < len(sentences); i++ {
		resWords = seg.CutAll(sentences[i][1])
		//resWords = seg.Cut(sentences[i][1], useHmm)
		ret = append(ret, resWords)
	}

	return ret
}

func QuerySegmenter(querys string) []string {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	//var useHmm = true

	var resWords []string
	ret := []string{}

	resWords = seg.CutAll(querys)
	ret = append(ret, resWords...)
	return ret
}

//func main() {
//	QuerySegmenter("清华大学")
//}
