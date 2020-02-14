package main

import "fmt"

func main(){

    //deklarasi terus inisialisasi
    var data map[string]int
    data = map[string]int{}
    data["one"] = 1

    fmt.Println(data["one"])


    //deklarasi barengan dengan inisialisasi
    var chicken = map[string]int{
        "januari":  50,
        "februari": 40,
    }

    //iteration
    for key, val := range chicken {
        fmt.Println(key, "  \t:", val)
    }

    // delete
    delete(chicken,"januari")

     //iteration
    for key, val := range chicken {
        fmt.Println(key, "  \t:", val)
    }
}