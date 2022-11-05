# Solution generator

Software for saving project templates and quickly applying them.

## Features

- Saving a template
- Applying a template
- Deleting a template

Solution generator is a simple cli tool to save your templates and quickly apply your templates to your new projects

## Installation

Solution generator requires [Go](https://go.dev/) v1.19+ to run.

```sh
go install github.com/bblbbl/solution-generator
```

## Usage

The first time you run any command, a .solution-generator folder is created in your home directory, it will store the saved templates

##### Save new solution
----
```
solution-generator save -n example-name
```

This command will save the template in the current directory to the .solution-generator folder

###### Arguments
* -name or -n: specify solution name to save
* -path or -p: specify path to solution for save

##### Apply solution
----
```
solution-generator apply -n example-name
```

This command will apply the selected template to the current directory

###### Arguments
* -name or -n: specify solution name to apply
* -path or -p: specify path to apply solution
* -dir or -d: specify dir name for applying template

##### Delete solution
----
```
solution-generator delete -n example-name
```

This command will remove the selected template

###### Arguments
* -name or -n: specify solution name to delete

## License

MIT