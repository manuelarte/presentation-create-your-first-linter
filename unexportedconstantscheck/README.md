# Unexported Constant Check

This linter check that the unexported constants starts with `_`.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

**Exception**: Unexported error values may use the prefix `err` without the underscore.

# How to use it

## Standalone CLI

Install it with:

```bash
go install github.com/manuelarte/presentation-create-your-first-linter/unexportedconstantscheck/cmd
```

And run it with:

```bash
unexportedconstantscheck [--fix=true|false] ./...
```

- `fix`: To suggest a fix

## As a golangci-lint plugin

* Create a file `.custom-gcl.yml`
* Add the path to this linter

```yaml
version: v2.6.1
plugins:
  # a plugin from a Go proxy
  - module: "github.com/manuelarte/presentation-create-your-first-linter/unexportedconstantscheck"
    path: "../unexportedconstantscheck"
```

* Create your custom `golangci-lint` binary and run it:

```bash
	golangci-lint custom -v
	./custom-gcl run --fix ./...
```
