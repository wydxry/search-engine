package main

//import "fmt"
//
///**
//本来是用来回溯切分query，但因为gojieba提供了全模式分词，所以这个就不用了
// */
//func main() {
//	input := []string{"1", "2", "3", "4"}
//	output := [100]string{}
//	dfs(input, 0, output, 0)
//}
//var ret []string
//func dfs(input []string, index int, output [100]string, outLength int) {
//	if len(input) == index {
//		//fmt.Println(output)
//		fmt.Println(split(output, outLength))
//		return
//	}
//
//	output[outLength] = input[index]
//
//	output[outLength + 1] = "_"
//
//	dfs(input, index + 1, output, outLength + 2)
//
//	if(len(input) != index + 1) {
//		dfs(input, index + 1, output, outLength + 1)
//	}
//}
//
//func split(output [100]string, outLength int) []string {
//	ret := []string{}
//	idx := 0
//	for i := 0; i < outLength; i++ {
//		if output[i] == "_" {
//			tmpStr := ""
//			for j := idx; j < i; j++ {
//				tmpStr += output[j]
//			}
//			ret = append(ret, tmpStr)
//			idx = i + 1
//		}
//	}
//	return ret
//
//}
