# Go Essentials

Values, Basic Types, & Core Language Features

- Understanding the **Key Components** of a Go Program
- Working with **Values** & **Types**
- Creating & Executing **Functions**
- Controlling Execution with **Control Structures**

## Understanding the Key Components of a Go Program

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello, world!")
}
```

- **Package Declaration**: `package main`
- **Import Declaration**: `import "fmt"`
- **Function Declaration**: `func main() { ... }`
- **Function Call**: `fmt.Print("Hello, world!")`

### Package Declaration

- The `package` keyword is used to declare the package to which the file belongs.
- The `main` package is a special package that tells the Go compiler to create an executable file.
- The `main` package must contain a `main` function.
- When writing Go code, you split your code into packages.
- You must have one `main` package in your program.

### Import Declaration

- The `import` keyword is used to import packages.
- The `fmt` package is used to format input and output.
- The `fmt` package is part of the Go standard library.

### Function Declaration

- The `func` keyword is used to declare a function.
- The `main` function is a special function that is called when the program is executed.
- The `main` function does not take any arguments.

### Main Function

- The `main` function is the entry point of a Go program.
- The `main` function is called when the program is executed.
- The `main` function must be declared in the `main` package.
