package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	print(
		`
░██████╗░░█████╗░░░░░░░░██╗░░░░░░░██╗░█████╗░████████╗░█████╗░██╗░░██╗
██╔════╝░██╔══██╗░░░░░░░██║░░██╗░░██║██╔══██╗╚══██╔══╝██╔══██╗██║░░██║
██║░░██╗░██║░░██║█████╗░╚██╗████╗██╔╝███████║░░░██║░░░██║░░╚═╝███████║
██║░░╚██╗██║░░██║╚════╝░░████╔═████║░██╔══██║░░░██║░░░██║░░██╗██╔══██║
╚██████╔╝╚█████╔╝░░░░░░░░╚██╔╝░╚██╔╝░██║░░██║░░░██║░░░╚█████╔╝██║░░██║
░╚═════╝░░╚════╝░░░░░░░░░░╚═╝░░░╚═╝░░╚═╝░░╚═╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝
`)
	path, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	// nodemon --exec "go build -o bin.exe $pathBase && $pathBase\bin.exe" --ignore "*.exe", "$path\.git\*" -e go
	command := fmt.Sprintf(`nodemon --exec "go build -o bin.exe %v && %v\bin.exe" --ignore "*.exe", "%v\.git\*" -e go`, path, path, path)
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
