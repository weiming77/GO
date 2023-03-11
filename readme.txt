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


Modules & Packages:
* Create three GO directories: bin, pkg and src in GO workspace.
* Update the said path of User Variables in environment variables section. They are GOPATH for GO workspace, GOBIN for GO compiled binaries.
* GO code is grouped into packages, and packages are grouped into modules.
* A package is a directory of .go files that organize the code into reusable units.
* A module is a collection of GO packages, with dependencies and versioning build-in.
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

package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}


when execute go mod tidy, what you will see in go.mod
// go.mod
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


