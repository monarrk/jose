package main

import (
	"fmt"
	"flag"

	"github.com/hegedustibor/htgo-tts"
	lol "github.com/kris-nova/lolgopher" 
)

const logo = 	
`	  ______________
	 |_____    _____|
	       |  |
	       |  |
	   ____|  |
	  |_______|ose`

func main() {
	var lang string
	flag.StringVar(&lang, "l", "en", "Select the language")
	flag.Parse()

	speech := htgotts.Speech{Folder: "audio", Language: lang}

	if flag.Arg(0) != "" {
		speech.Speak(flag.Arg(0))
		cleanup()
	}

	lol.Printf("%s\n\n", logo)
	fmt.Print("Welcome to the Jose console! What would you like to do?\n 1: Text\n 2: Speech\n > ")
	opt := scanI()
	switch opt {
	case 1:
		for {
			fmt.Print("Type something: ")
			inp := read()
			cmd(inp)
		}

	case 2:
		for {
        lol.Printf("Type something to say: ")
			inp := read()
			speech.Speak(inp)
		}
	}

	cleanup()
}
