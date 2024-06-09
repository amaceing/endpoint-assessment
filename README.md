# Endpoint Assessment

This Go program allows you to manage a directory tree structure with commands to create, list, delete, and move directories. The program reads commands from a text file and manipulates the directory tree accordingly.

## Features

- **CREATE**: Create a directory at the specified path.

    ```
    CREATE grains/squash
    ```
- **LIST**: List all subdirectories of the root.

    ```
    LIST
    fruits
    apples
        fuji
    ```
- **DELETE**: Delete a directory and all its subdirectories.

    ```
    DELETE fruits/apples
    Cannot delete fruits/apples - fruits does not exist
    DELETE foods/fruits/apples
    ```
- **MOVE**: Move a directory from one path to another.

    ```
    MOVE vegetables foods
    ```

## Prerequisites

- Go 1.16 or later

## Installation

### Install Go

1. Download the Go installer from the [official Go website](https://golang.org/dl/).
2. Follow the installation instructions for your operating system:
   - [Windows](https://golang.org/doc/install#install)
   - [macOS](https://golang.org/doc/install#install)
   - [Linux](https://golang.org/doc/install#install)

3. Verify the installation by opening a terminal and running:
   ```sh
   go version
   ```
You should see the installed Go version.

### Download and Extract the Project
Download the `endpoint-assessment.zip` file. Extract the contents of the zip file to a directory on your local machine.

### Clone git repository

```
git clone https://github.com/amaceing/endpoint-assessment.git
```

## Usage

### Run the Program
Open a terminal and navigate to the project directory.

Run the program using the following commands:

```
go build
./endpoint-assessment
```

### Test the program

```
go build
go test
```

The program will read commands from the `input.txt` file and output the results to the terminal.


