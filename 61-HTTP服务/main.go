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

func cookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "username",
		Value:    "zhangsan",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "password",
		Value:    "123456",
		HttpOnly: true,
	}

	// 响应报文设置cookie

	// 响应报文设置头部字段Set-Cookie，添加一个cookie
	// w.Header().Set("Set-Cookie", c1.String())
	// 再添加一个cookie
	// w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	// 从请求报文获取cookie

	// 获取所有cookie
	cs := r.Cookies()
	fmt.Fprintln(w, "cookies:", cs)

	// 获取指定名字的cookie
	username, _ := r.Cookie("username")
	password, _ := r.Cookie("password")
	fmt.Fprintln(w, "cookie:", username)
	fmt.Fprintln(w, "cookie:", password)
}

func main() {
	// 设置路由，为指定路径的请求选择处理器来处理请求
	http.HandleFunc("/index", index)
	http.HandleFunc("/readBody", readBody)
	http.HandleFunc("/login", login)
	http.HandleFunc("/tmpl", tmpl)
	http.HandleFunc("/cookie", cookie)
	// 静态文件服务
	http.Handle("/static/", http.StripPrefix("/static",
		http.FileServer(http.Dir("view/static"))))

	// 启动服务器，监听端口，使用默认的多路复用器分发请求
	http.ListenAndServe(":8080", nil)
}
