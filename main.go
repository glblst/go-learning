package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var yearBirth int
	var age int
	fmt.Println("Ваше имя:")
	fmt.Scan(&name)
	fmt.Println("Ваш год рождения:")
	fmt.Scan(&yearBirth)
	currentTime := time.Now()
	currentYear := currentTime.Year()
	age = currentYear - yearBirth

	if age <= 0 {
		fmt.Println("Возраст не может быть меньше нуля")
	} else {
		fmt.Printf("Приветствую %s, вам %d лет в %d году\n", name, age, currentYear)
	}

}
