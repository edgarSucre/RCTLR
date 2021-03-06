//Package utils provides common funcitonality for test
package utils

import (
	"fmt"
	"strings"
	"testing"
)

const (
	infoColor    = "\033[1;34m\t%s\033[0m\n\n"
	warningColor = "\033[1;33m\t%s\033[0m\n"
	errorColor   = "\033[1;31m\t%s\033[0m\n"
)

//AssertEquals log errors if values are not equal
func AssertEquals(expected, actual interface{}, message string, t *testing.T) {
	if expected != actual {
		Err(message)
		Warn(fmt.Sprint("Expected:", expected))
		Warn(fmt.Sprint("Actual:", actual, "\n"))
		t.Fail()
	}
}

//AssertContains log errores si el contenido no incluye el token
func AssertContains(content, token, message string, t *testing.T) {
	if !strings.Contains(content, token) {
		Err(message)
		Warn(fmt.Sprint("Mensaje: ", content))
		Warn(fmt.Sprint("No contiene: ", token))
		t.Fail()
	}
}

//AssertNotContains log errores si el contenido incluye el token
func AssertNotContains(content, token, message string, t *testing.T) {
	if strings.Contains(content, token) {
		Err(message)
		Warn(fmt.Sprint("Mensaje: ", content))
		Warn(fmt.Sprint("Contiene: ", token))
		t.Fail()
	}
}

//AssertTrue log errores si token es false
func AssertTrue(token bool, message string, t *testing.T) {
	if !token {
		Err(message)
		t.Fail()
	}
}

//AssertFalse log errores si token es true
func AssertFalse(token bool, message string, t *testing.T) {
	AssertTrue(!token, message, t)
}

//Err log error message
func Err(m string) {
	fmt.Printf(errorColor, m)
}

//Warn log warning message
func Warn(m string) {
	fmt.Printf(warningColor, m)
}

//Info log info message
func Info(m string) {
	fmt.Printf(infoColor, m)
}
