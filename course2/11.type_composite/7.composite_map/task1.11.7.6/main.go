package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

// filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
// Рекурсивно обрабатывает каждое предложение
func filterWords(text string, censorMap map[string]string) string {
	// Разделение текста на предложения с помощью splitSentences
	sentences := splitSentences(text)

	// Если предложений больше одного, рекурсивно обрабатываем каждое предложение
	if len(sentences) > 1 {
		// Рекурсивный вызов для обработки оставшихся предложений
		for i := range sentences {
			sentences[i] = filterWords(sentences[i], censorMap)
		}
	}

	// Если текст состоит из одного предложения, фильтруем это предложение
	if len(sentences) == 1 {
		return filterSentence(sentences[0], censorMap)
	}

	// Возвращаем объединённый текст
	return strings.Join(sentences, " ")
}

// filterSentence Фильтрует одно предложение
func filterSentence(sentence string, censorMap map[string]string) string {
	// Разделение текста на отдельные слова с помощью strings.Fields
	words := strings.Fields(sentence)

	// Создание пустой карты уникальных слов с помощью make(map[string]bool)
	uniqueWords := make(map[string]bool)
	// Карта для отслеживания замененных слов
	seenReplacements := make(map[string]bool)

	// Обработка каждого слова в цикле
	for i, word := range words {
		// Преобразуем слово в нижний регистр для проверки уникальности
		lowerWord := strings.ToLower(word)

		// Если слово содержится в карте цензурных слов
		if censoredWord, ok := censorMap[lowerWord]; ok {
			// Замена слова на цензурное
			replacement := CheckUpper(word, censoredWord)
			// Если это слово уже было заменено ранее, удаляем его (ставим пустое)
			if _, seen := seenReplacements[replacement]; seen {
				fmt.Println(replacement)
				words[i] = ""
				//words[i-1] = ""
			} else {
				// Отмечаем это слово как замененное
				seenReplacements[strings.ToLower(replacement)] = true
				words[i] = replacement

			}
		} else if _, found := uniqueWords[lowerWord]; found {
			// Если слово уже встречалось, удаляем его
			words[i] = ""
		} else {
			// Если слово уникальное, добавляем в карту уникальных слов
			uniqueWords[lowerWord] = true

		}
	}

	// Замена в слайсе слов при помощи карты уникальных слов и их индекса

	return WordsToSentence(words)
}

// WordsToSentence Удаляет пустые слова из слайса и объединяет их в предложение, добавляя в конце восклицательный знак
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// CheckUpper Проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

// splitSentences Разделяет текст на предложения по восклицательному знаку
func splitSentences(message string) []string {
	originSentences := strings.Split(message, "!")
	var orphan string
	var sentences []string

	for i, sentence := range originSentences {
		words := strings.Fields(sentence)

		if len(words) == 1 {
			// Если осталось одно слово, собираем его в orphan
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}

		if orphan != "" {
			// Если есть orphan, добавляем его к текущему предложению
			originSentences[i] = strings.Join([]string{orphan, sentence}, " ") + "!"
			orphan = ""
		}

		sentences = append(sentences, originSentences[i])
	}

	// Удаляем пустые строки, если они появились в результате разбиения
	var filteredSentences []string
	for _, sentence := range sentences {
		if len(sentence) > 0 {
			filteredSentences = append(filteredSentences, sentence)
		}
	}

	return filteredSentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "яблоки",
		"лайткоин": "яблоки",
		"эфир":     "яблоки",
	}

	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText)
}
