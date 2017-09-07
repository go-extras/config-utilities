# confutils
This package provides a set of functions to work with config files

## Provided functions

```golang
// Looks given text (given as slice of bytes) for %ENV(VARIABLE),
// where VARIABLE is any possible variable name.
// If environment variable `VARIABLE` exists, replaces the whole
// expression with the variable value (even if it's empty).
// If the environment variable does not exist, keeps the expression as is.
func SubstEnvVars(text []byte) []byte {
    // ...
}
```

```golang
// String version of SubstEnvVars()
func SubstStringEnvVars(text string) string {
    // ...
}
```