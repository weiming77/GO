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
