package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type msg struct {
	m string
}

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m.m)
}

func main() {
	Static()
	Msg()
	About()
	Users()
	Form()
	Template()
	Tempate2()

	StartServer()
}

func Static() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
}

func Msg() {
	msgHandler := msg{m: "Hello from Web Server in Go"}
	http.Handle("/msg", msgHandler)
}

func About() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
}

func Form() {

	http.HandleFunc("/form/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/user.html")
	})

	http.HandleFunc("/form/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
		age := r.FormValue("userage")

		fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	})
}

func Users() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		fmt.Fprintf(w, "Имя: %v \nВозраст: %s", name, age)
	})
}

type ViewData struct {
	Title   string
	Message string
}

func Template() {
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:   "World Cup",
			Message: "FIFA will never regret it",
		}

		//весь html тут или...
		//tmpl := template.Must(template.New("data").Parse(`<div>
		//    <h1>{{ .Title}}</h1>
		//    <p>{{ .Message}}</p>
		//</div>`))

		//весь html ввиде файла
		tmpl, err := template.ParseFiles("./static/template.html")

		if err != nil {
			fmt.Println("something went wrong")   // вывод в коносль
			fmt.Fprint(w, "something went wrong") // вывод на веб странице
		} else {
			tmpl.Execute(w, data)
		}

	})
}

type ViewData2 struct {
	Title     string
	Users     []User
	Available bool
}
type User struct {
	Name string
	Age  int
}

func Tempate2() {
	data := ViewData2{
		Title: "Users List",
		Users: []User{
			User{Name: "Tom", Age: 21},
			User{Name: "Kate", Age: 23},
			User{Name: "Alice", Age: 25},
		},
		Available: true,
	}
	http.HandleFunc("/template2", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.New("data").Parse(`<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{ .Title}}</title>
    </head>
    <body>
        <h1>{{ .Title}}</h1>
        <ul>
            {{range .Users}}
                <li>
                    <div><b>{{ .Name }}</b>: {{ .Age }}</div>
                </li>
            {{end}}
			{{if .Available}}
            	<p>Available</p>
            {{end}}
        </ul>
    </body>
</html>`))

		tmpl.Execute(w, data)
	})
}

func StartServer() {
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
