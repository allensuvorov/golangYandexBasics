// Статически создаём ошибку.
// var ErrFileReading = errors.New("read_text_file: read file error") //  хорошей практикой является начинать текст ошибки с названия пакета, где она объявлена, так будет проще её найти

// func ReadTextFile() (string, error) {
// 	if data, err := os.ReadFile(`nothing.txt`); err != nil {
// 		// будет вызван метод 'Error() string', который преобразует ошибку в строку
// 		fmt.Println(err)
// 		return ErrFileReading
// 	}
// 	fmt.Println(string(data))
// 	return string(data), nil
// }

// func ReadTextFileByName(filename string) (string, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		// вернём ошибку на русском языке
// 		return ``, fmt.Errorf(`не удалось прочитать файл (%s): %v`, filename, err)
// 	}
// 	return string(data), nil
// }

// func ReadTextFile(filename string) (string, error) {
//     data, err := os.ReadFile(filename)
//     if err != nil {
//         // возвратим ошибку на русском языке
//         return ``, fmt.Errorf(`не удалось прочитать файл (%s): %w`, filename, err)
//     }
//     return string(data), nil
// }

// Статически создаём ошибку.
var ErrFileReading = errors.New("read_text_file: read file error") 

func ReadTextFileByName(filename string) (string, error) {
    if data, err := os.ReadFile(filename); err != nil {
        // будет вызван метод 'Error() string', который преобразует ошибку в строку
        fmt.Println(err)
        return errors.Wrapf(ErrFileReading, "file not exist %s", filename)
    }   
    fmt.Println(string(data))
    return string(data), nil
}

////////////// 
import (
    "fmt"
    "time"
)
// Создадим собственный тип, который удовлетворяет интерфейсу error
// TimeError — тип для хранения времени и текста ошибки.
type TimeError struct {
    Time time.Time
    Text string
}

// Error добавляет поддержку интерфейса error для типа TimeError.
func (te TimeError) Error() string {
    return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}

// NewTimeError возвращает переменную типа TimeError c текущим временем.
func NewTimeError(text string) TimeError {
    return TimeError{
        Time: time.Now(),
        Text: text,
    }
}

func testFunc(i int) error {
    // несмотря на то что NewTimeError возвращает тип TimeError,
    // у testFunc тип возвращаемого значения равен error
    if i == 0 {
        return NewTimeError(`параметр в testFunc равен 0`)
    }
    return nil
}

func main() {
    if err := testFunc(0); err != nil {
        fmt.Println(err)
    }
}