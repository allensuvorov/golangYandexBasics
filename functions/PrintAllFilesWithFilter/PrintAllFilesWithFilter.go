package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	PrintAllFilesWithFilter(`C:\Users\miniPC\OneDrive\Documents\Allen Suvorov Files\Coding\yandex\golangYandexBasics\functions`, "Filter")
}

func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())

		// печатаем имя элемента если в пути есть строка filter
		if strings.Contains(filename, filter) {
			fmt.Println(filename)

		}
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}
