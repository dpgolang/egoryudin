package main

// Кастомный тип, имеющий одну ф-ию. Слайс типа bool проверяется на наличие хотя бы одного значения true.
type boolSlice []bool

func (b boolSlice) contains(trueOrFalse bool) bool {
	for _, a := range b {
		if a == trueOrFalse {
			return true
		}
	}
	return false
}
