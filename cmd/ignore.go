package cmd

// Ignore contains all files and extensions
// not to be observed during the application cycle
type Ignore struct {
	Files      []string `yaml:"file"`
	Extensions []string `yaml:"extensions"`
}
