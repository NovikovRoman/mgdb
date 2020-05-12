# Model Generator DB

> This is a Go application for generating database entity code.

## Table of Contents

- [Model Generator DB](#model-generator-db)
  - [Table of Contents](#table-of-contents)
  - [Build](#build)
  - [Usage](#usage)
    - [Model Interface](#model-interface)
    - [Model And Repository](#model-and-repository)
    - [JSON Column Structure](#json-column-structure)
  - [Help](#help)

## Build

```shell script
go build -o "mgdb" && sudo mv mgdb /usr/local/bin/
```

## Usage

### Model Interface

Creating a model interface.

```shell script
mgdb i
```
As a result, the `models` directory will be created in which created the [interface.go](samples/interface.md) file.

```shell script
mgdb m myInterface [path-to-dir/]mydir
```
As a result, the `mydir` directory will be created in which created the [myInterface.go](samples/myInterface.md) file. Package name `mydir`.

### Model And Repository

Creating a model with repository.

```shell script
mgdb m user
```

The files [user.go](samples/user.md) and [userRepository.go](samples/userRepository.md) will be created in the `models` directory.

```shell script
mgdb m user [path-to-dir/]mydir
```
Will be created in the `mydir` directory. Package name `mydir`.

### JSON Column Structure

Creating a structure for a json column.

```shell script
mgdb j stringArray
```

The file [stringArray.go](samples/stringArray.md) will be created in the `models` directory.

```shell script
mgdb m user [path-to-dir/]mydir
```
Will be created in the `mydir` directory. Package name `mydir`.

## Help

Any of the options:
- `mgdb`
- `mgdb h`
- `mgdb -h`
- `mgdb help`
- `mgdb -help`
- `mgdb --help`
 