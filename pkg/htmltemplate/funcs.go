package htmltemplate

import (
	"html/template"
	"strings"
	"time"
)

// Helper functions for templates
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		// --- String helpers ---
		"upper":      strings.ToUpper,
		"lower":      strings.ToLower,
		"title":      strings.Title,
		"trim":       strings.TrimSpace,
		"trimPrefix": strings.TrimPrefix,
		"trimSuffix": strings.TrimSuffix,
		"contains":   strings.Contains,
		"hasPrefix":  strings.HasPrefix,
		"hasSuffix":  strings.HasSuffix,
		"join":       strings.Join,
		"replace":    strings.ReplaceAll,
		"truncate": func(s string, length int) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},
		"substr": func(s string, start, end int) string {
			if start < 0 {
				start = 0
			}
			if end > len(s) {
				end = len(s)
			}
			if start > end {
				return ""
			}
			return s[start:end]
		},

		// --- Math helpers ---
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"mod": func(a, b int) int { return a % b },
		"gt":  func(a, b float64) bool { return a > b },
		"lt":  func(a, b float64) bool { return a < b },
		"eq":  func(a, b interface{}) bool { return a == b },
		"ne":  func(a, b interface{}) bool { return a != b },
		"gte": func(a, b float64) bool { return a >= b },
		"lte": func(a, b float64) bool { return a <= b },

		// --- Slice / iteration helpers ---
		"makeRange": func(min, max int) []int {
			a := make([]int, max-min+1)
			for i := range a {
				a[i] = min + i
			}
			return a
		},
		"first": func(arr []interface{}) interface{} {
			if len(arr) > 0 {
				return arr[0]
			}
			return nil
		},
		"last": func(arr []interface{}) interface{} {
			if len(arr) > 0 {
				return arr[len(arr)-1]
			}
			return nil
		},
		"len": func(arr interface{}) int {
			switch v := arr.(type) {
			case string:
				return len(v)
			case []interface{}:
				return len(v)
			case []string:
				return len(v)
			case []int:
				return len(v)
			default:
				return 0
			}
		},

		// --- Time helpers ---
		"now":      time.Now,
		"format":   func(t time.Time, layout string) string { return t.Format(layout) },
		"addDays":  func(t time.Time, d int) time.Time { return t.AddDate(0, 0, d) },
		"addHours": func(t time.Time, h int) time.Time { return t.Add(time.Duration(h) * time.Hour) },
		"weekday": func(day int) string {
			days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
			if day >= 0 && day < len(days) {
				return days[day]
			}
			return ""
		},
		"year":  func(t time.Time) int { return t.Year() },
		"month": func(t time.Time) time.Month { return t.Month() },
		"day":   func(t time.Time) int { return t.Day() },

		// --- HTML helpers ---
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		"safeAttr": func(s string) template.HTMLAttr { return template.HTMLAttr(s) },

		// --- Boolean helpers ---
		"not": func(b bool) bool { return !b },
		"and": func(a, b bool) bool { return a && b },
		"or":  func(a, b bool) bool { return a || b },
	}
}
