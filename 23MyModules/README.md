- go.sum file contains the checksum of the package which help to authenticate the pacakge its not malicious one

- package we get from outside will not stored inside the working directory that will be stored under GOPATH='/Users/athavan/go' under canche directory
- indirect comment in the go.mod file will be removed once we have started using that particular package
- `go mod verify` use the go.sum file and verify all the packages are safe to use
- `go list` list all the modules which are used in our application `go list all` will display all the modules irrespective of we are using it or not
- `go list -m -versions github.com/gorilla/mux` this will list all the version relased by the mux module
- `go mod tidy` go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

- `go mod why github.com/gorilla/mux` list where this module is used
- `go mod edit -go 1.16` used to change the go version (but even we can mannually change this in the go.mod file)
- `go mod edit -module myModules` used to change the module name (but even we can mannually change this in the go.mod file)
- `go mod vendor` this will bring the packages from the local instead of from web
- After using `go mod vendor` to run the application we have to use `go run -mod=vendor main.go`. This is to tell the application to look the local vendor for pacakges
