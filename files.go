package octoprint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

const (
	URIFiles = "/api/files"
)

// FilesRequest retrieve information regarding all files currently available and
// regarding the disk space still available locally in the system.
type FilesRequest struct {
	// Recursive if set to true, return all files and folders recursively.
	// Otherwise only return items on same level.
	Recursive bool
}

// Do sends an API request and returns the API response.
func (cmd *FilesRequest) Do(c *Client) (*FilesResponse, error) {
	uri := fmt.Sprintf("%s?recursive=%t", URIFiles, cmd.Recursive)

	b, err := c.doJSONRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	r := &FilesResponse{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, err
}

// UploadFileRequest uploads a file to the selected location or create a new
// empty folder on it.
type UploadFileRequest struct {
	// Location is the target location to which to upload the file. Currently
	// only `local` and `sdcard` are supported here, with local referring to
	// OctoPrint’s `uploads` folder and `sdcard` referring to the printer’s
	// SD card. If an upload targets the SD card, it will also be stored
	// locally first.
	Location string

	b *bytes.Buffer
	w *multipart.Writer
}

// AddFile adds a new file to be uploaded from a given reader.
func (req *UploadFileRequest) AddFile(filename string, r io.Reader) error {
	if req.w == nil {
		req.b = bytes.NewBuffer(nil)
		req.w = multipart.NewWriter(req.b)
	}

	w, err := req.w.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	return err

}

// Do sends an API request and returns the API response.
func (req *UploadFileRequest) Do(c *Client) (*UploadFileResponse, error) {
	req.w.Close()

	uri := fmt.Sprintf("%s/%s", URIFiles, req.Location)
	b, err := c.doRequest("POST", uri, req.w.FormDataContentType(), req.b)
	if err != nil {
		return nil, err
	}

	r := &UploadFileResponse{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, err
}
