- Mod file






o program can be of two types

 1 . Executable -  run it in our own system 
 2 , Library/ Package - imported libraries




go mod file Conitains - package name , versoion of go , external dependencies


main.go 

models , package

package are the way to define our files todiff diff packages
every file tht ends with .go will always contains one package and function 


fmt - external library that contains string manupulation related stuff 


functions are the way to group logically related code/expression 


# GO Build  

go support UTF-8  - unicode char - 120thousand char almost all

Meanwhile Others Support only - ASCII - like ony az AZ 0-9  - 127 

- Go Provides cross platform Single binry which means a Go program can generate a Executable binary 

CMD for generate /binary  - GOOS=windows GOARCH=amd64 go build .

during build it does not cache anything nor store anything  - so if you run go run thne it use cache in local 

# Package

- match to the folders name
- math/rand
- we can follow standard way of import packages 

-  if we are in diff packages and we want to access diff package from diff library the name will always in Capital letter are accessable : Encapsulation


# Functions

statically typed language

1 implicit return 
2 explicit return 


# Variables 

if we have vriable of same type then we add thetype at the end  : 
all the varable initialiized with 0 state in go a variable nevwr initialized in null state 


Rules of Vairbale :
- if initialized valur not matter  var i int 
- if we know the initial value  know then i:= use this 
- outside the funcition use var 