package executable

// Even though this could be imported, it's forbidden because the package
// has a main package which makes it executable
// Note that if the package was not called main, it would be importable even
// without the main function.

const (
	EXECUTABLE = "executable"
)
