// Hello,world
package hello

import (
	"fmt"
)

// 定义常量，使用 const
const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const chinese = "Chinese"
const chineseHelloPrefix = "你好, "
const spanishHelloPrefix = "Hola, "

// 定义函数，使用 func 函数名字(形参1 形参1的类型,形参2 形参2的类型) 返回值的类型 { 函数代码块 }
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
	// 当给出了函数返回值的变量名称时，直接使用 return 返回
	return
}

func main() {
	fmt.Println(Hello("世界!", "Chinese"))
}
