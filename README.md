# CLI-Calculator-2
# Using Go Lang Cobra command line parsing library
Cobra is a Powerful libaray in Go used for buliding modern CLI application. it helps pares arguments, handle subcommands,flags,and genrate helo menus auutomatically.

# Project Structure
cobra-calculator/<br>
|-- go.mod <br>
|-- main.go <br>
|__ cmd/ <br>
    |-- root.go <br>
    |-- add.go <br>
    |__ sub.go <br>
[main.go]- Start the CLI application<br>
[root.go]- Base command (handless all subcommmands)<br>
[add.go]- Subcommand: adds two numbers <br>
[sun.go]- Subcommand: subtracts two numbers<br>

# How to Run It
step 1: initialize module and install cobra <br>
go mod init cobra-calculator <br>
go get github.com/spf13/cobra@latest <br> 

step 2:Run the CLI application <br>
go run main.go add 10 5<br>
#Output: Result: 15<br>

go run main.go sub 10 5 <br>
#Output: Result: 5 <br>
[main.go]- Your program entry point<br>
[add/sub]- Cobra subcommand<br>
[10 5]- Argument passed to the Run()function<br>
[Output]- Parsed, command, and printed by your command<br>

[cobra library]- Handles parsing help messages and structure. 
