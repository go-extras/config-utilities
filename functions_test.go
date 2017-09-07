package strutils

import (
	"os"
	"testing"

	"gopkg.in/go-extras/testing.v1/assert"
)

func TestSubstEnvVars(t *testing.T) {
	testText := []byte("My text %ENV(TEST)\n%ENV(locase)\n%ENV(nonexistant)\n%ENV(ONE)%ENV(TWO)\n%ENV(NON VALID)\n%ENV(ITS_VALID)")
	expectedResult := []byte("My text 424242\nMy test value\n%ENV(nonexistant)\nTWOONE\n%ENV(NON VALID)\nEXISTS")
	os.Setenv("TEST", "424242")
	os.Setenv("locase", "My test value")
	os.Setenv("ONE", "TWO")
	os.Setenv("TWO", "ONE")
	os.Setenv("NON VALID", "WTF")
	os.Setenv("ITS_VALID", "EXISTS")
	result := SubstEnvVars(testText)
	os.Unsetenv("TEST")
	os.Unsetenv("locase")
	os.Unsetenv("ONE")
	os.Unsetenv("TWO")
	os.Unsetenv("NON VALID")
	os.Unsetenv("ITS_VALID")
	assertEqual := assert.Equal(t)
	assertEqual(string(result), string(expectedResult))
}

func TestSubstStringEnvVars(t *testing.T) {
	testText := "My text %ENV(TEST)\n%ENV(locase)\n%ENV(nonexistant)\n%ENV(ONE)%ENV(TWO)\n%ENV(NON VALID)\n%ENV(ITS_VALID)"
	expectedResult := "My text 424242\nMy test value\n%ENV(nonexistant)\nTWOONE\n%ENV(NON VALID)\nEXISTS"
	os.Setenv("TEST", "424242")
	os.Setenv("locase", "My test value")
	os.Setenv("ONE", "TWO")
	os.Setenv("TWO", "ONE")
	os.Setenv("NON VALID", "WTF")
	os.Setenv("ITS_VALID", "EXISTS")
	result := SubstStringEnvVars(testText)
	os.Unsetenv("TEST")
	os.Unsetenv("locase")
	os.Unsetenv("ONE")
	os.Unsetenv("TWO")
	os.Unsetenv("NON VALID")
	os.Unsetenv("ITS_VALID")
	assertEqual := assert.Equal(t)
	assertEqual(string(result), string(expectedResult))
}
