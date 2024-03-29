package models

type TemplateData struct {
	Strings   map[string]string
	Ints      map[string]int
	Floats    map[string]float64
	Data      map[string]any
	CsrfToken string
	Flash     string
	Warning   string
	Error     string
}
