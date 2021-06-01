# go-remote-debug
Remote IDEA Debugging for containerized server written in GoLang

**Note**
Pushed IDE files to retain the exact debug configuration used

## Instructions
* Followed [IDEA debugging](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-in-the-docker-container) documentation

## Observations
* Your source project should be under valid GOPATH directory
* Incase of problems with mapping executable to source code lines while debugging, make sure that your executable root path is same as the root path of your source
* Golang version on container and your local machine should be same
