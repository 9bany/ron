package cmd

// Watch contains all file and extensions to be observed
type Watch struct {
	Extensions []string `yaml:"extensions"`
}
