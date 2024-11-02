package server

import (
	"encoding/json"
	"net/http"
	"nice-progress/phrases"
)

type Server struct {
	PhraseManager *phrases.PhraseManager
}

func NewServer(pm *phrases.PhraseManager) *Server {
	return &Server{
		PhraseManager: pm,
	}
}

func (s *Server) Start(port string) error {
	http.HandleFunc("/random", s.handleRandomPhrase)
	http.HandleFunc("/themes", s.handleGetThemes)
	http.HandleFunc("/phrases", s.handleGetPhrasesByTheme)
	return http.ListenAndServe(":"+port, nil)
}

func (s *Server) handleRandomPhrase(w http.ResponseWriter, r *http.Request) {
	theme := r.URL.Query().Get("theme")
	if theme == "" {
		theme = "original"
	}
	phrase, err := s.PhraseManager.GetRandomPhrase(theme)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(w, map[string]string{"phrase": phrase})
}

func (s *Server) handleGetThemes(w http.ResponseWriter, r *http.Request) {
	themes := s.PhraseManager.GetAllThemes()
	jsonResponse(w, themes)
}

func (s *Server) handleGetPhrasesByTheme(w http.ResponseWriter, r *http.Request) {
	theme := r.URL.Query().Get("theme")
	if theme == "" {
		http.Error(w, "theme parameter is required", http.StatusBadRequest)
		return
	}
	phrases, err := s.PhraseManager.GetPhrasesByTheme(theme)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(w, phrases)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
