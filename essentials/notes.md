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

## Why Main?

- The `main` function is the entry point of a Go program.
- The `main` function is called when the program is executed.
- The `main` function must be declared in the `main` package.
- The `main` function is a special function that is called by the Go runtime.

## Go Types & Null Values

### Basic Types

Go comes with a couple of built-in basic types:

`int` => A number WITHOUT decimal places (e.g., -5, 10, 12 etc)

`float64` => A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)

`string` => A text value (created via double quotes or backticks: "Hello World", `Hi everyone')

`bool` => true or false

But there also are some noteworthy "niche" basic types which you'll typically not need that often but which you should still know about:

`uint` => An unsigned integer which means a strictly non-negative integer (e.g., 0, 10, 255 etc)

`int32` => A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890)

`rune` => An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')

`uint32` => A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295

`int64` => A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

There also are more types like `int8` or `uint8` which work in the same way (=> integer with smaller number range)

### Null Values

All Go value types come with a so-called "null value" which is the value stored in a variable if no other value is explicitly set.

For example, the following `int` variable would have a default value of `0` (because 0 is the null value of int, int32 etc):

`var age int // age is 0`
Here's a list of the null values for the different types:

`int` => 0

`float64` => 0.0

`string` => "" (i.e., an empty string)

`bool` => false

## Ways of Declaring Variables

1. **Type Inference**: Go can infer the type of a variable based on the value assigned to it.

```go
name := "John"
```

2. **Explicit Declaration**: You can explicitly declare the type of a variable.

```go
var age int
```

3. **Multiple Declaration**: You can declare multiple variables at once.

```go
var name, age = "John", 25
```

4. **Zero Value**: Variables that are declared but not initialized are assigned a zero value.

```go
var age int // age is 0
```

5. **Short Declaration**: You can declare and initialize a variable in a single line.

```go
name := "John"
```

6. **Multiple Short Declaration**: You can declare and initialize multiple variables in a single line.

```go
name, age := "John", 25
```

7. **Var Block**: You can declare multiple variables in a block.

```go
var (
  name = "John"
  age  = 25
)
```

## Constants

- Constants are values that do not change.
- Constants are declared using the `const` keyword.
- Constants must be assigned a value when declared.
- Constants cannot be declared using the `:=` syntax.

```go
const name = "John"
```

## Getting User Input

- You can get user input using the `fmt` package.
- The `Scan` function is used to get user input.
- The `Scan` function takes a pointer to the variable where the input will be stored.

```go
var name string
fmt.Scan(&name)
```

In the above code snippet, the `Scan` function takes a pointer to the `name` variable.
This is because the `Scan` function needs to modify the value of the `name` variable.

## Formatting Output and Strings

- You can format output using the `fmt` package.
- The `Print` function is used to print output.
- The `Printf` function is used to format output.
- The `Sprintf` function is used to format output without printing it.
- The `Println` function is used to print output with a newline character.

```go
fmt.Print("Hello, world!")
fmt.Printf("Hello, %s!", name)
fmt.Println("Hello, world!")
```

In the above code snippet, the `Printf` function is used to format the output by replacing `%s` with the value of the `name` variable.

```go
formatted := fmt.Sprintf("Hello, %s!", name)
fmt.Print(formatted)
```

In the above code snippet, the `Sprintf` function is used to format the output without printing it.

```go
fmt.Print(`You can use backticks to print "quotes", 'single quotes', and newlines`)
```

In the above code snippet, backticks are used to print quotes and single quotes. Backticks are also used to print newlines.

## Functions

- Functions are blocks of code that perform a specific task.
- Functions are declared using the `func` keyword.
- Functions can take zero or more arguments.
- Functions can return zero or more values.

```go
func greet() {
  fmt.Print("Hello, world!")
}
```

In the above code snippet, the `greet` function is declared using the `func` keyword.
The `greet` function does not take any arguments and does not return any values.

```go
func greet(name string) {
  fmt.Printf("Hello, %s!", name)
}
```

In the above code snippet, the `greet` function takes a `name` argument of type `string`.

```go
func greet(name string) string {
  return fmt.Sprintf("Hello, %s!", name)
}
```

In the above code snippet, the `greet` function returns a formatted string.

Functions can also return multiple values.

```go
func add(a, b int) (int, error) {
  if a < 0 || b < 0 {
    return 0, errors.New("both numbers must be positive")
  }
  return a + b, nil
}

sum, err := add(5, 10)
if err != nil {
  fmt.Print(err)
}
```

In the above code snippet, the `add` function returns the sum of two numbers and an error if either of the numbers is negative.

Another way of returning multiple values is by using named return values.

```go
func add(a, b int) (sum int, err error) {
  if a < 0 || b < 0 {
    err = errors.New("both numbers must be positive")
    return
  }
  sum = a + b
  return
}

sum, err := add(5, 10)
if err != nil {
  fmt.Print(err)
}
```

In the above code snippet, the `add` function uses named return values to return the sum of two numbers and an error if either of the numbers is negative.

## Control Structures

- Control structures are used to control the flow of a program.
- Go has three main control structures: `if`, `for`, and `switch`.

### If Statement

- The `if` statement is used to execute a block of code if a condition is true.
- The `if` statement can have an optional `else` block.
- The `if` statement can have an optional `else if` block.

```go
if age >= 18 {
  fmt.Print("You are an adult")
} else {
  fmt.Print("You are a child")
}
```

In the above code snippet, the `if` statement checks if the `age` variable is greater than or equal to `18`.

### For Loop

- The `for` loop is used to execute a block of code repeatedly.
- The `for` loop can have an optional `init`, `condition`, and `post` statement.

```go
for i := 0; i < 5; i++ {
  fmt.Print(i)
}
```

In the above code snippet, the `for` loop prints the numbers from `0` to `4`.

In pseudo-code, the `for` loop can be represented as follows:

```
init statement (executed before the loop starts)
while condition statement (evaluated before each iteration of the loop)
  code block
  post statement (executed after each iteration of the loop)
end while
```

The `init` statement is executed before the loop starts.
The `condition` statement is evaluated before each iteration of the loop.
The `post` statement is executed after each iteration of the loop.

Translated to Go code, the `for` loop can be represented as follows:

```go
for init; condition; post {
  // code block
}
```

#### Break and Continue

- The `break` statement is used to exit a loop.
- The `continue` statement is used to skip the current iteration of a loop.

```go
for i := 0; i < 5; i++ {
  if i == 3 {
    break
  }
  fmt.Print(i)
}
```

In the above code snippet, the `break` statement is used to exit the loop when `i` is equal to `3`.

```go
for i := 0; i < 5; i++ {
  if i == 3 {
    continue
  }
  fmt.Print(i)
}
```

In the above code snippet, the `continue` statement is used to skip the iteration when `i` is equal to `3`.

### Switch Statement

- The `switch` statement is used to execute a block of code based on the value of an expression.
- The `switch` statement can have multiple `case` blocks.
- The `switch` statement can have an optional `default` block.

```go
switch day {
case "Monday":
  fmt.Print("It's Monday")
case "Tuesday":
  fmt.Print("It's Tuesday")
default:
  fmt.Print("It's another day")
}
```

In the above code snippet, the `switch` statement checks the value of the `day` variable and executes the corresponding block of code.

## File I/O

- You can read and write files in Go using the `os` and `io/ioutil` packages.
- The `os` package provides functions for file operations.
- The `io/ioutil` package provides functions for reading and writing files.

### Reading Files

- You can read the contents of a file using the `ReadFile` function from the `io/ioutil` package.

```go
data, err := ioutil.ReadFile("file.txt")
if err != nil {
  fmt.Print(err)
}
fmt.Print(string(data))
```

In the above code snippet, the `ReadFile` function reads the contents of the `file.txt` file and stores it in the `data` variable.

### Writing Files

- You can write data to a file using the `WriteFile` function from the `io/ioutil` package.

```go
data := []byte("Hello, world!")
err := ioutil.WriteFile("file.txt", data, 0644)
if err != nil {
  fmt.Print(err)
}
```

In the above code snippet, the `WriteFile` function writes the data to the `file.txt` file.

### Creating Files

- You can create a new file using the `Create` function from the `os` package.

```go
file, err := os.Create("file.txt")
if err != nil {
  fmt.Print(err)
}
defer file.Close()
```

In the above code snippet, the `Create` function creates a new file named `file.txt`.

### Closing Files

- You should always close files after you are done using them.
- You can close a file using the `Close` method.

```go
file, err := os.Open("file.txt")
if err != nil {
  fmt.Print(err)
}
defer file.Close()
```

In the above code snippet, the `Close` method is called on the `file` variable using the `defer` keyword.

## Comma Ok Syntax/Idiom

- The comma ok syntax is used to check if a map lookup was successful.
- The comma ok syntax returns two values: the value and a boolean indicating if the lookup was successful.

```go
ages := map[string]int{
  "John": 25,
  "Jane": 30,
}

age, ok := ages["John"]
if ok {
  fmt.Printf("John's age is %d", age)
} else {
  fmt.Print("John's age is unknown")
}
```

In the above code snippet, the comma ok syntax is used to check if the lookup for the key `"John"` was successful.

## Defer Statement

- The `defer` statement is used to delay the execution of a function until the surrounding function returns.
- The `defer` statement is often used to close files, release resources, or perform cleanup tasks.

```go
file, err := os.Open("file.txt")
if err != nil {
  fmt.Print(err)
}

defer file.Close()
```

In the above code snippet, the `Close` method is called on the `file` variable when the surrounding function returns.

## Panic and Recover

- The `panic` function is used to cause a runtime error.
- The `recover` function is used to recover from a panic.

```go
func divide(a, b int) int {
  if b == 0 {
    panic("division by zero")
  }
  return a / b
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Print("recovered from panic")
    }
  }()

  result := divide(10, 0)
  fmt.Print(result)
}
```

In the above code snippet, the `divide` function panics if the second argument is `0`.
The `recover` function is used to recover from the panic and print a message.
