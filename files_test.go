package octoprint

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadFileRequest_Do(t *testing.T) {
	cli := NewClient("http://localhost:5000", "")

	r := &UploadFileRequest{Location: "local"}
	err := r.AddFile("foo.gcode", bytes.NewBufferString("foo"))
	assert.NoError(t, err)

	state, err := r.Do(cli)
	assert.NoError(t, err)
	assert.Equal(t, "foo.gcode", state.File.Local.Name)
}

func TestUploadFilesRequest_Do(t *testing.T) {
	TestUploadFileRequest_Do(t)

	cli := NewClient("http://localhost:5000", "")

	r := &FilesRequest{}
	files, err := r.Do(cli)
	assert.NoError(t, err)

	assert.Len(t, files.Files, 1)
	assert.Equal(t, "foo.gcode", files.Files[0].Name)
}
