package server

import (
	"bytes"
	"compress/flate"
	"fmt"
	"mime"
	"path/filepath"
)

// ResourceType specifies the content type of the resource
type ResourceType = string

// CompressionType specifies the content compression type
type CompressionType = string

// InMemoryResource hold the resouce that can be served by the server
type InMemoryResource struct {
	Name            string
	Type            ResourceType
	CompressionType CompressionType
	Size            int
	InitialSize     int
	Content         []byte
}

// NewInMemoryResource method will initialize a new InMemoryResource object and return it
func NewInMemoryResource(name string, content []byte, compressionType CompressionType) (InMemoryResource, error) {
	compressedContent, err := compressResource(content)
	if err != nil {
		return InMemoryResource{}, err
	}
	return InMemoryResource{
		Name:            name,
		InitialSize:     len(content),
		Size:            len(compressedContent),
		Content:         compressedContent,
		Type:            mime.TypeByExtension(filepath.Ext(name)),
		CompressionType: compressionType,
	}, nil
}

// String method will genereate a string representation of the InMemoryResource object
func (resource *InMemoryResource) String() string {
	return fmt.Sprintf("{ name: '%s', size: %d, initial-size: %d, contentType: '%s', compressionType: '%s' }",
		resource.Name,
		resource.Size,
		resource.InitialSize,
		resource.Type,
		resource.CompressionType,
	)
}

// compressResource receives a []byte with the resource content and returns a compressed version
func compressResource(content []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	flateWriter, _ := flate.NewWriter(buf, flate.BestCompression)
	_, err := flateWriter.Write(content)
	flateWriter.Flush()
	if err != nil {
		return nil, err
	}
	defer flateWriter.Close()
	return buf.Bytes(), nil
}
