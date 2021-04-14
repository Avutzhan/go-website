# go-website

BUG MUGS

ENV BUGS

1. GOROOT is root directory for common go files (for go itself)
   for example in go version go1.16.3 linux/amd64
   GOROOT="/usr/local/go"

2. GOPATH is my projects directory
   GOPATH="/home/avutzhan/go-workspace"

3. ok but there is another bugs
   in your GOPATH must be those directories: bin pkg src 
   in src located your new projects 

4. use the latest stable version go version go1.16.3 linux/amd64
   because you will get some error

5. in new version of golang make go mod init for tracking packages versions
   this is functionality similar to pakcage.json composer.json


CODE BUGS
1. handle all errors because golnag cant do it itself like in laravel
   php node.js
   
2. in templates in html siles when you work with "html/template" dont use comment
  because comment will crash the program