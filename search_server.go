package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"mySearch/Indexer"
	"mySearch/Segment"
	"mySearch/Sorter"
	"mySearch/excel"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

var (
	wukongData   = "data/wukong50k_release.xlsx"
	staticFolder = flag.String("static_folder", "static", "静态文件目录")
	//staticFolder  = flag.String("static_folder", "D:\\Study\\code\\Go\\wukong-master\\examples\\codelab\\static", "静态文件目录")
	wordScoreTable = map[string][]int{}
	bm25           = &Indexer.BM25{}
	completeIndex  chan bool
)

/*******************************************************************************
    索引
*******************************************************************************/
func indexExcel() {
	// 读入数据
	input := excel.GetExcelData(wukongData)[1:100] //防止爆内存，暂时只读取100条数据
	segmentWords := Segment.Segmenter(input)

	Indexer.CalcBM25(bm25, segmentWords, input)
	//wordScoreTable = Indexer.GenerateRelationTable(bm25)
	log.Print("添加索引")
	log.Printf("索引了%d条excel数据\n", len(input))
	log.Print("倒排索引生成中...")
	completeIndex <- true
}

/*******************************************************************************
    JSON-RPC
*******************************************************************************/
type JsonResponse struct {
	Docs []string `json:"docs"`
}

func JsonRpcServer(w http.ResponseWriter, req *http.Request) {
	querys := req.URL.Query().Get("query")
	querysArr := Segment.QuerySegmenter(querys)
	idAndScoreArr := [][]float64{}
	for _, query := range querysArr {
		idAndScoreArr = append(idAndScoreArr, findIdAndScore(query)...)
	}
	//按bm25排序由高到低排序
	idAndScoreArr = Sorter.Sorter(idAndScoreArr)
	docs := []string{}
	for _, idAndScore := range idAndScoreArr {
		wordId := int(idAndScore[0])
		docId := int(idAndScore[1])
		docs = append(docs, strings.Replace(bm25.Docs[docId][1], bm25.WordsMap[wordId], "<font color=red>"+bm25.WordsMap[wordId]+"</font>", -1))
	}
	response, _ := json.Marshal(&JsonResponse{Docs: docs})
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(response))
}

func findIdAndScore(query string) [][]float64 {
	wordMapId := 0
	//找出query这个词对应于bm25.wordMapId里的index
	for i := 0; i < len(bm25.WordsMap); i++ {
		if bm25.WordsMap[i] == query {
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

/*******************************************************************************
	主函数
*******************************************************************************/
func main() {

	// 索引
	completeIndex = make(chan bool)
	log.Print("建索引开始")
	go indexExcel()
	log.Print("建索引完毕")

	// 捕获ctrl-c
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			log.Print("捕获Ctrl-c，退出服务器")
			os.Exit(0)
		}
	}()

	http.HandleFunc("/json", JsonRpcServer)
	http.Handle("/", http.FileServer(http.Dir(*staticFolder)))
	log.Print("服务器启动")
	<-completeIndex
	log.Print("倒排索引已生成，服务器启动完成")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
