package main

import (
	"fmt"
	"strings"
)

func getName() {
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
func lenString() {
	var getStr string
	fmt.Println("Введите строку:")
	fmt.Scan(&getStr)
	fmt.Println("Длина строки:", len(getStr))
}
func twoStr() {
	var str1, str2 string
	fmt.Println("Введите первую строку:")
	fmt.Scan(&str1)
	fmt.Println("Введите вторую строку:")
	fmt.Scan(&str2)
	fmt.Println("Объединенные строки с дефисом", str1+"-"+str2)
}
func inclA() {
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)
	if strings.Contains(str, "a") || strings.Contains(str, "A") {
		fmt.Println("Строка содержит букву a.")
	} else {
		fmt.Println("Строка не содержит букву а.")
	}
}
func lowReg() {
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)
	fmt.Println("Строка в нижнем регистре:", strings.ToLower(str))
}
func lengthStr() {
	var str string
	fmt.Println("Введите строку:")
	fmt.Scan(&str)
	if len(str) > 5 {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
func main() {

	lengthStr()
	lowReg()
	inclA()
	twoStr()
	getName()
	lenString()
}
