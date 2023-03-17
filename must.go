package must

// Must panics if the function result contains an error.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

//go:generate ./generate.sh ./must_generated.go 1 8
