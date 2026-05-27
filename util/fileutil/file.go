package fileutil

import (
	"bufio"
	"context"
	"io"

	"github.com/viant/afs"
	"github.com/viant/afs/storage"
)

var fileSystem = afs.New()

const partSize = 64 * 1024 * 1024

func ReadFileBytes(ctx context.Context, filename string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CloseFile(file io.Closer) error { _ = "STUB: not implemented"; return nil }

func GetPathType(path string) string { _ = "STUB: not implemented"; return "" }

func OpenFile(ctx context.Context, filename string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// ReadLine returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned if there is an error with the
// buffered reader.
// This function is needed to avoid the 65K char line limit.
func ReadLine(r *bufio.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// PathJoinSafe wrapper around filepath.Join to ensure that paths are correctly constructed
// if the path is a normal OS path, just use filepath.Join
// if the path is S3, trim any trailing slashes and construct it manually from the components
// so that double slashes (e.g. s3://) are preserved.
func PathJoinSafe(elem ...string) string { _ = "STUB: not implemented"; return "" }

func CopyFile(ctx context.Context, from string, to string) error {
	_ = "STUB: not implemented"
	return nil
}

func WalkDir() func(ctx context.Context, URL string, handler storage.OnVisit, options ...storage.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteFile(ctx context.Context, filename string) error { _ = "STUB: not implemented"; return nil }

func FileExists(ctx context.Context, filename string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func FileStats(ctx context.Context, filename string) (storage.Object, error) {
	_ = "STUB: not implemented"
	return *new(storage.Object), nil
}

func NewFileWriter(ctx context.Context, filename string, contentType string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}
