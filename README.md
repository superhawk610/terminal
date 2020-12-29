# terminal

This package contains a handful of utility functions for interacting with the terminal.

## Migrating from `v0.0.1`

The initial release of this package performed all initialization in `init()`, assuming
that only a single terminal was available and providing no ability to disable initialization
(for example, during tests or debugging).

An example usage prior to `v0.1.0` may have looked something like

```go
import (
  "github.com/superhawk610/terminal"
)

func main() {
  terminal.ClearLine()
  terminal.Overwritef("Hello, %s", "world")
}
```

To perform the same operations post-`v0.1.0`, simply add an explicit `Terminal` initialization:

```go
import (
  "github.com/superhawk610/terminal"
)

func main() {
  t := terminal.New()

  t.ClearLine()
  t.Overwritef("Hello, %s", "world")
}
```

If you'd like to emulate the old behaviour, you can simply initialize the `Terminal` in an `init()`
block in your own code.

## Changelog

See [CHANGELOG.md](CHANGELOG.md).

## License

Copyright &copy; 2019 Aaron Ross, all rights reserved. View the ISC license [here](LICENSE).
