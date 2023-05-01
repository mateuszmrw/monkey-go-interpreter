# Monkey Language Interpreter in Go

The code in this repository contains an interpreter for "Monkey" language, written by following [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball

Monkey is a simple and intuitive programming language designed for learning and understanding the fundamentals of interpreters and compilers. 

## Features

- Lexer for tokenizing Monkey source code
- Parser for generating an Abstract Syntax Tree (AST)
- Evaluator for interpreting the AST and executing the Monkey code
- REPL (Read-Eval-Print Loop) for an interactive programming environment
- Support for integer, boolean, string, array, and hash data types
- First-class functions and closures
- Error handling and custom error messages

## Getting Started

### Prerequisites

To run the Monkey interpreter, you will need:

- [Go](https://golang.org/doc/install) (version 1.16 or later)

### Installation

1. Clone the repository:

```bash
    git clone https://github.com/mateuszmrw/monkey-interpreter-go.git
```

2. Change into the project directory:

``` bash
    cd monkey-interpreter-go
```

3. Build the project:

```bash
    go build -o monkey
```

## Usage

### Running the REPL

To start the REPL and interact with the Monkey interpreter, simply run:
```bash
    ./monkey
```

You will be greeted with a prompt where you can enter Monkey code.
