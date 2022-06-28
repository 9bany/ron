package cmd

func randomConfigForNode() *Configuration {
	return &Configuration{
		RootPath:        "./",
		ExecPath:        "index.js",
		Language:        "node",
		WatchExtensions: []string{"js", "ts"},
		IgnorePath:      []string{"./tests/"},
	}
}

func randomConfigForGo() *Configuration {
	return &Configuration{
		RootPath:        "./",
		ExecPath:        "main.go",
		Language:        "go",
		WatchExtensions: []string{"go"},
		IgnorePath:      []string{"./tests/"},
	}
}
