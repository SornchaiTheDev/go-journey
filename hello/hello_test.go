package main

import "testing"

func TestHello(t *testing.T) {
	t.Run(
		"saying hello to Person Name", func(t *testing.T) {

			actual := Hello("Chokun", "")
			expected := "Hello Chokun"

			assertTestError(t, actual, expected)
		})
	t.Run(
		"saying Hello World when no argument pass",
		func(t *testing.T) {
			actual := Hello("", "")
			expected := "Hello World"

			assertTestError(t, actual, expected)
		})
	t.Run(
		"saying Hello In Thai language",
		func(t *testing.T) {
			actual := Hello("โชกุน", "th")
			expected := "สวัสดี! โชกุน"

			assertTestError(t, actual, expected)
		})
	t.Run(
		"saying Hello In Spanich language",
		func(t *testing.T) {
			actual := Hello("Chokun", "es")
			expected := "Hola, Chokun"

			assertTestError(t, actual, expected)
		})
}

func assertTestError(t testing.TB, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
