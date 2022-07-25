package main

import "testing"

func Test_stringStat(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			"Старт",
			"Старт",
			`с - 1
т - 2
а - 1
р - 1`,
		},
		{
			"Финиш",
			"ФиНИшшш",
			`ф - 1
и - 2
н - 1
ш - 3`,
		},
		{
			"Булки с чаем",
			"Съешь ещё этих мягких французских булок, да выпей чаю",
			`с - 2
ъ - 1
е - 3
ш - 1
ь - 1
щ - 1
ё - 1
э - 1
т - 1
и - 3
х - 3
м - 1
я - 1
г - 1
к - 3
ф - 1
р - 1
а - 3
н - 1
ц - 1
у - 2
з - 1
б - 1
л - 1
о - 1
, - 1
д - 1
в - 1
ы - 1
п - 1
й - 1
ч - 1
ю - 1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringStat(tt.word); got != tt.want {
				t.Errorf("stringStat() = %v, want %v", got, tt.want)
			}
		})
	}
}
