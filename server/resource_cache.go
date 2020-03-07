package server

import (
	"fmt"
	"log"
)

// ResourceCache hold all the resources to be served
type ResourceCache struct {
	resources map[string]InMemoryResource
}

// NewCache creates a Cache object
func NewCache(files []ResourceFile) ResourceCache {
	rsrcs := make(map[string]InMemoryResource, len(files))
	for _, file := range files {
		content, err := GetResourceContent(file.AbsPath)
		if err != nil {
			log.Printf("Fail to get file content %s: %v", file.AbsPath, err)
		}

		rsrcs[file.RelPath], _ = NewInMemoryResource(file.RelPath, content)
	}
	return ResourceCache{
		resources: rsrcs,
	}
}

// GetResource return the in memory resource content
func (cache ResourceCache) GetResource(name string) (InMemoryResource, error) {
	if resource, found := cache.resources[name]; found {
		return resource, nil
	}
	return InMemoryResource{}, fmt.Errorf("Resource %s not found", name)
}

// Size returns the number of resources in memory
func (cache ResourceCache) Size() int {
	return len(cache.resources)
}

// Exists return true if the resource exists on cache and false otherwise
func (cache ResourceCache) Exists(name string) bool {
	_, found := cache.resources[name]
	return found
}
