package cmd

func randomConfigForNode() *Configuration {
	return &Configuration{
		RootPath: "./",
		ExecPath: "index.js",
		Language: "node",
		Watch: Watch{
			Extensions: []string{"js", "ts"},
		},
		Ignore: Ignore{
			Files: []string{"./tests/"},
		},
	}
}

func randomConfigForGo() *Configuration {
	return &Configuration{
		RootPath: "./",
		ExecPath: "main.go",
		Language: "go",
		Watch: Watch{
			Extensions: []string{"go"},
		},
		Ignore: Ignore{
			Files: []string{"./tests/"},
		},
	}
}
