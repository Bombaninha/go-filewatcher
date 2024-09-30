package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func main() {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	fmt.Println("What path do you want to watch?")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	path := scanner.Text()
	fmt.Println()

	if err := filepath.Walk(path, watchDir); err != nil {
		fmt.Println("Error:", err)
	}

	done := make(chan bool)

	eventMessages := map[fsnotify.Op]string{
		fsnotify.Create: "created file:",
		fsnotify.Write:  "modified file:",
		fsnotify.Remove: "removed file:",
		fsnotify.Rename: "renamed file:",
		fsnotify.Chmod:  "permission changed:",
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				for eventType, message := range eventMessages {
					if event.Has(eventType) {
						fmt.Println(message, event.Name)
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
