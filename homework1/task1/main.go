package main

import (
	"fmt"
	"os"
)

func main() {
	var rubles float64
	const course float64 = 70.04

	fmt.Println("Программа для конвертации рублей в доллары")
	fmt.Println("Введите сумму в рублях:")
	_,err:=fmt.Scanf("%f", &rubles)
	if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
	fmt.Println("Вы ввели: ",rubles)
	fmt.Println("Расчитываем по курсу $:",course)
	dollars := rubles/course;
	fmt.Println("Итого долларов:", dollars)
	fmt.Println("The end")

}
