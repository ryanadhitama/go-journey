package main

import "fmt"
import "net/http"
import "html/template"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo!")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
        var data = map[string]string{
            "name":    "Ryan Adhitama Putra",
            "message": "Halo apa kabar",
        }

        var t, err = template.ParseFiles("template/about.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        t.Execute(w, data)
    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}