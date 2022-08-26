package main

if data, err := os.ReadFile(`nothing.txt`); err != nil {
    // будет вызван метод 'Error() string', который преобразует ошибку в строку
    fmt.Println(err)
} else {
    fmt.Println(string(data))
}
