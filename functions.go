package confutils

import (
	"log"
	"os"
	"regexp"
)

// Looks given text (given as slice of bytes) for %ENV(VARIABLE),
// where VARIABLE is any possible variable name.
// If environment variable `VARIABLE` exists, replaces the whole
// expression with the variable value (even if it's empty).
// If the environment variable does not exist, keeps the expression as is.
func SubstEnvVars(text []byte) []byte {
	reg, err := regexp.Compile("%ENV\\(([A-Za-z0-9]+)\\)")
	if err != nil {
		log.Fatal(err)
	}

	result := reg.ReplaceAllFunc(text, func(substr []byte) []byte {
		envvar := reg.FindSubmatch(substr)[1]
		envvarValue, exists := os.LookupEnv(string(envvar))
		if exists {
			// return the substituted value from the environment
			return []byte(envvarValue)
		}

		// return the original value unmodified
		return substr
	})

	return result
}

// String version of SubstEnvVars()
func SubstStringEnvVars(text string) string {
	bytes := []byte(text)

	return string(SubstEnvVars(bytes))
}
