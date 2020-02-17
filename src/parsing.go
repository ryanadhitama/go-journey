package main

import "fmt"
import "net/url"

func main(){
    var urlString = "https://blog.ryanadhitama.com/post?name=sabar&page=2"

    var u, e = url.Parse(urlString)
    if e != nil {
        fmt.Println(e.Error())
        return
    }

    fmt.Printf("url: %s\n", urlString)

    fmt.Printf("\nprotocol: %s\n", u.Scheme) // https
    fmt.Printf("host: %s\n", u.Host)       // blog.ryanadhitama.com
    fmt.Printf("path: %s\n", u.Path)       // post

    var name = u.Query()["name"][0]
    var page = u.Query()["page"][0]

    fmt.Println("Query name : "+name+", page : "+page);
}
