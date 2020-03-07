package server

import (
	"bytes"
	"compress/flate"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestNewInMemoryResource(t *testing.T) {
	type args struct {
		name            string
		content         []byte
		compressionType Compression
	}
	tests := []struct {
		name    string
		args    args
		want    InMemoryResource
		wantErr bool
	}{

		{
			name: "Create In Memory Resource",
			args: args{
				name:            "/index.html",
				content:         []byte("File Content to Test aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
				compressionType: CompressionDeflate,
			},
			want: InMemoryResource{
				Name:            "/index.html",
				CompressionType: CompressionDeflate,
				InitialSize:     77,
				Size:            30,
				Type:            "text/html; charset=utf-8",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInMemoryResource(tt.args.name, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInMemoryResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want.String() {
				t.Errorf("NewInMemoryResource() = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestInMemoryResource_String(t *testing.T) {
	type fields struct {
		Name            string
		Type            ResourceType
		CompressionType Compression
		Size            int
		InitialSize     int
		Content         []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "",
			fields: fields{
				Name:            "/file.txt",
				CompressionType: CompressionDeflate,
				Content:         []byte("Text to test compression"),
				InitialSize:     24,
				Size:            12,
				Type:            "text/plain",
			},
			want: "{ name: '/file.txt', size: 12, initial-size: 24, contentType: 'text/plain', compressionType: 'deflate' }",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := &InMemoryResource{
				Name:            tt.fields.Name,
				Type:            tt.fields.Type,
				CompressionType: tt.fields.CompressionType,
				Size:            tt.fields.Size,
				InitialSize:     tt.fields.InitialSize,
				Content:         tt.fields.Content,
			}
			if got := resource.String(); got != tt.want {
				t.Errorf("InMemoryResource.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressResource(t *testing.T) {
	type test struct {
		name    string
		args    []byte
		want    string
		wantErr bool
	}
	tests := []test{{
		name:    "Compress and Uncompress match",
		args:    []byte("Text to test compression"),
		want:    "Text to test compression",
		wantErr: false,
	}, {
		name:    "Compress and Uncompress nil content",
		args:    nil,
		want:    "",
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compressResource(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("compressResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotD, errD := decompress(got)
			if (errD != nil) != tt.wantErr {
				t.Errorf("compressResource() decompress error = %v, wantErr %v", errD, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.want) {
				t.Errorf("compressResource() = '%v', want '%v'", gotD, tt.want)
			}
		})
	}
}

func decompress(content []byte) (string, error) {
	fr := flate.NewReader(bytes.NewBuffer(content))
	defer fr.Close()
	c, err := ioutil.ReadAll(fr)
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", err
	}
	return string(c), nil
}
