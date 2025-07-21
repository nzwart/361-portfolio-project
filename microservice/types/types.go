// File: types/types.go
package types

// Struct for the complete response structure
type PasswordPackage struct {
	Password string           `json:"password"`
	Analysis CompositionStats `json:"analysis"`
	Security SecurityScore    `json:"security"`
	Metadata Metadata         `json:"metadata"`
}

// Struct for password character distribution analysis
type CompositionStats struct {
	TotalLength      int `json:"totalLength"`
	UppercaseCount   int `json:"uppercaseCount"`
	LowercaseCount   int `json:"lowercaseCount"`
	NumberCount      int `json:"numberCount"`
	SpecialCharCount int `json:"specialCharCount"`
}

// Struct for password strength assessment
type SecurityScore struct {
	Score      int    `json:"scoreLevel"`
	Evaluation string `json:"evaluation"`
}

// Struct for information about the service and compliance
type Metadata struct {
	Standard   string `json:"standard"`
	Compliance string `json:"compliance"`
	Version    string `json:"version"`
}
