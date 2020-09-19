package main

type MyError struct {
	code    string
	message string
}

func (m *MyError) Error() string {
	return "Custom Error: " + m.code + " - " + m.message
}

func executeCustomError() (string, error) {
	return "message ok", &MyError{"04", "the end?"}
}
