// Package twofer create a sentence of the form "One for X, one for me.
package twofer

import "fmt"

// ShareWith return a string with the message:
//```
//One for X, one for me.
//```
//Where X is the given name.
//However, if the name is missing, return the string:
//```
//One for you, one for me.
//```
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
