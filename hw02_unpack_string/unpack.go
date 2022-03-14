package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

// ConvertStrToInt - функция конвертации строки в целое число.
func ConvertStrToInt(str string) (int, error) {
	result, err := strconv.Atoi(str)
	if err == nil {
		return result, nil
	}
	return result, err
}

func Unpack(str string) (string, error) {
	// итоговая возвращаемая переменная
	var result string
	// временная переменная для не цифровых значений
	var tempStr string
	// счетчик для случая идущих друг за другом чисел
	var counter int
	// проверка на пустое значение аргумента
	if str == "" {
		return "", nil
	}
	// основной цикл
	for i, v := range str {
		// проверка цифр идущих друг за другом
		if counter > 1 {
			return "", ErrInvalidString
		}
		intCheck, err := ConvertStrToInt(string(v))
		// проверка на первую цифру в строке
		if i == 0 && err == nil {
			return "", ErrInvalidString
		}
		// запись единичных строковых значений, после которых нет числа + счетчик
		if err != nil {
			tempStr = string(v)
			result += tempStr
			counter = 0
		} else {
			// проверка на число 0 и удаление записанного в результат последнего строкового значения
			if intCheck == 0 {
				result = result[:utf8.RuneCountInString(result)-1]
			} else {
				// запись в результат строки по значению числа
				for i := 0; i < intCheck-1; i++ {
					result += tempStr
				}
			}
			counter++
		}
	}
	return result, nil
}
