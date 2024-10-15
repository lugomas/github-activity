package main

import (
	"fmt"
	"os"
	"roadmaps/projects/github-activity/cmd"
)

func main() {
	// Check if the username is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}
	username := os.Args[1]
	cmd.GetEvent(username)

}
