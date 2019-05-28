package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"bufio"
)

func cleanup() {
	fmt.Println("Cleaning up...")
	files, err := ioutil.ReadDir("./audio")
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range files {
		file := "audio/" + f.Name()
		err = os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(0)
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func scanI() int {
	var i int
	fmt.Scanf("%d", &i)
	return i
} 

func cmd(inp string) int {
	switch inp {
	case "exit":
		cleanup()

	case "console":
		main()
	
	default:
		fmt.Printf("I don't currently know how to handle that, sorry.\n")
	}
	return 0
}