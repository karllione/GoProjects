package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		a, b, c float64
	)

	fmt.Println("[Поиск корней квадратного уравнения ]")
	fmt.Println("Введите значение для переменной a: ")
	for {
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Println("Пожалуйста, введите корректное значение a: ")
			continue
		}
		break
	}

	fmt.Println("Введите значение для переменной b: ")
	for {
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Println("Пожалуйста, введите корректное значение b: ")
			continue
		}
		break
	}

	fmt.Println("Введите значение для переменной c: ")
	for {
		if _, err := fmt.Scan(&c); err != nil {
			fmt.Println("Пожалуйста, введите корректное значение c: ")
			continue
		}
		break
	}

	if a == 0 {
		fmt.Println("Уравнение квадратным не является, поскольку коэффициент а равен 0")
		return
	}

	D := (b * b) - 4*(a*c)
	if D > 0 {
		var (
			x1, x2 float64
		)

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Уравнение имеет 2 корня.\nD = " + fmt.Sprint(D))
		fmt.Println("Корень x1 = " + fmt.Sprint(x1) + "\tКорень x2 = " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)
		fmt.Println("Уравнение имеет 1 корень.\nD = " + fmt.Sprint(D))
		fmt.Println("Корень x = " + fmt.Sprint(x))
	} else {
		fmt.Println("Уравнение корней не имеет.\nD = " + fmt.Sprint(D))
	}
}
