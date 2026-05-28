# gosh

A Unix shell written in Go. Built as part of a 12-week project challenge — one project per week, finished and shipped.

---

## Build

```bash
go build -o gosh ./cmd/gosh
```

## Run

```bash
./gosh
```

---

## Built-in commands

| Command | Description |
|---|---|
| `cd [dir]` | Change directory. No argument changes to home |
| `pwd` | Print current working directory |
| `echo [args]` | Print arguments to stdout |
| `exit [code]` | Exit with given code, 0 by default |
| `help` | List built-in commands |

---

## Supported syntax

| Syntax | Description |
|---|---|
| `cmd1 \| cmd2` | Pipe output of cmd1 to input of cmd2 |
| `cmd > file` | Redirect stdout to file, overwriting |
| `cmd >> file` | Redirect stdout to file, appending |
| `cmd < file` | Redirect file to stdin |
| `cmd1 \| cmd2 > file` | Pipe output and redirect final result to file |

**Note**: An arbitrary number of commands can be piped together. 

---

## Known limitations

- No tab completion
- No command history
- No shell scripting, conditionals, or loops
- No background jobs or job control
- No environment variable assignment
- No glob expansion
- `cd` does not support absolute paths
- Built-in commands do not support pipes or redirection