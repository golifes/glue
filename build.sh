#!/bin/bash

CGO_ENABLED=0 gox -osarch="linux/amd64" -ldflags "-s -w " -gcflags -m -output dist/glue
# docker run --rm -it -v $GOPATH:/go golang:latest bash -c 'cd $GOPATH/src/glue && CGO_ENABLED=0 go build -ldflags "-s -w " -gcflags -m -o dist/glue'
# cp -rf conf dist
upx -f -9 dist/glue && docker build -t api .

# delivery ---

# docker run  -v $GOPATH/src/glue/static:/root/app/static -p 8080:8080  api
# docker exec -it d460ba92ea0a  /bin/bash
