package utils

type Set map[string]bool

func (s Set) Add(element string) {
	s[element] = true
}

func (s Set) Remove(element string) {
	delete(s, element)
}

func (s Set) Contains(element string) bool {
	return s[element]
}
