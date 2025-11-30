package htmltemplate

import "fmt"

// TemplateError represents an error in template parsing or execution
type TemplateError struct {
	Op   string // "parse" or "execute"
	File string // template file path
	Err  error  // original error
}

func (e *TemplateError) Error() string {
	return fmt.Sprintf("%s template %s: %v", e.Op, e.File, e.Err)
}

// Unwrap allows errors.Is / errors.As
func (e *TemplateError) Unwrap() error {
	return e.Err
}
