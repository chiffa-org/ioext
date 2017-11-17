package ioext

import (
	"io"
)

type ReaderFunc func(p []byte) (n int, err error)
type WriterFunc func(p []byte) (n int, err error)
type CloserFunc func() error
type SeekerFunc func(offset int64, whence int) (int64, error)

func (r ReaderFunc) Read(p []byte) (n int, err error)             { return r(p) }
func (w WriterFunc) Write(p []byte) (n int, err error)            { return w(p) }
func (c CloserFunc) Close() error                                 { return c() }
func (s SeekerFunc) Seek(offset int64, whence int) (int64, error) { return s(offset, whence) }

func CompositeReadWriter(r io.Reader, w io.Writer) io.ReadWriter {
	return struct {
		io.Reader
		io.Writer
	}{Reader: r, Writer: w}
}

func CompositeReadCloser(r io.Reader, c io.Closer) io.ReadCloser {
	return struct {
		io.Reader
		io.Closer
	}{Reader: r, Closer: c}
}

func CompositeWriteCloser(w io.Writer, c io.Closer) io.WriteCloser {
	return struct {
		io.Writer
		io.Closer
	}{Writer: w, Closer: c}
}

func CompositeReadWriteCloser(r io.Reader, w io.Writer, c io.Closer) io.ReadWriteCloser {
	return struct {
		io.Reader
		io.Writer
		io.Closer
	}{Reader: r, Writer: w, Closer: c}
}

func CompositeReadSeeker(r io.Reader, s io.Seeker) io.ReadSeeker {
	return struct {
		io.Reader
		io.Seeker
	}{Reader: r, Seeker: s}
}

func CompositeWriteSeeker(w io.Writer, s io.Seeker) io.WriteSeeker {
	return struct {
		io.Writer
		io.Seeker
	}{Writer: w, Seeker: s}
}

func CompositeReadWriteSeeker(r io.Reader, w io.Writer, s io.Seeker) io.ReadWriteSeeker {
	return struct {
		io.Reader
		io.Writer
		io.Seeker
	}{Reader: r, Writer: w, Seeker: s}
}
