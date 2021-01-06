package Hello_World

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hello World, %v!", name)
	return msg
}
