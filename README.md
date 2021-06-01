# go-remote-debug
Remote IDEA Debugging for containerized server written in GoLang

**Note:**
Pushed IDEA run and debug configurations.

Followed [IDEA debugging](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-in-the-docker-container) documentation


## Instructions
* Start docker on your local machine
* Clone the project under `$GOPATH/src` directory and open it in IntelliJ IDEA
* Run the docker configuration in IDE to start the container
* Run `remote` configuration in IDE to start remote debugging
* To test, put some breakpoints in source code and run request in `.http` file

## Observations
* Your source project should be under valid GOPATH directory
* In case of problems with mapping executable to source code lines while debugging, make sure that your executable root path is same as the root path of your source
* Golang version on server container, and your local machine should be same
