// Статически создаём ошибку.
var ErrFileReading = errors.New("read_text_file: read file error") //  хорошей практикой является начинать текст ошибки с названия пакета, где она объявлена, так будет проще её найти

func ReadTextFile() (string, error) {
	if data, err := os.ReadFile(`nothing.txt`); err != nil {
		// будет вызван метод 'Error() string', который преобразует ошибку в строку
		fmt.Println(err)
		return ErrFileReading
	}
	fmt.Println(string(data))
	return string(data), nil
}

func ReadTextFileByName(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		// вернём ошибку на русском языке
		return ``, fmt.Errorf(`не удалось прочитать файл (%s): %v`, filename, err)
	}
	return string(data), nil
}