package main

import (
	"fmt"
	"strings"
)

func main() {
	var name, lastName, fullName string
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	if strings.TrimSpace(name) == "" {
		fmt.Println("Ошибка: имя и фамилия не могут быть пустыми")
		return
	}
	fmt.Println("Введите фамилию:")
	fmt.Scan(&lastName)
	if strings.TrimSpace(lastName) == "" {
		fmt.Println("Ошибка: имя и фамилия не могут быть пустыми")
		return
	}
	fullName = name + " " + lastName
	fmt.Println("Полное имя:", fullName)
	fmt.Println("Длина:", len(fullName))
	fmt.Println("Верхний регистр:", strings.ToUpper(fullName))
}
