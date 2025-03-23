package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)

	if age < 0 {
		fmt.Println("Ошибка:", name, ", возраст не может быть отрицательным!")
	} else if age < 18 {
		fmt.Println("Привет,", name, "! Ты ещё молод, тебе всего", age, "лет.")
	} else {
		fmt.Println("Здравствуйте,", name, "! Вам уже", age, "лет, отличный возраст!")
	}
}
