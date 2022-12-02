package Segment

//package main
import (
	"github.com/yanyiwu/gojieba"
)

type Segmenter struct {
}

func (segmenter *Segmenter) Segmenter(sentences [][]string) [][]string {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	//var useHmm = true

	var resWords []string
	ret := [][]string{}
	//var sentence = "计算机视觉"

	for i := 0; i < len(sentences); i++ {
		resWords = seg.CutAll(sentences[i][1])
		ret = append(ret, resWords)
	}

	return ret
}

func (segmenter *Segmenter) QuerySegmenter(querys string) []string {
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
//	QuerySegmenter("中国足球")
//}
