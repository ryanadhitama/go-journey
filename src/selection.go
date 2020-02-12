package main

import "fmt"

func main(){
    var angka int = 10
    angka_2 := 12

    //if else
    if angka > angka_2 {
        fmt.Println("false")
    } else {
        fmt.Println("true")
    }

    //switch case
    switch angka {
        case 1 :
            fmt.Println("Angka 1")

        case 10 : 
            fmt.Println("Angka lebih dari 1")
    }

    //switch with if else
    switch {
        case angka < 10  :
            fmt.Println("Angka kurang dari 10")
            break
        default :
            fmt.Println("Angka lebih dari 10")
    }
}