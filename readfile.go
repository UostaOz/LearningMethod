package readfile //пакет для чтения файлов и получения сегмента значений и возможной ошибки

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntFileRead(name string) ([]int, error) { //Код для чтения int
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Can't read this file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text []int
	for scanner.Scan() {

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Can't make int")
		}
		text = append(text, i)
	}

	return text, nil
}
func FloatFileRead(name string) ([]float64, error) { //Для чтения типа float64
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Can't read this file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text []float64
	for scanner.Scan() {

		i, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("Can't make int")
		}
		text = append(text, i)
	}

	return text, nil
}
func StringFileRead(name string) ([]string, error) { // Для чтения строк
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Can't read this file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {

		i := scanner.Text()
		i = strings.TrimSpace(i)
		text = append(text, i)
	}

	return text, nil
}
