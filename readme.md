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

* Go has 2 basic data structs to handle list s of records:
	* Array: fixed length list of things
	* Slice: an array that can grow or shrink
* Slices and arrays must have a defined data type (like in Java or C). So no mix type arrays.
* in *main.go* we need to find a way to create multiple cards at a time.
* we create a new variable called cards. we use the := syntax and declare it as a slice
* the := syntax of declaring a slice is: `name := []type{element1,element2}` e.g `cards := []string{"Card1","Card2"}` declares a slice of type string named cards and initializes it with two strings. instead of literals we can pass in function calls or vars
* how do add new elements to a slice? The syntax is:  `cards = append(cards, "New Card")` or a general rule could be `slice = append(slice, newElement)`
* append adds an element to the end of the slice and returns  A NEW slice so essentialy does not modify the existing slice. but the assignement we use is what modifies the originam slice we pass in the append func
* how do we iterate over a slice?

```
for i, element := range slice {
	fmt.Println(i, element)
}
```

* index: index of the element in the array
* element: current element we are iterating over
* range slice: take the slice and loop over it
* fmt.Println(i, element): the code that runs on every iteration

* in the for loop syntax we kind of redeclare index and element for each iteration with teh assignement := . this is fien as in each iteration we discard the previous values
* in go all declared vars must be used to compile , even indexes

### Lecture 17 - OO APproach vs GO approach

* Go is NOT an OO language. there is no notion of classes in go
* In an OO lang we would build a deck of cards by creating a Deck class. with methods like print(),shuffle() and saveToFile() and an attribute cards to be an array of strings. We would then create a Deck Instance and use this object in our code
* In GO the way to do it is very different. we work with extending basic types (e.g string,int,float, array, map)
* what we want to achieve is to extend a base type and add extra functionality to it. e.g we declare in Go an array of strings `type deck []string` and add a bunch of functions specificllly made to work with it.
* These are functions with deck as *receiver*. A function with a receiver is like a method. a functions that belongs to the instance
* we will createa separate file for our class like new type to describe what it is and how it works
* this new *deck.go* file recides in teh same folder as main.go and belongs to same package (package main)

### Lecture 18 - Custom Type Declarations

* we declare the new type `type deck []string` as a slice of type string
* the immediate effect is that we can replace the `cards := []string{"Ace of Diamonds", newCard()}` in main.go with `cards := deck{"Ace of Diamonds", newCard()}`
* to get deck.go in mu runtime and not get error of undefined: deck i add deck.go in my go run command ` go run main.go deck.go`
* so far type deck has no actual value
* we add a method to id (receiver function)

```
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

* this implements the for loop we had in main. it is called as a receiver because it is bound with the deck type. this binding is done with this new syntax style of (varname type) before the functioname.
* we call this function in the main as a method of the cards variable which was declared of type deck (which has the method). the way to call it is `cards.print()` in a oo style call.

### Lecture 19 - Receiver Functions

* the `(internalvar mytype)` is the reciever on a function `func (internalvar mytype) myfunc() {}`. is what makes the function a receiver function. what is means is that any variable in our app of type mytype has access to this myfunc method. 
* what reciever does behind the scens is to make avalable the copy of the mytype we are working with available in the myfuc method as a variable called internalvar. also every variable of type mytype can call this function on itself
* the internalvar is like this in JS. In Go the is a convention to name the internal reference to the reciver type with 1 or 2 leter abbreviation of the type name

### Lecture 20 - Creating a New Deck

* we want to have a method to create anew deck of cards. this in OO would be done witha constructor. but we dont have constuctors. so we can create afunction that re turns the deck type and build the deck in there. 
* we start by making a new deck var as an empty slice `cards := deck{}`
* we want to populate this slice with data . we choose to do it with 2 nested for loops

```
cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

for i, suit := range cardSuits {
	for j, value := range cardValues {
		cards = append(cards, suit+" of "+value)
	}
}
```

* compiler complains that we declared the indexes i and j and never used them. to make go compiler stop complaining we replace the var name with _. this tells go that this var is not meant ot be used.

* in main.go we replace the old code of instantiating the cards var manually with `cards := newDeck()` and use our deck generator

### Lecture 21 - Slice Range Syntax

* we will go and implement the deal function. we will take an existing deck of cards, take out a number of cards to deal out and create anew deck of the specified size (remainder cardds?). hands and remainder deck will be of same type (deck). essentialy we need to split up the slice
* slices are zero-indexed. if we want to access a specific elemetn in the slice we use `slice[index]` like in any other lang array. like in  python we can extract a part of the slice with the syntax `slice[startIndexIncluding:stopIndexNotIncluding]` e.gfruits[0:2]. with this we can select a range ina slice.
* if our startIndex is 0 we can ommit it `slice[:stopindex]` same holds for stopindex `slice[startindex:]`
* an easy way to solve our dealing problem is to set a handsize var and set hand = cards[:handsize] restpfdeck = cards[handsize:]

### Lecture 22 - Multiple Return Values

* we implement the deal function

```
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

