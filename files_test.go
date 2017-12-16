package octoprint

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUploadFileRequest_Do(t *testing.T) {
	cli := NewClient("http://localhost:5000", "")

	r := &UploadFileRequest{Location: Local}
	err := r.AddFile("foo.gcode", bytes.NewBufferString("foo"))
	assert.NoError(t, err)

	state, err := r.Do(cli)
	assert.NoError(t, err)
	assert.Equal(t, "foo.gcode", state.File.Local.Name)
}

func TestUploadFileRequest_DoWithFolder(t *testing.T) {
	cli := NewClient("http://localhost:5000", "")

	r := &UploadFileRequest{Location: Local}
	err := r.AddFolder("qux")
	assert.NoError(t, err)

	state, err := r.Do(cli)
	assert.NoError(t, err)
	assert.Equal(t, true, state.Done)

}

func TestFilesRequest_Do(t *testing.T) {
	TestUploadFileRequest_Do(t)

	cli := NewClient("http://localhost:5000", "")

	files, err := (&FilesRequest{}).Do(cli)
	assert.NoError(t, err)

	assert.True(t, len(files.Files) >= 1)

	r := &FileRequest{Location: Local, Filename: "foo.gcode"}
	file, err := r.Do(cli)
	assert.NoError(t, err)

	assert.Equal(t, "foo.gcode", file.Name)
}

func TestSelectFileRequest_Do(t *testing.T) {
	cli := NewClient("http://localhost:5000", "")

	ur := &UploadFileRequest{Location: Local}
	err := ur.AddFile("foo2.gcode", bytes.NewBufferString("foo"))
	assert.NoError(t, err)
	_, err = ur.Do(cli)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 300)

	r := &SelectFileRequest{Location: Local, Path: "foo2.gcode"}
	err = r.Do(cli)
	assert.NoError(t, err)
}

func TestFilesRequest_DoWithLocation(t *testing.T) {
	cli := NewClient("http://localhost:5000", "")

	files, err := (&FilesRequest{Location: SDCard}).Do(cli)
	assert.NoError(t, err)

	assert.Len(t, files.Files, 0)
}
