package main

import (
	"fmt"
)

func parenthesisStr(preStr string,lnum, rnum int, strList *[]string){
	// fmt.Printf("preStr:%s,lnum:%d,rnum:%d,strList:%+v\n",preStr,lnum,rnum,strList)
	// time.Sleep(time.Second)
	if lnum==0&&rnum==0{
		*strList = append(*strList,preStr)
		return
	}
	if lnum == 0{
		preStr = preStr + ")"
		parenthesisStr(preStr,lnum,rnum-1,strList)
		return
	}
	if lnum == rnum {
		preStr = preStr + "("
		parenthesisStr(preStr,lnum-1,rnum,strList)
		return
	}
	parenthesisStr(preStr+"(",lnum-1,rnum,strList)
	parenthesisStr(preStr+")",lnum,rnum-1,strList)
}

func generateParenthesis(n int) []string {
	strList := make([]string,0,n*n)
	parenthesisStr("",n,n,&strList)
	return strList
}

func main(){
	n := 3
	strList := generateParenthesis(n)
	fmt.Printf("n:%d,strList:%+v\n",n,strList)
}