package main

import "fmt"

func main() {
    var name string
    var age int
    var hobby string

    fmt.Print("Как вас зовут? ")
    fmt.Scan(&name)

    fmt.Print("Сколько вам лет? ")
    fmt.Scan(&age)

    fmt.Print("Какое у вас хобби? ")
    fmt.Scan(&hobby)

    if age < 0 {
        fmt.Println("Возраст не может быть отрицательным!")
    } else {
        fmt.Println("Привет,", name, "! Тебе", age, "лет, и твоё хобби -", hobby)
    }
}