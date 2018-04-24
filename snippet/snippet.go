// Package snippet provides snippet models.
package snippet

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Rules       []string
	Body        []string
	Description string
}

/**
 * So one file should return []Snippet
 * AtomSnip --> CSON
 * CodeSnip --> JSON
 * SublSnip --> XML
 * UltiSnip --> TXT
 */
