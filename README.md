# aws-sdk-experiments

The purpose of this project is to make it easy to generate tiny executables for testing the AWS v2 go SDK. To build all of the executables in the build/ directory, run:
```bash

make build_commands
```


To create a new executable, create a new directory in cmd/  and add a main.go. You can copy the contents of cmd/listS3Contents and strt from there

You also need to add the new diretory to the makefile:
```makefile
CMDS := listClusterArns listS3Contents
```