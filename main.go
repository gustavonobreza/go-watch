package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func cls() {
	print("\033[H\033[2J")
}

func banner() {
	print(
		`
░██████╗░░█████╗░░░░░░░░██╗░░░░░░░██╗░█████╗░████████╗░█████╗░██╗░░██╗
██╔════╝░██╔══██╗░░░░░░░██║░░██╗░░██║██╔══██╗╚══██╔══╝██╔══██╗██║░░██║
██║░░██╗░██║░░██║█████╗░╚██╗████╗██╔╝███████║░░░██║░░░██║░░╚═╝███████║
██║░░╚██╗██║░░██║╚════╝░░████╔═████║░██╔══██║░░░██║░░░██║░░██╗██╔══██║
╚██████╔╝╚█████╔╝░░░░░░░░╚██╔╝░╚██╔╝░██║░░██║░░░██║░░░╚█████╔╝██║░░██║
░╚═════╝░░╚════╝░░░░░░░░░░╚═╝░░░╚═╝░░╚═╝░░╚═╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝
`)
}

func printWaiting() {
	println("Watching for changes...")
}

func printSetup() {
	println("Preparing and starting...")
}

func main() {
	cls()
	banner()
	printWaiting()
	path, err := os.Getwd()
	if err != nil {
		print(path)
		panic(err.Error())

	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	Run(path)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					// log.Println("modified file:", event.Name)
					separedByDot := strings.Split(event.Name, ".")
					if separedByDot[len(separedByDot)-1] == "go" {
						// Is a go file
						cls()
						banner()
						printSetup()
						Run(path)
						printWaiting()
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
func Run(path string) {
	command := fmt.Sprintf(`go build -o bin.exe %v; %v\bin.exe`, path, path)
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
