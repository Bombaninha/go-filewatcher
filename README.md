# go-filewatcher

This is a Go-based file watcher program that monitors changes in a specified directory using the [fsnotify](https://github.com/fsnotify/fsnotify) package. 

It detects events such as file creation, modification, deletion, renaming, and permission changes in real-time.

## Features

- Monitors changes in files and directories.
- Detects the following file system events:
  - File creation
  - File modification
  - File deletion
  - File renaming
  - File permission changes

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Bombaninha/go-filewatcher.git
   cd go-filewatcher
   ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run cmd/cli/main.go
    ```
