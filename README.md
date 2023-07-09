# GO with examples

Every folder is an app with focus on a sepecific feature.

## Commands
To initialize a module in a go application:
- ` $ go mod init example/goroutines`

To run a go application:
- ` $ go run .`

----
To work with multiple go applications within one repo, first create and populate the file by executing go work:
````
 cd /my/parent/dir
 go work init
 go work use project-one
 go work use project-two
```