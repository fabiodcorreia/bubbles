# Ozone HTTP Server [![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/fabiodcorreia/ozone/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/fabiodcorreia/ozone)](https://goreportcard.com/report/github.com/fabiodcorreia/ozone)

Ozone is a minimalist Cloud Native HTTP Server to serve single page applications with no configuration and high performance.

## Motivation

There are many of great high-performance HTTP Servers, such as NGINX, Apache Httpd and others, so why another one?
Ozone was designed to be small, fast like the others but also to work on cloud containerized environments like Kubernetes.

## How it Works

When Ozone starts it loads all the resources of a specified directory in memory and compress them, after that it's ready to start serving these resources by http. It only loads and compresses the resources one time so no IO to the filesystem or CPU usage to compress the resouces on each call, also since the resouces are compresses the memory footprint its also lower.

1. Scan the directory and sub directories and list all the resources
2. For each resource load the content in memory and compress it
3. Each resource is stored in a hash table where they key is the resource path/name
4. The resource tree reflects the filesystem tree /index.html -> directory/index.html,  directory/static/js/main.js -> /static/js/main.js
5. By default all 404 will redirect to /index.html that also includes the root path /
6. To check if the server is ready we can call the endpoint /api/server/status, that returns 200 code if yes
7. When a resouce is requested the server fech the resource matching from the hashmap and returned it
8. If the default-resouce doesn't exists it fails

## Installation

## Usage
```
USAGE:
   ozone [global options]

GLOBAL OPTIONS:
   --hostname value, -n value  Set the hostname (default: "127.0.0.1")
   --port value, -p value      Set the listening port (default: "8080")
   --default value, -d value   Set the default resource to use on 404 (default: "/index.html")
   --root value, -r value      Set root directory to serve (default: ".")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```


