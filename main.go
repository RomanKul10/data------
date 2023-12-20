package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("data.txt")
	// створюємо змінні, щоб зберігати туди помилку відкриття файлу
	if err != nil { //код помилки( перевірка на помилку)
		log.Fatal(err)
	}
	//сканер нашого файлу
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// сканер для перевірки файлу по рядках
		fmt.Println(scanner.Text()) //викликаємо у вигляді тексту наш файл через сканер
	}

	err = file.Close() // закриваємо файл

	if err != nil { //код помилки( перевірка на помилку)
		log.Fatal(err)
	}
	if scanner.Err() != nil { // робимо перевірку на випадок помилки при скануванні
		log.Fatal(scanner.Err())
	}
}
