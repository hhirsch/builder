# Builder
Write code to commission GNU/Linux systems

## Features
- Domain Specific Language inspired by Basic
- Command Line Interface
- Commands to install packages, copy files and take snapshots of the installed packages

## Example
```
// prompt the user to give information about the host
setupHost myhost
step Install Net Tools If Not Installed
ensurePackage net-tools
```

## Help
To read the help page, use the following command:
```
builder help
```
