if you have code base dependancy to other source units in different package/module
you should run go mod command in the main project folder. ie:
go mod init my-first-go-project
PS: 
You can init any main package name. However this is only when project is running alone in local machine. if you decide to deploy it as repo mod, just follow the naming convention <repoURL>/<userName>/<mainPackageName>. ie
go mod init github.com/weiming77/hello-world

when you are perform referencing in main package. just
import "github.com/weiming77/hello-world/greeting
ps: greeting is subfolder name in main package project dir.

when you want to access the method in package just refer the package name and it funtion or object. ie:
greetings.Hello

PS: 
1. greetings is package name, greeting is subfolder that contains packages go files.
2. Variables, Objects or function are accessible by other package need to started with Uppercase. 

Quick Note:
GOROOT, GOPATH & GOBIN
* $GOROOT denotes go installation path

* $GOPATH directory contains all the go code(workspace) in the host which 
defaults to $HOME/GO on unix and %USERPROFILE%/go on Windows.
- Unix: export GOPATH=$(HOME)/go-projects
- Windows: Create new environment variable GOPATH

* $GOPATH directory should have three directories
- src: for source files
- pkg: for compiled file whose suffix is .a
- bin: for executables files

* $GOBIN directory contains all the final executables. Defaults to $GOPATH/bin or to be set.


Common GO commands and its associated:
GO help - Help with description
GO env - list out the GO environment variables
GO version - display the correct version
GO mod init <module-name> - Create a module
GO get <module-name> - Get third party packages from SCM
GO run <filename>
GO build - Compile the executable file and put it in same directory
GO build-O <output-file-name> - Save executable with custom name
GO install - Compile the executable file and move to $GOPATH/bin or $GOBIN if set
echo %GOPATH%
PS: go clean // clean the executable following go build

Modules & Packages:
* Create three GO directories: bin, pkg and src in GO workspace.
* Update the said path of User Variables in environment variables section. They are GOPATH for GO workspace, GOBIN for GO compiled binaries.
* GO code is grouped into packages, and packages are grouped into modules.
* A package is a directory of .go files that organize the code into reusable units.
* A module is a collection of GO packages stored in a file tree with a go.mod at its root, with dependencies and versioning build-in.
* Programs start running in package named main, with functon main.
* go.mod lists out other dependent modules, usually from source code repositories like GitHub.
* The basic go.mod file in the repo assumes your project is hosted on Github, but it is not a requirement.
* https://github.com/golang-standards/project-layout
* https://go.dev/doc/tutorial/getting-started

go.mod vs go.sum
* go.mod and go.sum files usually resides at the root of your project.
* Precise version of a dependency without breaking anything.
* Defines modules import path & project's depenencies requirement and also locks them to their correct version.
* All the dependencies will be downloaded in the $GOPATH/pkg/mod directory with versioning.
* go mod vendor for making a vendor a copy of your dependencies.

package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}


when execute go mod tidy, what you will see in go.mod
// go.mod -- by remove the 3rd packages that no longer required.
// also prepare the dependency that auto install the required 3rd party package(s)
module github.com/weiming77/go-microservices
go 1.17
require rsc.io/quote v1.5.2
require (
	golamg.org/x/text v0.0.0-201709150032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

﻿

* go.sum file contains the cryptographic checksums of the content of specific module versions that you use in your program:
* Each time a dependancy is used, its checksum is added to go.sum, if missing or else required to match the existing entry in go.sum.
* Go maintains a cache of downloaded packages and computes and records the cryptographic checksum of each package at download time.
* In normal operation, the go command checks these pre-computed checksums against the main module's go.sum file, instead of recomputing them on each command invocation.
* So checking in go.sum file can help others to verify that they are using the same modules as you.
//go.sum
golang.org/x/text v8.0.0-20170915032832-14c0c48eadac h1:qgoY5wgZOaTkIIMVji8= golang.org/x/text v8.0.6-20170915032832-14c0c48ead9c/go.mod h1: NCMBEUOU14njkJ3fqMW+ rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/a0+QwRjYZOKNam
rsc.io/quote v1.5.2/go.mod h1:Lzx7hef JvL54yjefDEDHNONDJII
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+ROD
rsc.io/sampler v1.3.0/go.nod h1:T1hPZKmBbMNahiBXFy5HгXp

TIPS:
go get <3rd party package>
* ie go get rsc.io/quote or 
* what you can do via go mod tidy to download the used packages and update dependency in go.mod file automatically.


GO Basics
Data types:
bool

Numeric Types: 
* int8, int16, int32, int64, int
* uint8, uint16, uint32, uint64, uint
* float32, float64
* complex64, complex128

byte: is an alias of uint8

rune: is an alias of int32

string 

Variable & Constants
var x int = 10
var y = 20
var z = 22.2

var x, y, z = 1, 2, 3
var w float64
w = 33.546

var name string = "Hello World"
var isWorking bool = true

const percent float32 = 5.5

Shorthand declaration
x := 10
x, y, z := 1, 2, 3
w := 33.546
name := "Hello World!"
isWorking= false

Fmt Package vs Variable type & size
* The type of a variable can be printed using %T format.
* Go has a package unsafe which has a Sizeof function which returns in bytes the size of the variable passed to it.
* Unsafe package should be used with care as the code using it might have portability issues. ie

// main.go
package main

import ( 
	"fmt"
	"unsafe"
)

func main() {
	x, y, z := 1, 2, 3
w := 33.546
name := "Estella"
fmt.Println("Name is: ", name)
fmt.Printf("Type of name is %T, size of name is %d", name, unsafe.Sizeof(name))
} 

Functions
* Block of code that does a specific task.

Declaration:
func funcName(parameters type) returnType {
	// body
}

Declaration:
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) *2 
	return area, perimeter
}

