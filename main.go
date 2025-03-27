package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func numberOfStudents() {
	var n, sum int
	var name string
	var grade int
	students := make(map[string]int)
	fmt.Println("Введите количество студентов:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Ошибка: количество должно быть положительным")
		return
	}
	for i := 1; i <= n; i++ {
		fmt.Println("Введите фамилию студента:")
		fmt.Scan(&name)
		fmt.Println("Введите оценку студента:")
		fmt.Scan(&grade)
		students[name] = grade
		sum += grade
	}
	for age, value := range students {
		fmt.Printf("%s: %d\n", age, value)
	}

	fmt.Println("Средняя оценка:", float64(sum)/float64(len(students)))
}
func countryCapital() {
	countries := make(map[string]string)
	countries["США"] = "Вашингтон"
	countries["Russia"] = "Москва"
	countries["Япония"] = "Токио"
	for key, value := range countries {
		fmt.Printf("%s: %s\n", key, value)
	}

}
func nameRequest() {
	names := make(map[string]int)
	names["Мария"] = 30
	names["Piter"] = 22
	names["Alex"] = 33
	names["Jon"] = 44
	var name string

	fmt.Println("Enter name:")
	fmt.Scan(&name)
	key, ok := names[name]
	if ok {
		fmt.Printf("Есть такое имя %s, %d года.\n", name, key)
	} else {
		fmt.Println("Такого имени нет.")
	}
}
func product() {
	products := make(map[string]float64)
	products["Apple"] = 22.3
	products["Banana"] = 35.7
	products["Jam"] = 88.12
	products["Mango"] = 12.90
	var max float64
	var product string
	for key, value := range products {
		if value > max {
			max = value
			product = key

		}
	}
	fmt.Println("Самый дорогой продукт: ", product, max)

}
func emptyKey() {
	names := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	var name string
	var age int
	for {
		fmt.Println("Enter name (or press Enter to finish):")
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
		if name == "" {
			break

		}
		fmt.Println("Enter age:")
		fmt.Scan(&age)
		names[name] = age
	}

	for name, age := range names {
		fmt.Printf("%s: %d\n", name, age)
	}
}
func heightAbove() {
	names := make(map[string]int)
	names["Maria"] = 130
	names["Pavel"] = 150
	names["Alex"] = 160
	names["Pavel"] = 172
	names["Jon"] = 192
	for name, height := range names {
		if height > 170 {
			fmt.Printf("%s: %d\n", name, height)
		}
	}

}
func main() {
	heightAbove()
	emptyKey()
	product()
	nameRequest()
	countryCapital()
	numberOfStudents()

}
