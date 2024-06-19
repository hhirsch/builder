# Builder
Write code to commission GNU/Linux systems. For example install packages and copy files.
At the core it is designed to be run on a client system so for your CI pipelines you might
find other tools more suitable.

## Features
- Domain Specific Language inspired by Basic
- Command Line Interface
- Commands to install packages, copy files and take snapshots of the installed packages

## Example
```
// tools.bld
// prompt the user to give information about the host
setupHost myhost
step Install Net Tools If Not Installed
ensurePackage net-tools
```

Run it with the following command
```
builder script tools.bld
```

## Help
To read the help page, use the following command:
```
builder help
```

## Development
To make your development binary available in your system as bdev you can use the following target.
```
make linkBinary
```
The binary is called bdev so it can be differentiated from your regular builder install.
