# CLI-Calculator-2
# Using Go Lang Cobra command line parsing library
Cobra is a Powerful libaray in Go used for buliding modern CLI application. it helps pares arguments, handle subcommands,flags,and genrate helo menus auutomatically.

# Project Structure
cobra-calculator/<br>
|-- go.mod
|-- main.go
|__ cmd/
    |-- root.go
    |-- add.go
    |__ sub.go
[main.go]- Start the CLI application
[root.go]- Base command (handless all subcommmands)
[add.go]- Subcommand: adds two numbers 
[sun.go]- Subcommand: subtracts two numbers

# How to Run It
step 1: initialize module and install cobra
go mod init cobra-calculator
go get github.com/spf13/cobra@latest

step 2:Run the CLI application
go run main.go add 10 5
#Output: Result: 15

go run main.go sub 10 5
#Output: Result: 5
[main.go]- Your program entry point
[add/sub]- Cobra subcommand
[10 5]- Argument passed to the Run()function
[Output]- Parsed, command, and printed by your command

[cobra library]- Handles parsing help messages and structure. 
