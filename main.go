package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"strings"
)

func main() {
	var (
		arg1, arg2 string
	)
	if len(os.Args) < 3 {
		arg1 = "data.txt"
		arg2 = "result.txt"
	} else {
		arg1 = os.Args[1]
		arg2 = os.Args[2]
	}

	dataFile := getLinesFiles(arg1)
	outputFile := getOPutputFile(arg2)

	defer outputFile.Close()

	// отбираем выражения типа 2+3=?
	re := regexp.MustCompile(`([0-9]+)([+-/*]{1})([0-9]+)\=\?`)

	writer := bufio.NewWriter(outputFile)
	for index, line := range dataFile {
		index++
		data := []byte(line)

		testStr := string(data)

		arr := re.FindAllStringSubmatch(testStr, -1)
		if arr != nil {
			fmt.Printf("строка %d записана\n", index)
			result, err1 := getResult(arr[0])
			if err1 != nil {
				fmt.Println("Что то пошло не так: ", err1)
			}
			// формируем формат записи в выходной файл
			dsave := fmt.Sprintf("%s%s%s%s%d\n", arr[0][1], arr[0][2], arr[0][3], "=", int(result))
			_, err := writer.Write([]byte(dsave))
			if err != nil {
				fmt.Println("Ошибка при записи данных: ", err)
				return
			}
		} else {
			fmt.Printf("строка %d не записана\n", index)
			continue
		}
	}
	// Вызов Flush() для записи данных из буфера в файл
	err := writer.Flush()
	if err != nil {
		fmt.Println("Ошибка при сбросе буфера:", err)
		return
	}
}

// Получение результата математических действий для дальнейшей записи в файл вывода
func getResult(s []string) (float64, error) {
	n1, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Println(err, n1)
		return 0, err
	}
	n2, err := strconv.Atoi(s[3])
	if err != nil {
		fmt.Println(err, n2)
		return 0, err
	}
	switch s[2] {
	case "+":
		return float64(n1 + n2), nil
	case "-":
		return float64(n1 - n2), nil
	case "*":
		return float64(n1 * n2), nil
	case "/":
		return float64(n1 / n2), nil
	}
	return 0, nil
}

// возвращает файл вывода
func getOPutputFile(nameFile string) *os.File {
	file, err := os.OpenFile(nameFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла: ", err)
		os.Exit(1)
	}
	return file
}

// вернет массив строчек входного файла
func getLinesFiles(nameFile string) []string {
	dataFile, err := ioutil.ReadFile(nameFile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dataFile), "\n")
	return lines
}
