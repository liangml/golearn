package main

import (
	"fmt"
	"net/http"
	"text/template"
)
// UserInfo 是一个测试用例
type UserInfo struct{

	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./src/GinLean/hello.tpl")
	if err != nil {
		fmt.Println("Parse template faile")
		return
	}

	// 3.渲染模板
	u1 := UserInfo{
		Name: "梁孟麟",
		Gender: "男",
		Age: 18,
	}
	m1 := map[string]interface{}{
		"name": "梁孟麟",
		"gender": "男",
		"age": 18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
	// name := "梁孟麟"
	// err = t.Execute(w, name)
	// if err != nil {
	// 	fmt.Printf("rander template faild,err:%v", err)
	// 	return
	// }
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start faild ,err:%v", err)
		return
	}
}
