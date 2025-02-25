# Builder
Write code to commission GNU/Linux systems. 
At the core it is designed to be run on a client system so for your CI pipelines you might
find other tools more suitable.

## Features
- Touring incomplete Domain Specific Language inspired by Basic.
- commands to 
  - install packages
  - upload files
  - manage users
  - turn binaries into system services
  - take snapshots of the installed packages
- command line interface
- self documenting commands

## Planned Features
- incremental Updates
- variables
- abstract syntax tree
- if statements
- y/n prompt
- switch statements
- Nginx configuration
- Wiki install
- Wordpress install
- Wordpress backup
- change Wordpress hostname
- MySQL-Dump
- host aliases json file

## Example
tools.bld
```
setupHost myhost
step Install Net Tools If Not Installed
ensurePackage net-tools
```

Run it with the following command
```
builder script tools.bld
```

This will connect via ssh with key authentification and install net-tools on a remote Debian system.
On the first try it will prompt you for the IP or hostname of myhost.
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
