package Indexer

//package main
import (
	"math"
)

type Indexer struct {
	Docs          [][]string
	WordsMap      []string
	wordSplitArr  [][]string
	happenWord    []float64
	words         []string
	frequency     [][]float64
	b             float64
	k1            float64
	Bm25Value     [][][]float64
	totalNumWords int
	idf           []float64
	avgWordNum    float64
}

func (indexer *Indexer) getMap(arrTmp [][]string) {
	for _, value := range arrTmp {
		//arrTmp := strings.Split(v, "|")
		indexer.wordSplitArr = append(indexer.wordSplitArr, value)
		indexer.words = append(indexer.words, value...)
		indexer.totalNumWords += len(arrTmp)
	}

	//去用map重
	tmpMap := make(map[string]int)
	for _, value := range indexer.words {
		tmpMap[value] = tmpMap[value] + 1
	}
	indexer.WordsMap = []string{}
	for key, _ := range tmpMap {
		indexer.WordsMap = append(indexer.WordsMap, key)
	}
}

func (indexer *Indexer) getFrequency() {

	//切分后的word数量
	docNum := len(indexer.wordSplitArr)
	wordNum := len(indexer.WordsMap)
	//初始化
	indexer.frequency = make([][]float64, wordNum)
	for i := 0; i < len(indexer.frequency); i++ {
		indexer.frequency[i] = make([]float64, docNum)
	}

	idx := 0
	//出现该分词词的文档数目
	indexer.happenWord = make([]float64, wordNum)
	//计算第i个word出现在第j个文档的词频
	for _, word := range indexer.WordsMap {
		for j, arr := range indexer.wordSplitArr { //wordSplitArr[i]存储着第i个句子分词后的结果
			isWriteHappenWord := false
			for _, splitWord := range arr {
				if word == splitWord {
					if !isWriteHappenWord {
						isWriteHappenWord = true
						indexer.happenWord[idx]++
					}
					//第i个word出现在第j个文档的词频
					indexer.frequency[idx][j]++
				}
			}
		}
		idx++
	}
}

func (indexer *Indexer) getIDF() {
	//切分后的word数量
	wordNum := len(indexer.WordsMap)
	docNum := float64(len(indexer.wordSplitArr))
	indexer.idf = make([]float64, wordNum)
	for i := 0; i < wordNum; i++ {
		if indexer.happenWord[i] == 0 {
			indexer.idf[i] = 0
		} else {
			indexer.idf[i] = math.Log2((docNum / indexer.happenWord[i]) + 1.0)
		}

	}
}

func (indexer *Indexer) InitBM25Param(strs, input [][]string) {
	//indexer.WordsMap = make(map[string]int)
	indexer.Docs = input
	indexer.getMap(strs)
	indexer.getFrequency()
	indexer.getIDF()
	indexer.b = 0.75
	indexer.k1 = 2.0
	indexer.avgWordNum = float64(indexer.totalNumWords / len(indexer.wordSplitArr))
	indexer.Bm25Value = make([][][]float64, len(indexer.WordsMap))
	for i := 0; i < len(indexer.Bm25Value); i++ {
		indexer.Bm25Value[i] = make([][]float64, len(indexer.wordSplitArr))
		for j := 0; j < len(indexer.wordSplitArr); j++ {
			indexer.Bm25Value[i][j] = make([]float64, 3)
		}
	}
}

func (indexer *Indexer) CalcBM25(strs, input [][]string) {
	indexer.InitBM25Param(strs, input)
	wordNum := len(indexer.WordsMap)
	docNum := len(indexer.wordSplitArr)
	for i := 0; i < wordNum; i++ {
		for j := 0; j < docNum; j++ {
			indexer.Bm25Value[i][j][0] = float64(i)
			indexer.Bm25Value[i][j][1] = float64(j)
			indexer.Bm25Value[i][j][2] = (indexer.idf[i] * indexer.frequency[i][j] * (indexer.k1 + 1)) /
				(indexer.frequency[i][j] + indexer.k1*(1-indexer.b+(indexer.b*float64(len(indexer.wordSplitArr[j]))/indexer.avgWordNum)))
		}
	}

}

//func main() {
//
//
//	input := excel.GetExcelData("data/wukong50k_release.xlsx")[1:100]
//	fmt.Println(input)
//	segmentWords := Segment.Segmenter(input)
//	fmt.Println(segmentWords)
//	bm25 := &BM25{}
//
//	CalcBM25(bm25, segmentWords, input)
//
//	fmt.Println(222)
//
//}
