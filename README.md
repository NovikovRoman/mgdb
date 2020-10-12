# Model Generator DB

> This is a Go application for generating database entity code.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Build](#build)
- [Usage](#usage)
    - [Model Interface](#model-interface)
    - [Model And Repository](#model-and-repository)
    - [Array Structure For JSON Columns](#array-structure-for-json-columns)
    - [Structure For JSON Columns](#structure-for-json-columns)
    - [Migration Tools](#migration-tools)
- [Help](#help)
- [License](#license)

## Build

```shell script
go build -o "mgdb" && sudo mv mgdb /usr/local/bin/
```

## Usage

### Model Interface

Creating a model interface.

```shell script
mgdb i -c
```
Flag `-c` - create with context.
As a result, the `models` directory will be created in which created the [interface.go](docs/interface.md) file.

```shell script
mgdb i -c myInterface [path-to-dir/]mydir
```
As a result, the `mydir` directory will be created in which created the [myInterface.go](docs/myInterface.md) file. Package name `mydir`.

### Model And Repository

Creating a model with repository.

```shell script
mgdb m -c user
```
Flag `-c` - create with context.
The files [user.go](docs/user.md) and [userRepository.go](docs/userRepository.md) will be created in the `models` directory.

```shell script
mgdb m -c user [path-to-dir/]mydir
```
Flag `-c` - create with context.
Will be created in the `mydir` directory. Package name `mydir`.

### Array Structure For JSON Columns

Creating an array structure for a json column.

```shell script
mgdb j stringArray
```

The file [stringArray.go](docs/stringArray.md) will be created in the `models` directory.

### Structure For JSON Columns
Creating a structure for a json column.

```shell script
mgdb j simpleStruct
```

The file [simpleStruct.go](docs/simpleStruct.md) will be created in the `models` directory.

### Migration Tools

Creating migration tools code.

```shell script
mgdb t
```

This will create the directory `mydir`, with a migration tools and [documentation](docs/migrateDocs.md).

```shell script
mgdb t [path-to-dir/]mydir
```

## Help

Any of the options:
- `mgdb -h|--help`

## License
[MIT License](LICENSE) © Roman Novikov