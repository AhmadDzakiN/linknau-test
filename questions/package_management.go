package questions

/*
Go's package management system revolves around the concept of modules
If you are starting a new Go project, you need to initialize a new module using go mod init {{your project name}} command

Go Modules provide a way to manage dependencies for Go projects. Each Go module is a collection of related Go packages that are versioned together
Modules are stored in version control repositories (like Git) and identified by a module path, which typically looks like a URL
*/

/*
How to import a Go standard library:
Suppose we want print formatted output to the console/terminal
1. import "fmt"
2. fmt.Println("Hello World!")
*/

/*
How to import a third party package in a Go project:
Suppose we want to import goplayground/validator (Use for value validations for structs and individual fields based on tags)
1. You need to import the package/dependency to the go.mod using go get {{path to the third party package}} command -> go get github.com/go-playground/validator/v10
2. Make sure all dependencies related to that package are already written in go.mod file
3. Import the required dependency to your go file using import command (ex. import "github.com/go-playground/validator/v10")
4. Use the required method according to your needs (ex. Struct() method to validate struct based on the tags in each struct field)

if you do not use the third party package anywhere in your project, Go will automatically remove that package/dependency from go.mod file or you can use go mod tidy command
*/
