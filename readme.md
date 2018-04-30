# Udeny Course - Go: The Complete Developers Guide
[course](https://www.udemy.com/go-the-complete-developers-guide/learn/v4/overview)
[repo](https://github.com/StephenGrider/GoCasts)

## Section 1

### Lecture 3 - Environment Setup

* development environment setup
	* install go
	* install text editor (ok)
	* configure text editor for go
	* start dev
* install go [dowload](https://golang.org/dl/)
* add to .profile the followin

```
export GOROOT=$HOME/go1.X
export PATH=$PATH:$GOROOT/bin
```

* source ~/.profile and set terminal to run as login terminal
* or logout and login again
* test installation running `go` in terminal

### Lecture 5 - Go support in VSCode

* in VSCode: View -> Extensions -> Go 

## Section 2 - A Simple Start

### Lecture 6 - Boring Old Hello World

* we make a folder called *helloworld*
* inside itr we make a *main.go* file
* our code looks like

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi there!")
}

```

### Lecture 7 - Five Important Questions

* How do we run the code in our project? in our project folder wher our source resides we use the go cli and use the run command `go run <sourcefile>.go`. THe supported cli commands are:
	* `go build`: compiles a bunch of go source code files
	* `go run`: compiles and executed one or two go source code files
	* `go fmt`: format all the code in each file in the current directory
	* `go install`: compiles and *installs* a package
	* `go get`: downloads the raw source code of someone else's package
	* `go test`: runs any tests associated with the current project
* if we run the `go build main.go` command we get an *executable* main file.which we can execute our self. this file is not executed with the build command
* go install and go get handle dependencies in our project

### Lecture 8 - Go Packages

* What does `package main` mean?
* when we see the word package we think it refers to a project or a workspace.
* A package is a collection of common source code files
* it can have many related go files inside it.
* the only requirement for every file in a package is that at the very first line of each file it must declare the package it belongs to
* so the file we just wrote belong to the package main
* so why we chose the name main for our project package.
* there are 2 types of packages. 
	* executable: when compiled generates a file that we can run
	* reusable: code used as *helpers*. Good place to put reusable logic. they are libraries or code dependencies
* how do we know if we are in a executable or reusable package? is the name of the package we use that determines that. the word main is the only one that produces an executable package. all other packages when build with `go build ` DO NOT produce an executable
* an other precondition for executable package (main) is that it contains a main function that will be called first. to run our program

### Lecture 9 - Import Statements

* [golang standard lib packages](https://golang.org/pkg/)
* What does `import fmt` mean? import gives our code access to another package , so code writen by someone else.
* fmt is a standard go library package (shorthand for format). contains utils to printout to console ,debugging and the like
* other go standart lib packages are: debug, fmt, math, encoding, crypto, io
* we can import user created libraries as well. (reusable packages)

### Lecture 10 - File organization

* What's the 'func' thing? is shorthand for function. syntax is like most traditional prog langs
* How is the `main.go` file organized?
* a typical go file is structured
	* Package Declaration: e.g `package main`
	* Import other packages that we need: e.g. `import "fmt"`
	* Declare functions, tell go to do things, use imports: 

```
func main() {
	fmt.Println("Hi there!")
}
```

## Section 3 - Deeper into Go

### Lecture 12 - Project Overview

* our project will be a package that will simulate playing around witha  deck of playing cards
* our package will have several function to simulate a card game
	* newDeck: createa list of playing cards, essentially an array of strings. eache element will represent a playimng card
	* print: log out the contents of a deck of cards
	* shuffle: shuffles all the cards in a deck
	* deal: create ahand of cards
	* saveToFile: save a list of cards to a file on a local machine
	* newDeckFromFile: load a list of cards from the local machine 

### Lecture 13 - New Project Folder

* we creates a *cards* folder to host our project
* we create a *main.go* file inside
* we add the boilerplate code from last project

### Lecture 13 - Variable Declarations

* a typical variable declaration in go `var card string = "Ace of Spades"`. nothing special, only special is that datatype is declared and is placed after the var name. *var* declares we are going to createa anew variable, *card* is the name of the var, *string* is the type of the variable *='Hi there'* assigns the value to the variable
* What we see in this declaration is that Go is a statically typed language like Java and C++ unlike dynamic type languages like JS,Ruby and Python
* unlike Java or C++ where we define vars as `int mynum = 1` in go we declare them as `var mynym int = 1`
* the var declaration notation we saw for go is one of the ways we can use. the other is infering the type from the value. (so maybe it will look a bit Javascript like the declaration but is statically typed anyhow - no type reassignement)
* the type inferred version of var declaration or shorthand is `card := 'Ace of Spades`. this is 100% equivalent to the `var card string = "Ace of Spades". The second syntax with := is used only when we first declare the var not in subsequent reassignemtns of values (of same type) then we use just = `card = "Five of Diamonds"`
* Go does not accept both single quotes and doublequotes for string literals. only doublequotes. singlequotes are used only for single chars (like C)
* go basic types include bool,string,float64,int
* we cannot initialize AND assign values to vars OUTSIDE of FUNCIONS
* we can initialize (var varname type) vars and assign values to them (varname = value) afterwords in our functions
*  We can initialize a variable outside of a function, we just can't assign a value to it.
* we cannot assigne values to vars before they are initialized

### Lecture 15 - Functions and Return Types

* in a realistic world i dont know before exeuting the program the card i want so i need to make it parametrical passing the card to main through a function call
* we define a new func *new card* which returns the string. return statement is like any other language

```
func newCard() {
	return "Five of Diamonds"
}
```

* But there is a Caveat the code above throows an error because in go after the functionb declaration we need to declare the returned variable type. the previous code says we will return nothing but we return a string. Go is statically typed lang . in c++ and java functions are defined as `int main() {}` so there as well we define the return type. just the syntax is different

```
func newCard() string {
	return "Five of Diamonds"
}
```

* the syntax of function dewclaration is broken like this:
	`func functionName(argument1, argument2) return var datatype { return variable}`
* inconsistency between return type declaration and actual data returned throws an error
* Files in the same package can freely call functions defined in other files. the build order does not matter

### Lecture 16 - Slices and For Loops

* 