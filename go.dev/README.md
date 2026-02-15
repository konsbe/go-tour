# go.dev


### Init Modules

[create go module](https://go.dev/doc/tutorial/create-module)

You can check for packages in go pkg official site [go packages](https://pkg.go.dev/).

```bash
go get rsc.io/quote
```
Install packages from go.mod with:

```bash
go mod tidy
```

## Create workspace

### Initialize the workspace

[multi-modules workspace](https://go.dev/doc/tutorial/workspaces#create_folder)
In the workspace directory, run: `go work init ./<directory>`
The go work init command tells go to create a go.work file for a workspace containing the modules in the ./hello directory.

The go command produces a go.work file that looks like this:
```
go 1.18

use ./<directory>
```

The go directive tells Go which version of Go the file should be interpreted with. It’s similar to the go directive in the go.mod file.

The use directive tells Go that the module in the hello directory should be main modules when doing a build.

So in any subdirectory of workspace the module will be active.

Run the program in the workspace directory¶
In the workspace directory, run:
```
$ go run ./<directory>
output
```

