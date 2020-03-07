package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Configuration allow to setup the server configurations
type Configuration struct {
	Hostname        string
	Port            string
	DefaultResource string
	RootPath        string
}

var cache ResourceCache
var fallback string

func handler(response http.ResponseWriter, request *http.Request) {
	log.Printf("Requested %s\n", request.URL.Path)
	content, err := cache.GetResource(request.URL.Path)
	if err != nil {
		log.Printf("Resource %s not found, fallback to %s\n", request.URL.Path, fallback)
		http.Redirect(response, request, fallback, 307)
		return
	}
	response.Header().Add("Content-Encoding", content.CompressionType)
	response.Header().Add("Content-Type", content.Type)
	response.Header().Add("Content-Length", strconv.Itoa(content.Size))
	response.Write(content.Content)
}

func readyHandler(response http.ResponseWriter, request *http.Request) {
	if cache.Size() > 0 {
		response.WriteHeader(200)
	} else {
		response.WriteHeader(500)
	}
}

// StartServer starts the HTTP Server and loads the resources
func StartServer(config Configuration) error {
	log.Printf("Ozone scanning resourses at %s\n", config.RootPath)
	files, errFiles := SearchFiles(config.RootPath)
	if errFiles != nil {
		return errFiles
	}
	log.Printf("Ozone found %d resources\n", len(files))
	cache = NewCache(files)
	for name, resource := range cache.resources {
		log.Printf("Loaded %s - %d bytes => %d bytes\n", name, resource.InitialSize, resource.Size)
	}
	log.Printf("Resources ready to serve\n")
	fallback = config.DefaultResource

	if !cache.Exists(fallback) {
		return fmt.Errorf("Default resouce %s not found in resources", fallback)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/service/ready", readyHandler)

	log.Printf("Ozone starting at http://%s:%s\n", config.Hostname, config.Port)
	log.Printf("Ozone rediness service at http://%s:%s/service/ready\n", config.Hostname, config.Port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Hostname, config.Port), nil)
	if err != nil {
		return fmt.Errorf("Ozone fail to start: %v", err)
	}
	return nil
}
