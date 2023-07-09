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
 cd /my/parent/repo_dir
    $ go work init     // One time thing

Every time you create a new project inside the repo. 
Run this command from cd /my/parent/repo_dir
    $ go work use <PROJECT_FOLDER_NAME>
```