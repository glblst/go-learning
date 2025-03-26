package main

import "fmt"

func main() {
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
	for key, value := range students {
		fmt.Printf("%s: %d\n", key, value)
	}
	
	fmt.Println("Средняя оценка:", float64(sum)/float64(len(students)))
}
