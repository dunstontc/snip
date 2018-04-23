package snippet

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Rules       []string
	Body        []string
	Description string
}
