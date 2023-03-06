package main

import (
	"strconv"
	"strings"
)

type view struct {
	char  string
	count int
}

func main() {
	word := "Съешь ещё этих мягких французских булок, да выпей чаю"
	stringStat(word)
}

func stringStat(word string) string {
	views := getViews(word)
	result := getString(views)

	return result
}

func getViews(word string) []view {
	count := 0
	key := ""
	index := 0
	oldIndex := 0
	exists := false
	mapCharIndex := make(map[string]int)
	views := make([]view, 0, len(word))

	for _, char := range word {
		if char == ' ' {
			continue
		}
		key = strings.ToLower(string(char))
		oldIndex, exists = mapCharIndex[key]
		if exists {
			count = views[oldIndex].count
			count++
			views[oldIndex] = view{key, count}
			continue
		}

		views = append(views, view{key, 1})
		mapCharIndex[key] = index
		index++
	}

	return views
}

func getString(views []view) string {
	count := 0
	char := ""
	result := ""

	for i := 0; i < len(views); i++ {
		char = views[i].char
		count = views[i].count
		if i > 0 {
			result = result + "\n"
		}

		result = result + char + " - " + strconv.Itoa(count)
	}

	return result
}