// or 
func rectProps(length, width float64) (area, perimeter float64) {
	var area = length * width
	var perimeter = (length + width) *2 
	return
}

Invoke:
area, perimeter := rectProps(10, 2)
fmt.Printf("Area %f Perimeter %f", area, perimeter)

PS: make the function name Capital letter if you want to export the function to another package, otherwise keep it lowercase for local package usage.


Maps:
* Maps are like Python dictionary datatype.
* A map is a built-in type in Go which associates a value to a key. The value can be retrieved using the corresponding key.
* If Key is absent, value 0 will be returned.

Syntax: map[key_type]value_type
var daysOfMonth map[string]int
daysOfMonth["April"] = 31

var daysOfMonth = map[string]int{"Jan": 31, "Feb": 28}
daysOfMonth:= map[string]int{"Jan": 31, "Feb": 28}

fmt.Println(daysOfMonth["Feb"])


Importing Local Packages
* There are various way of import local packages. We can
a. Alias
- Import the package name first: geo "github.com/xyz/go-microservices/geometry"
  PS: geo is alias name of geometry package
- then we can call the function from the geometry package: 
	area := geo.Area(3, 3)
	diagonal := geo.Diagonal(3, 3)

b. direct using "." operator
- Bring the complete package into current namespace: . "github.com/xyz/go-microservices/geometry".
- so any function or variable can be direct reference using ".". 
- then we can call the function from the geometry package: 
	area := Area(3, 3)
	diagonal := Diagonal(3, 3)

c. use of blank identifier
- It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so. Same is the case for variables that are defined but not used.
- The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time.
- use "_" in order to avoid unused function or variable removed: _ "github.com/xyz/go-microservices/geometry" 


Cross Compilation
- GO can compile binaries that will work on the oter operating system.
- Go suppports a variety of platform, operating systems and architectures.
- Go tool dist list 
  android/amd64
  android/arm
  android/arm64
  darwin/amd64
  freebsd/amd64
  freebsd/arm
  ios/amd64
  ios/arm64
  linux/386
  linux/amd64
  linux/arm
  linux/arm64
  windows/386
  windows/amd64
  windows/arm
  windows/arm64
- ie
env GOOS=linus GOARCH=arm64 go build -o main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\nArch: %s\n", runtime.GOOS, runtime.GOARCH)
)


Golang Package Structure 
- There are executable and non-executable package. 
- Executable package is the main entry point of the Go application, the package naming usually known as package main, and the main entry point is func main. But the go file not neccessary main.go.
- Non executable package also known as utility package and it is group of functions and procedures.
- The package name must be descriptive and not camel/snake. Usually same as name directory name. The Go unit files are usually small and many.
- The package scope will be break into block {}, exported (Upper case) and non-exported (lower case).
- package structure:
  /app
	|---/handlers
	|	|__ handlers.go {ninja handler, dojo handler}
	|
	|---/models
	|	|__ models.go { ninja struct, dojo struct }
	|	|__ requests.go { requests }
	|	|__ responses.go { responses }
	|
	|---/persistence
	|	|__ persistence.go PS: Relevant to databases
	|
	|---/main
		|__ main.go { main function, router setup, server kickoff }
