package Query

import (
	"mySearch/Indexer"
	"mySearch/Segment"
	"mySearch/Sorter"
	"strings"
)

var (
	bm25      = &Indexer.Indexer{}
	sorter    = &Sorter.Sorter{}
	segmenter = &Segment.Segmenter{}
)

type Query struct {
}

func (query *Query) getQueryRet(querys string) []string {
	querysArr := segmenter.QuerySegmenter(querys)
	idAndScoreArr := [][]float64{}
	for _, queryWord := range querysArr {
		idAndScoreArr = append(idAndScoreArr, query.findIdAndScore(queryWord)...)
	}
	//按bm25排序由高到低排序
	idAndScoreArr = sorter.Sorter(idAndScoreArr)
	docs := []string{}
	for _, idAndScore := range idAndScoreArr {
		wordId := int(idAndScore[0])
		docId := int(idAndScore[1])
		docs = append(docs, strings.Replace(bm25.Docs[docId][1], bm25.WordsMap[wordId], "<font color=red>"+bm25.WordsMap[wordId]+"</font>", -1))
	}
	return docs
}

func (query *Query) findIdAndScore(queryWord string) [][]float64 {
	wordMapId := 0
	//找出query这个词对应于bm25.wordMapId里的index
	for i := 0; i < len(bm25.WordsMap); i++ {
		if bm25.WordsMap[i] == queryWord {
			wordMapId = i
			break
		}
	}
	idAndScoreArr := [][]float64{}
	for _, DocIdAndScore := range bm25.Bm25Value[wordMapId] {
		if DocIdAndScore[2] != 0 {
			idAndScoreArr = append(idAndScoreArr, []float64{DocIdAndScore[0], DocIdAndScore[1], DocIdAndScore[1]})
		}
	}
	return idAndScoreArr
}