* and call it 

```
hand, remainignDeck := deal(cards, 5)
```

* what we see here is how to pass arguments in a function (nothing new) and how to return multiple values from a function (this is Unusual for other langs)
* the way to return multiple variables from afunction is to define theirt types in a parentherses after the functionname() and before the brackets.  also we return the variables sepaartade by comma. Order matters and type is static and binding
* we can declare and assign these variables to external vars in our program as we see. using the := synbtax to assing 2 vars separated by comma. Order Matters here as well
* extracting ranges from a slice does not alter it .  We created two new references that point at subsections of the 'cards' slice. We never directly modified the slice that 'cards' is pointing at.

### Lecture 23 - Byte Slices

* [ioutil package doc](https://golang.org/pkg/io/ioutil/)
* we will now move to the saveToFile function. our aim is to take a deck and save its representation to a file on our local drive
* the WriteFile funct from go standard lib *ioutil* takes a filename and a byte slice []byte. this is known as bytestream or data stream in other langs and a permissions flag of os.FileMode
* we need to transform our deck a.k.a []string to a []byte to pass it in the file
* a byteslice is a stream of ascii characters/bytes. so our string is covereted to asn ascii stream
* "hi there!" => [72 105 32 116 104 101 114 101 33] *string to byteslice*

### Lecture 24 - Deck to String

* we will use type conversion. type conversion in go is used to take one type of value and turn it into another
* we do *type conversion* by writing the type that we want+parenthesis+in the parenthesis the value that we have. It is more Value casting not real type casting. `[]byte("Hi there!")`
* our deck is actually a slice of strings. we have to flatten it out into a string (using join strings)and the covert it to a byteslice
* we will add a helper function that takes a deck(stringslice) and turns it into a string `func (d deck) toString() string {}`
* conceptual it is preferable to be formed as a receiver function (accessory to the deck type)

### Lecture 25 - Joining a Slice of Strings
	
* [strings lib doc](https://golang.org/pkg/strings/)
* we use `[]string(d)` to type convert the decl to stringslice
* we make use of the *strings* lib to find a readymade Join func we can use. this func takes as argument a stingslice and a char to interpolate between elements while forming the joined string

```
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

* the way to import multiple packages/libs in out source file is

```
import (
	"fmt"
	"strings"
)
```

### Lecture 26 - Saving Data to the Hard Drive

* our save to file method signature looks like this

```
func (d deck) saveToFile(filename string) error {

}
```

* it is a receiver method as we want to attach it to our deck. also it wraps the ioutil WriteFile method so it needs to porvide it with a filename (string). ALso the function it wraps produces an error type. we return it as well as we dont want to implkement error handling there.
* the completed function looks like

```
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```

* as a last param WriteFile needs `os.FilePerm` file permissions (UNIX) so we pass a default of 0666.
* we test our func in main. SUCCESS.

### Lecture 27 - Reading fromn the Hard Drive

* we will make use ReadFIle func from ioutils.  reviewing its signature we see it returns a byteslice AKA string AND and error type 
* so we need to revverse what we did splitting the byteslice on commas et all.
* we name our func newDeckFromFile and we dont attach it as a  receiver. its like newDeck() so it will creatrea a deck from afile.
* we expect a second return value from ReadFile of type error. if all is ok err will have a value of *nil*, if there is an error it will have an error. a common pattern in go is to check for nil on error immediately after returning from the func that might produce it.
* if statement in go is like any other language w/ bracketrs 

```
if err != nil {
	
}
```

* error handling in go is up to programmer. apply common sense and ituition, log the rror
* we make use of Exit() function from *os* package to quit our program. we can pass an error code there (like C) 0 means all ok 1 means error
* our complete error handling code looks like

```
if err != nil {
	// Option 1: log the error and return a call to newDeck()
	// Option 2: log the error and quit the program
	fmt.Println("Error:", err)
	os.Exit(1)
}
```

### Lecture 28 - Error Handling

* we need to do the reverse flow to create adeck out of a byte slice. []byte => string => string.split(",") => []string => deck
* indeed *string* pkg contains a `func Split(s, sep string) []string` that takes the string, the separator string and returns a []string stringslice
* our complete function looks like

```
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: log the error and return a call to newDeck()
		// Option 2: log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
```

* we test it in main and it works

### Lecture 29 - Shuffling a Deck

* [Intn doc](https://golang.org/pkg/math/rand/#Intn)
* we will implement now the shuffle function. it will take a deck and randomize the order
* standard go lib does NOT have a shuffle a list/slice method. we have to make it

```
for each index,card in cards
	Generate a random number between 0 and len(cards) -1
	swap the current card and the card at cards[randomNumber]
```

* standard lib has a rand method in math pkg. we use the Intn func that has a 0 -n range
* our func shuffle will be a receiver function as it is an accessory method of cards

```
func (d deck) shuffle() {
	for i := range d {
		newPosition  := rand.Intn(len(d) -1)
		d{i], d[newPosition] = d[newPosition], d[i]
	}
}
```

* in the for loop we iterate on the index. we get the new position using the intn rand func and then we swap
* for swaping we use double assignment. a unique syntax of Go.
* we test the func and we see that rerunning gets the same results from rand. thi is because we dont seed it with new num.

### Lecture 30 - Random Number Generation

* random num genrerator is a pseudo random generator. it needs a seed
* in the math/rand doc we see that `type Rand` is a source of random numbers. when we create it using the New() function we pass in a Source type var. this Source type is the seed
* to make a Source type var we use NewSource() function passing in a int64 seed num
* to get a random seed we will use *time* package and the current timestamp `func Now()` and pass it into UnixNano() to get a int64 timestamp (not good for crypto)
* our shuffle method with true randomness

```
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
```

### Lecture 31 - Testing With Go

* Go testing is no RSpec, mocha, jasmin etc.
* to make a test on a source file create anew file ending in <sc name>_test.go `deck_test.go`
* to run all tests in a package, run the command `go test`
* we create a `deck-test.go` file in our package folder and add boilerplate code `package main`

### Lecture 32 - Writing Useful Tests

* how we decide what to test? we can write tests for our functions in deck.go
* to see how it words we choose to test the newDeck() method. its up to us to decide what we want to test in this func. we are interested in 3 things (lenght of slice, first element, last element)
* whnever we identify a function that we want to test out of our code base we are going to create Test<Funcname> for it `TestNewDeck`
* we can make Test function to test multiple closely related functions. the convention is to chaintheyrnames preseded with Test `func TestSaveToDeckandNewDeckFromFile`
* how we write test code? e.g code to make sure that a deck is created with x number of cards. Create a new deck => Write if statement to see if the deck has the right number of cards => if it doesn't tell teh go test handler that sthing went wrong
* writing TestFunctions with capital T is a convention
* we pass in our test fun a test handler. `func TestNewDeck(t *testing.T) {}`
* our test file after successfully adding first test looks like

```
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}
```

* its assertion is in the form of if statement.  if the asertion fails we return the T type var setting the error message with its .Errorf method

### Lecture 33 - Asserting Elements in a Slice

* the second test will make sure that eh te first card is an Ace of Spades: Create a NEW deck => write is statemetn to see if the first card in the deck is equal to "Ace of Spades" => if it doesn't tell the go test handler that something went wrong
* our remaining 2 assertions (2 more tests)

```
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d)-1])
	}
```

* go test handler is primitive, so it does not return the number of test passeing or failing. just a general message.

### Lecture 34 - Testing File IO

* we will write a test that creates a deck => saves it to a file => will say file created => attempt to load the file => if it crashes then it will not clean up the created file so that we can reporoduce the test.
* in go we have to take care of cleaning up ourselves
* our test will be: testing saveToDeck and newDeckFromFile. the test code wiil: delete any files in current working directory with the name *_decktesting* => create a deck => save to file "_decktesting" => load from file => assert deck length => delete any files in current working directory with the name "_decktesting"
* we find a delete file func in "os" pkg with name Remove()
* our new test function name will be *TestSaveToDeckAndNewDeckFromFile*. this name although long is an exact vcopy of the functions it tests so it helps us inthe long term to  trace bugs
* our complete test

```
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove test file
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
```

### Lecture 35 - Project Review

* Assignemtn solution

```
package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range list {
		if val%2 == 0 {
			fmt.Printf("%v is even \n", val)
		} else {
			fmt.Printf("%v is odd \n", val)
		}
	}
}
```

## Section 4 - Organizing Data with Structs

### Lecture 36 - Structs in Go

* 