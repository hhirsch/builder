# Builder
Write code to commission GNU/Linux systems. 
At the core it is designed to be run on a client system so for your CI pipelines you might
find other tools more suitable.

## Features
- Touring incomplete Domain Specific Language inspired by Basic
- Commands to 
  - Install packages
  - Upload files
  - Manage users
  - Turn binaries into system services
  - Take snapshots of the installed packages
- Command Line Interface
- Turn your binary into a service for systemd

## Planned Features
- Migration management
- Variables
- Abstract syntax tree
- If statements
- y/n prompt
- Self constructing help
- Switch statements
- Nginx configuration
- Wiki install
- Wordpress install
- Wordpress backup
- Change Wordpress hostname
- MySQL-Dump

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

This will install net-tools on a Debian system
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
