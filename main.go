package main

import (
	"fmt"
	"log"
	"myapp/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt") // створюємо змінні і викликаємо файл, за методом getfloats, посилаючись на наш файл -data.txt

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0 // змінна де буде зберігатись сума, і за замовчуванням вона = 0

	for _, number := range numbers {
		sum += number
	}
	count := float64(len(numbers))
	fmt.Print(sum / count)

}
