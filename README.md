# Arguments Fetcher: argsfet

A simple tool to fetch command line arguments.

# Build

```shell
rsrc -manifest app.manifest -o app.syso
go build -ldflags="-w -s" -o argsfet.exe
```

# Usage

Execute

```shell
argsfet.exe <ProcessName>
```

then will show the command line arguments and output a file named `command_line_<ProcessName>.txt`.
There into `<ProcessName>` is the process name which you want to fetch.

# Reference

[build - https://github.com/akavel/rsrc](https://github.com/akavel/rsrc)

[argsfet.dll - https://github.com/Ninohana/argsfet](https://github.com/Ninohana/argsfet)