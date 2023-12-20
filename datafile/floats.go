package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {

	var numbers [3]float64

	file, err := os.Open(fileName)
	// створюємо змінні, щоб зберігати туди помилку відкриття файлу
	if err != nil { //код помилки( перевірка на помилку)
		return numbers, err
	}

	i := 0
	//сканер нашого файлу
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// сканер для перевірки файлу по рядках
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //викликаємо у вигляді тексту наш файл через сканер
		if err != nil {                                          //код помилки( перевірка на помилку)
			return numbers, err
		}
		i++
	}

	err = file.Close() // закриваємо файл

	if err != nil { //код помилки( перевірка на помилку)
		return numbers, err
	}
	if scanner.Err() != nil { // робимо перевірку на випадок помилки при скануванні
		return numbers, scanner.Err()
	}

	return numbers, nil
}
