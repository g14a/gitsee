package models

type LanguageSet struct {
	Languages map[string]bool
}

func NewLanguageSet() *LanguageSet {
	l := &LanguageSet{}
	l.Languages = make(map[string]bool)

	return l
}

func (l *LanguageSet) Add(language string) {
	if _, ok := l.Languages[language]; !ok {
		l.Languages[language] = true
	}
}
