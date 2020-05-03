# sizo-db 


## Overview
`sizo-db` is a CLI tool and a library to manipulate Sizo DB.

### Library
Sizo uses `sizo-db` internally to manipulate vulnerability DB. This DB has vulnerability information from NVD, Red Hat, Debian, etc.

### CLI
`sizo-db` builds vulnerability DBs on GitHub Actions and uploads them to GitHub Release periodically.

```
NAME:
   sizo-db - Sizo DB builder

USAGE:
   main [global options] command [command options] image_name

VERSION:
   0.0.1

COMMANDS:
     build    build a database file
     upload   upload database files to GitHub Release
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
