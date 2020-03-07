package server

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	var cache = make(map[string]InMemoryResource)
	cache["/index.html"], _ = NewInMemoryResource("/index.html", nil)
	type args struct {
		files []ResourceFile
	}
	tests := []struct {
		name string
		args args
		want ResourceCache
	}{

		{
			name: "Cache Empty",
			args: args{
				files: []ResourceFile{},
			},
			want: ResourceCache{
				resources: make(map[string]InMemoryResource),
			},
		},
		{
			name: "Cache with Resources",
			args: args{
				files: []ResourceFile{{AbsPath: "/app/index.html", RelPath: "/index.html"}},
			},
			want: ResourceCache{
				resources: cache,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceCache_GetResource(t *testing.T) {
	var cache = make(map[string]InMemoryResource)
	mr, _ := NewInMemoryResource("/index.html", nil)
	cache["/index.html"] = mr
	type fields struct {
		resources map[string]InMemoryResource
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    InMemoryResource
		wantErr bool
	}{

		{
			name: "Get Existing Resource",
			args: args{
				name: "/index.html",
			},
			fields: fields{
				resources: cache,
			},
			want:    mr,
			wantErr: false,
		},
		{
			name: "Get Non Existing Resource",
			args: args{
				name: "/fake.html",
			},
			fields: fields{
				resources: cache,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := ResourceCache{
				resources: tt.fields.resources,
			}
			got, err := cache.GetResource(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceCache.GetResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceCache.GetResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceCache_Size(t *testing.T) {
	var cache = make(map[string]InMemoryResource)
	cache["/index.html"], _ = NewInMemoryResource("/index.html", nil)
	type fields struct {
		resources map[string]InMemoryResource
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{

		{
			name: "Size of Cache is 1",
			fields: fields{
				resources: cache,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := ResourceCache{
				resources: tt.fields.resources,
			}
			if got := cache.Size(); got != tt.want {
				t.Errorf("ResourceCache.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceCache_Exists(t *testing.T) {
	var cache = make(map[string]InMemoryResource)
	cache["/index.html"], _ = NewInMemoryResource("/index.html", nil)
	type fields struct {
		resources map[string]InMemoryResource
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{

		{
			name: "Resource Found",
			args: args{
				name: "/index.html",
			},
			fields: fields{
				resources: cache,
			},
			want: true,
		},
		{
			name: "Resource Not Found",
			args: args{
				name: "/fake.html",
			},
			fields: fields{
				resources: cache,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := ResourceCache{
				resources: tt.fields.resources,
			}
			if got := cache.Exists(tt.args.name); got != tt.want {
				t.Errorf("ResourceCache.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
