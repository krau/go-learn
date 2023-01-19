package hello

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const chinese = "Chinese"
const chineseHelloPrefix = "你好, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("世界!", "Chinese"))
}
