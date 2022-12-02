package Sorter

//package main
import (
	"sort"
)

type Sorter struct {
}

type Obj struct {
	idAndScoreArr [][]float64
}

func (obj Obj) Len() int {
	return len(obj.idAndScoreArr)
}

func (obj Obj) Less(i, j int) bool {
	return obj.idAndScoreArr[i][1] > obj.idAndScoreArr[j][1]
}

func (obj Obj) Swap(i, j int) {
	obj.idAndScoreArr[i], obj.idAndScoreArr[j] = obj.idAndScoreArr[j], obj.idAndScoreArr[i]
}

func (s *Sorter) Sorter(idAndScoreArr [][]float64) [][]float64 {
	obj := &Obj{
		idAndScoreArr: idAndScoreArr,
	}
	sort.Sort(obj)
	//fmt.Println(obj)
	return obj.idAndScoreArr
}

//func main() {
//	n := make([][]float64,3)
//	n[0] = []float64{1,4}
//	n[1] = []float64{2,3}
//	n[2] = []float64{3,5}
//	Sorter(n)
//}
