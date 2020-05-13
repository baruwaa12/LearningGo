// Package paasio provides functionality around counting operations
package paasio

import (
	"io"
	"sync"
)

type WriteCount struct {
	writer io.Writer

	mutex        *sync.RWMutex
	calls        int64
	bytesWritten int64
}

func NewWriteCounter(w io.Writer) WriteCount {
	wc := WriteCount{
		writer:       w,
		bytesWritten: 0,
		mutex:        &sync.RWMutex{},
	}
	return &wc
}

func (wc *WriteCount) Write(p []byte) (n int, err error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	wc.calls++

	n, err = wc.writer.Write(p)
	if err != nil {
		return 0, err
	}

	wc.bytesWritten += int64(n)
	return n, nil
}

func (wc *WriteCount) WriteCount() (n int64, nops int) {
	wc.mutex.RLock()
	defer wc.mutex.RUnlock()
	return wc.bytesWritten, int(wc.calls)
}

type ReadCount struct {
	reader io.Reader

	mutex     *sync.RWMutex
	calls     int
	bytesRead int64
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &ReadCount{
		reader:    r,
		bytesRead: 0,
		calls:     0,
		mutex:     &sync.RWMutex{},
	}
}

func (rc *ReadCount) Read(p []byte) (n int, err error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	rc.calls++

	n, err = rc.reader.Read(p)
	if err != nil {
		return 0, err
	}

	rc.bytesRead += int64(n)

	return n, nil
}

func (rc *ReadCount) ReadCount() (n int64, nops int) {
	rc.mutex.RLock()
	defer rc.mutex.RUnlock()
	return rc.bytesRead, rc.calls
}

type ReadWriteCount struct {
	writer WriteCounter
	reader ReadCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &ReadWriteCount{
		writer: NewWriteCounter(rw),
		reader: NewReadCounter(rw),
	}
}

func (rw *ReadWriteCount) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

func (rw *ReadWriteCount) ReadCount() (n int64, nops int) {
	return rw.reader.ReadCount()
}

func (rw *ReadWriteCount) Write(p []byte) (n int, err error) {
	return rw.writer.Write(p)
}

func (rw *ReadWriteCount) WriteCount() (n int64, nops int) {
	return rw.writer.WriteCount()
}