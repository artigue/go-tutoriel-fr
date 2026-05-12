package hello

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("World"))
}
