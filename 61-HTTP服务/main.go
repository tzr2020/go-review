package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")

	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "Path:", r.URL.Path)

	fmt.Fprintln(w, "RawQuery:", r.URL.RawQuery)

	fmt.Fprintln(w, "Header:", r.Header)
	fmt.Fprintln(w, "Accept-Encoding:", r.Header.Get("Accept-Encoding"))
}

func readBody(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, "Body:", string(body))

	default:
		fmt.Fprintln(w, "请发送POST请求")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Fprintln(w, "RawQuery:", r.URL.RawQuery)

		/*
			获取请求参数
			1.查询字符串
			2.请求主体
		*/
		r.ParseForm()
		fmt.Fprintln(w, "Form:", r.Form)
		fmt.Fprintln(w, "PostForm:", r.PostForm)

		// fmt.Fprintln(w, "username:", r.FormValue("username"))
		// fmt.Fprintln(w, "password:", r.FormValue("password"))

		// fmt.Fprintln(w, "username:", r.PostFormValue("username"))
		// fmt.Fprintln(w, "password:", r.PostFormValue("password"))

	default:
		fmt.Fprintln(w, "请发送POST请求")
	}
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	// t := template.New("index.html")
	// template.Must(t.ParseFiles("index.html"))
	// template.Must(template.ParseFiles("index.html"))
	// t.Execute(w, nil)

	t2, _ := template.ParseFiles("about.html", "home.html")
	// t2, _ := template.ParseGlob("*html")
	t2.ExecuteTemplate(w, "home.html", nil)
}

func main() {
	// 设置路由，为指定路径的请求选择处理器来处理请求
	http.HandleFunc("/index", index)
	http.HandleFunc("/readBody", readBody)
	http.HandleFunc("/login", login)
	http.HandleFunc("/tmpl", tmpl)

	// 启动服务器，监听端口，使用默认的多路复用器分发请求
	http.ListenAndServe(":8080", nil)
}
