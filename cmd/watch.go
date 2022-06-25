package cmd

// Watch contains all file and extensions to be observed
type Watch struct {
	Files      []string `yaml:"files"`
	Extensions []string `yaml:"extensions"`
}
