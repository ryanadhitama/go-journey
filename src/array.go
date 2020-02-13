package main

import "fmt"

func main(){
    // declaration
    var name[4] string;
    
    // initialization
    name[0] = "Ryan"
    name[1] = "Adit"

    fmt.Println(name[0],name[1])


    //declaration and initialization together
    var names = [2]string{"Ryan Adhitama", "Adhitama Putra"}

    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }
}