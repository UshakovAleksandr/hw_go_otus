package hw03frequencyanalysis

import "strings"

// TextPrepare - подготовка входных данных.
func TextPrepare(text string) []string {
	// очищение от лишних пробелов. Строка
	stripText := strings.Join(strings.Fields(text), " ")
	// разделение на отдельные элементы по пробелу. Слайс
	return strings.Split(stripText, " ")
}

// CountWordInSlice - заполнение словарика.
func CountWordInSlice(word *string, cache map[string]int) {
	cache[*word]++
}

// FindMaxKV - поиск максимальных значений, их удаление.
// сортировка отбираемых ключей.
func FindMaxKV(cache map[string]int) string {
	// временные переменных для k, v
	var maxSortKey string
	var maxVal int
	// цикл поиска максимального значения
	for k, v := range cache {
		// условие сортировки через bool-сравнение строк (k < maxSortKey)
		if v > maxVal || (v == maxVal && k < maxSortKey) {
			maxVal = v
			maxSortKey = k
		}
	}
	// удаление максимальных k, v
	delete(cache, maxSortKey)

	return maxSortKey
}

// GetTop10 - сбор самых частых и отсортированных слов.
func GetTop10(cache map[string]int, flag int) []string {
	// инициализация пустого слайса с cap == flag
	result := make([]string, 0, flag)
	// заполнение слайса
	for i := 0; i < flag; i++ {
		result = append(result, FindMaxKV(cache))
	}

	return result
}

// Top10 - основная функция.
func Top10(text string) []string {
	// проверка на пустую строку
	if text == "" {
		return nil
	}
	// количество самых частых слов
	const MaxRepeatWordsNumber = 10
	// подготовка входных данных
	sliceOfWords := TextPrepare(text)
	// инициализация кэша, len(sliceOfWords) указано для экономии памяти
	cache := make(map[string]int, len(sliceOfWords))
	// заполнение кэша: подсчет количества слов
	for _, v := range sliceOfWords {
		// для решения придирок линтера - G601: Implicit memory aliasing in for loop. (gosec)
		v := v
		CountWordInSlice(&v, cache)
	}
	// получение финального слайса и возрат его в main(в тест, данном случае)
	return GetTop10(cache, MaxRepeatWordsNumber)
}
