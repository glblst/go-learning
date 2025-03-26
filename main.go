package main

import "fmt"

func numbers() {
	n := 0
	sum := 0
	number := 0
	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Ошибка: количество должно быть положительным.")
		return
	}
	numbers := []int{}
	for i := 1; i <= n; i++ {
		fmt.Println("Введите цифру:")
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Все числа в срезе:", numbers)
	fmt.Println("Сумма чисел в срезе:", sum)
	fmt.Println("Среднее значение среза:", float64(sum)/float64(len(numbers)))
}
func sumArray() {

	sum := 0
	array := [5]int{1, 2, 3, 4, 5}
	for _, num := range array {
		sum += num
	}
	fmt.Println("Сумма чисел массива:", sum)
}
func threeSlice() {
	var str string
	threeSl := []string{}
	for i := 1; i <= 3; i++ {
		fmt.Println("Введите строку:")
		fmt.Scan(&str)
		threeSl = append(threeSl, str)
	}
	fmt.Println("В обратном порядке:")
	for i := 1; i <= 3; i++ {

		fmt.Println(threeSl[len(threeSl)-i])
	}

}
func maxNum() {
	numbers := []int{3, 7, 2, 4}
	max := numbers[0]
	for _, num := range numbers {
		if max < num {
			max = num
		}

	}
	fmt.Printf("Максимальное значение в срезе %v = %d\n", number, max)
}
func whileZero() {
	num := 0
	numbers := []int{}
	for {
		fmt.Println("Введите число. По окончании введите 0.")
		fmt.Scan(&num)
		if num != 0 {
			numbers = append(numbers, num)
		} else {
			fmt.Println(numbers)
			break
		}
	}
}
func onlyEven() {
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}
func main() {
	onlyEven()
	whileZero()
	maxNum()
	threeSlice()
	sumArray()
	numbers()
}
