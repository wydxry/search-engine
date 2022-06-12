package Indexer

//package main
import (
	"math"
)

type BM25 struct {
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

func getMap(bm25 *BM25, arrTmp [][]string) {
	for _, value := range arrTmp {
		//arrTmp := strings.Split(v, "|")
		bm25.wordSplitArr = append(bm25.wordSplitArr, value)
		bm25.words = append(bm25.words, value...)
		bm25.totalNumWords += len(arrTmp)
	}

	//去用map重
	tmpMap := make(map[string]int)
	for _, value := range bm25.words {
		tmpMap[value] = tmpMap[value] + 1
	}
	bm25.WordsMap = []string{}
	for key, _ := range tmpMap {
		bm25.WordsMap = append(bm25.WordsMap, key)
	}
}

func getFrequency(bm25 *BM25) {

	//切分后的word数量
	docNum := len(bm25.wordSplitArr)
	wordNum := len(bm25.WordsMap)
	//初始化
	bm25.frequency = make([][]float64, wordNum)
	for i := 0; i < len(bm25.frequency); i++ {
		bm25.frequency[i] = make([]float64, docNum)
	}

	idx := 0
	//出现该分词词的文档数目
	bm25.happenWord = make([]float64, wordNum)
	//计算第i个word出现在第j个文档的词频
	for _, word := range bm25.WordsMap {
		for j, arr := range bm25.wordSplitArr { //wordSplitArr[i]存储着第i个句子分词后的结果
			isWriteHappenWord := false
			for _, splitWord := range arr {
				if word == splitWord {
					if !isWriteHappenWord {
						isWriteHappenWord = true
						bm25.happenWord[idx]++
					}
					//第i个word出现在第j个文档的词频
					bm25.frequency[idx][j]++
				}
			}
		}
		idx++
	}
}

func getIDF(bm25 *BM25) {
	//切分后的word数量
	wordNum := len(bm25.WordsMap)
	docNum := float64(len(bm25.wordSplitArr))
	bm25.idf = make([]float64, wordNum)
	for i := 0; i < wordNum; i++ {
		if bm25.happenWord[i] == 0 {
			bm25.idf[i] = 0
		} else {
			bm25.idf[i] = math.Log2((docNum / bm25.happenWord[i]) + 1.0)
		}

	}
}

func InitBM25Param(bm25 *BM25, strs, input [][]string) {
	//bm25.WordsMap = make(map[string]int)
	bm25.Docs = input
	getMap(bm25, strs)
	getFrequency(bm25)
	getIDF(bm25)
	bm25.b = 0.75
	bm25.k1 = 2.0
	bm25.avgWordNum = float64(bm25.totalNumWords / len(bm25.wordSplitArr))
	bm25.Bm25Value = make([][][]float64, len(bm25.WordsMap))
	for i := 0; i < len(bm25.Bm25Value); i++ {
		bm25.Bm25Value[i] = make([][]float64, len(bm25.wordSplitArr))
		for j := 0; j < len(bm25.wordSplitArr); j++ {
			bm25.Bm25Value[i][j] = make([]float64, 3)
		}
	}
}

func CalcBM25(bm25 *BM25, strs, input [][]string) {
	InitBM25Param(bm25, strs, input)
	wordNum := len(bm25.WordsMap)
	docNum := len(bm25.wordSplitArr)
	for i := 0; i < wordNum; i++ {
		for j := 0; j < docNum; j++ {
			bm25.Bm25Value[i][j][0] = float64(i)
			bm25.Bm25Value[i][j][1] = float64(j)
			bm25.Bm25Value[i][j][2] = (bm25.idf[i] * bm25.frequency[i][j] * (bm25.k1 + 1)) /
				(bm25.frequency[i][j] + bm25.k1*(1-bm25.b+(bm25.b*float64(len(bm25.wordSplitArr[j]))/bm25.avgWordNum)))
		}
	}
	//计算完把doc按Bm25Value进行排序

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
