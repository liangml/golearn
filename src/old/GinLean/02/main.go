package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func f1(w http.ResponseWriter,r *http.Request){
	// 要么有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	kua := func (name string)(string, error)  {
		return name + "年轻又帅气",nil
	}

	t,err := template.New("f.tpl").Funcs(template.FuncMap{
		"kua": kua,
	}).ParseFiles("./src/GinLean/02/f.tpl")
	if err != nil{
		fmt.Printf("fail %v",err)
		return
	}
	name := "梁孟麟"
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/",f1)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("http faild %v",err)
	}
}