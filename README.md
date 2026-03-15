# loglinter

### Linter for Golang that checks log messages for the following rules:
- message should begin with a lowercase letter
- message should be in English (Latin alphabet only)
- message should not contain special characters and emojis
- message should not contain sensitive data (keywords: password, token, etc.)

Supported loggers: `log/slog` and `go.uber.org/zap`.

## Building:
```
git clone https://github.com/ADM307/loglinter.git
cd loglinter
go mod tidy
```
Get binary utility  
`go build -o loglinter ./cmd/loglinter`  
Get binary plugin (optional)  
`go build -o loglinter-plugin ./cmd/plugin`  

## Usage as golangci-lint plugin
Add the following in the `.golangci.yml` in your project:
```
version: "2"
linters:
  settings:
    custom:
      loglinter:
        path: github.com/ADM307/loglinter/cmd/plugin@latest
        description: Checks log messages for style and security.
  enable:
    - loglinter
```
Then run with `golangci-lint run`

## Usage as standalone binary
Running on all package files:  
`loglinter ./...`  
Running on a specific file:  
`loglinter <filename>.go`  

## Examples
| Code | Messages |
| --- | --- |
| `slog.Info("Starting server")` | log message should start with a lowercase letter |
| `slog.Error("Ошибка")` | log message should use only English letters<br>log message should start with a lowercase letter |
| `slog.Info("server started!")` | log message should not contain special characters |
| `slog.Debug("emoji 😀")` | log message should not contain emoji |
| `slog.Info("user password is 123")` | log message may contain sensitive data: "password" |
