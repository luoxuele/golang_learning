package main

import "log"

func main() {

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
