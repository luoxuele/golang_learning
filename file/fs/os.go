package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	getenv()
}

func play01() {
	envs := os.Environ()
	for _, env := range envs {
		// fmt.Println(env)
		cache := strings.Split(env, "=")
		fmt.Printf(`
		key: %s value: %s
		`, cache[0], cache[1])
	}
}

func getenv() {
	str := os.Getenv("PATH")
	fmt.Println(str)
}
