package main

import "fmt"

const (
	helloPrefixEN = "Hello "
	helloPrefixTH = "สวัสดี! "
	helloPrefixES = "Hola, "
)

func getPrefixByLanguage(language string) (prefix string) {
	switch language {
	case "th":
		prefix = helloPrefixTH
	case "es":
		prefix = helloPrefixES
	default:
		prefix = helloPrefixEN
	}
	return
}

func Hello(name string, lang string) string {
	prefix := getPrefixByLanguage(lang)

	if name == "" {
		name = "World"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chokun", ""))
}
