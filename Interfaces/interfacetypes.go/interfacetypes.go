package main

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(string)
}

// Closer interface
type Closer interface {
	Close() error
}

// ReadWriter interface
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser interface
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}


// Alternate example of interface declaration
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}