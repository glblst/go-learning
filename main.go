package main

import "fmt"

func gradeOfStudents() {
	type Student struct {
		Name  string
		Grade int
	}

	var n, grade, sumGrade int
	var name string
	students := []Student{}
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
		students = append(students, Student{Name: name, Grade: grade})

	}
	fmt.Println("Список студентов:")
	for _, student := range students {
		fmt.Printf("%s: %d\n", student.Name, student.Grade)
		sumGrade += student.Grade

	}
	fmt.Println("Средняя оценка: ", float64(sumGrade)/float64(len(students)))
}

type Book struct {
	Title string
	Pages int
}

func book() {

	book := Book{Title: "Война и мир", Pages: 1200}
	fmt.Printf("Название: %s, Страниц: %d\n", book.Title, book.Pages)

}

func threeStudents() {
	type Student struct {
		Name  string
		Grade int
	}
	var name string
	var grade int
	students := []Student{}
	for i := 1; i <= 3; i++ {
		fmt.Printf("Введите имя %d: ", i)
		fmt.Scan(&name)
		fmt.Printf("Введите оценку %d: ", i)
		fmt.Scan(&grade)
		students = append(students, Student{Name: name, Grade: grade})
	}
	fmt.Println("Студенты с оценкой выше 3:")
	for _, student := range students {
		if student.Grade > 3 {
			fmt.Printf("%s\n", student.Name)
		}
	}
}
func fasterCar() {
	type Car struct {
		Model string
		Speed int
	}
	var maxSpeed int
	var maxSpeedCar Car
	cars := []Car{
		{Model: "BMW", Speed: 240},
		{Model: "Toyota", Speed: 180},
	}

	for _, car := range cars {
		if car.Speed > maxSpeed {
			maxSpeed = car.Speed
			maxSpeedCar = car
		}
	}
	fmt.Printf("Самая быстрая машина: %s, скорость: %d\n", maxSpeedCar.Model, maxSpeedCar.Speed)
}

type Student struct {
	Name  string
	Grade int
}

func averageGrade(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}
	var sum int
	for _, student := range students {
		sum += student.Grade
	}
	average := float64(sum) / float64(len(students))
	return average
}

type Person struct {
	Name string
	Age  int
}

func oldPerson() {
	persons := []Person{}
	var name string
	var age, maxOld int
	var maxOldPerson Person
	for i := 1; i <= 2; i++ {
		fmt.Printf("Введите имя %d: ", i)
		fmt.Scan(&name)
		fmt.Printf("Введите возраст %d: ", i)
		fmt.Scan(&age)
		persons = append(persons, Person{Name: name, Age: age})

	}
	for _, person := range persons {
		if person.Age > maxOld {
			maxOld = person.Age
			maxOldPerson = person
		}
	}
	fmt.Printf("Самый старший: %s\n", maxOldPerson.Name)
}
func bookMore200() {
	var title string
	var pages int

	books := []Book{}
	for i := 1; i <= 4; i++ {
		fmt.Printf("Название %d: ", i)
		fmt.Scan(&title)
		fmt.Printf("Страниц %d: ", i)
		fmt.Scan(&pages)
		books = append(books, Book{Title: title, Pages: pages})

	}
	fmt.Println("Книги с более чем 200 страницами:")
	for _, book := range books {
		if book.Pages > 200 {
			fmt.Println(book.Title)
		}
	}

}

type Product struct {
	Name  string
	Price float64
}

func cheapProduct() {

	var minProduct Product
	products := []Product{}
	products = append(products, Product{Name: "Яблоко", Price: 10.5})
	products = append(products, Product{Name: "Банан", Price: 15})
	products = append(products, Product{Name: "Хлеб", Price: 5})
	var minPrice float64 = products[0].Price
	for _, product := range products {
		if product.Price < minPrice {
			minPrice = product.Price
			minProduct = product
		}
	}
	fmt.Printf("Самый дешевый продукт: %s, цена: %0.2f\n", minProduct.Name, minProduct.Price)
}
func main() {
	cheapProduct()
	bookMore200()
	oldPerson()
	student := []Student{{Name: "Иван", Grade: 5}, {Name: "Петров", Grade: 4}}
	fmt.Printf("Средняя оценка: %.1f\n", averageGrade(student))
	fasterCar()
	threeStudents()
	book()
	gradeOfStudents()

}
