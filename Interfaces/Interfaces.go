package fmt

import (
	"bytes"
	"io"
	"os"
)

// Fprintf function
// func Fprintf(w io.Writer, format string, args ...interface{}) (int, error){
// 	// return (int, error)
// }

// Printf Function
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

// Sprintf function
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()

}
