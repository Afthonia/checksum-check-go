# Checksum-Check-Go Docs

## Operating System: Ubuntu 22.04.3 LTS (Jammy Jellyfish) --Linux

First of all, clone the repository with the following command into your directory:

`git clone https://github.com/Afthonia/checksum-check-go.git`

After cloning, in the directory of the project, there are two ways to run the program:

##### First Way

You can just type the following command with the command line flags that are declared for this project:

`go run main.go`


##### Second Way (This way is recommended)

Compile and then run the binary directly with the command below and command line flags that is declared for this project:

```go
go build main.go

./main
```

### Command Line Flags (Arguments)

The command line flags and the information about them such as their usages or types are as follows:

```go
  -check
        states if the checkChecksum function will be run or not
  -checksum string
        given checksum file path
  -file string
        the file whose checksum will be created
  -n int
        bit size of the checksum that will be made (16, 32 or 64)
```

This information can be reached with a default flag as well:

`-h    --help`

The command line flags can be used by the following ways:

```go
-check=        --check=
-checksum=     --checksum=    -checksum    --checksum
-file=         --file=        -file        --file
-n=            --n=           -n           --n
```

-check flag cannot be written with a space between the value and the flag name while others can. Because with boolean flags, the given value can be miscalculated and results a confusion so please make not to use statements like "-check false" or "--check true".

The examples of the correct way of running the program after compiling is as followings:

```go
./main -file <file-path> -n=<bit-size>

./main -file <file-path> -n <bit-size>

./main -file <file-path> -check=true -checksum <checksum-file-path>

./main -file=<file-path> -check=true -checksum=<checksum-file-path>

.
.
.

```

After running, the output will be either a binary of the checksum of the file and a txt file to store it or a boolean value to state that the provided checksum file equals to the on that the program will be calculate.
