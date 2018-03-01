// package util contains utility functions that should probably move into their own library someday. These ones are 4 Windows.
package util

import "os"

func Exists(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {

		return false
	}
	return true
}
