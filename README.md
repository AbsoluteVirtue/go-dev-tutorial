# go-dev-tutorial
## installation
https://go.dev/doc/install

        rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

Add /usr/local/go/bin to the PATH environment variable by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
        
        export PATH=$PATH:/usr/local/go/bin

## basics
https://go.dev/doc/tutorial/getting-started

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path:

        go mod init example/hello
Run your code to see the greeting:

        go run .
Use the following command to get a list of the others:

        go help

## packages
Visit https://pkg.go.dev and search for a package.

Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module:

        go mod tidy
Publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is):

        go mod edit -replace example.com/greetings=../greetings

## tests
At the command line in the greetings directory, run the go test command to execute the test.

The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.

        go test -v

## compilation
The go build command compiles the packages, along with their dependencies, but it doesn't install the results.

The go install command compiles and installs the packages.

You can discover the install path by running the go list command, as in the following example:

        go list -f '{{.Target}}'
As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

        go env -w GOBIN=/path/to/your/bin
or
        go env -w GOBIN=C:\path\to\your\bin
Once you've updated the shell path, run the go install command to compile and install the package:
        go install

## links
 
        https://gobyexample.com/
        https://go.dev/blog/declaration-syntax
        https://go.dev/blog/slices-intro
        https://go.dev/blog/defer-panic-and-recover
        https://go.dev/doc/code
        https://go.dev/doc/articles/wiki/
        https://go.dev/doc/effective_go
