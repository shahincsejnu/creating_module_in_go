package testmod

import "fmt"

// Hi returns a just friendly greeting

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}