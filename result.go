package vsdep

// Result stores information about which solutions needs to build and
// which tests needs to run.
type Result struct {
	Solutions map[string]string `json:"solutions,omitempty"`
	Tests     map[string]string `json:"tests,omitempty"`
}
