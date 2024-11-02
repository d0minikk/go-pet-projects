package phrases

import (
	"errors"
	"math/rand"
)

type PhraseManager struct {
	themes map[string][]string
}

func NewPhraseManager() *PhraseManager {
	return &PhraseManager{
		themes: themes,
	}
}

func (pm *PhraseManager) GetRandomPhrase(theme string) (string, error) {
	phrases, exists := pm.themes[theme]
	if !exists {
		return "", errors.New("theme not found")
	}
	if len(phrases) == 0 {
		return "", errors.New("no phrases available for the theme")
	}

	randomIndex := rand.Intn(len(phrases))
	return phrases[randomIndex], nil
}

func (pm *PhraseManager) GetAllThemes() []string {
	keys := make([]string, 0, len(pm.themes))
	for k := range pm.themes {
		keys = append(keys, k)
	}

	return keys
}

func (pm *PhraseManager) GetPhrasesByTheme(theme string) ([]string, error) {
	phrases, exists := pm.themes[theme]
	if !exists {
		return nil, errors.New("theme not found")
	}

	return phrases, nil
}
