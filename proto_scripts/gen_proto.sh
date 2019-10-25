#!/bin/sh

##################################################
###   proto & api
##################################################
#modify the path to your self path.
#export PATH_AGENT=/go/gonet2/agent/src/client_handler
#export PATH_GAME=/go/gonet2/game/src/client_handler
export GOPATH=$(pwd)
export PATH_AGENT=./proto_code/agent/
export PATH_GAME=./proto_code/game/
export PATH_CLIENT=./proto_code/client/
export PATH_SIMULATE=./proto_code/simulate/

#go get github.com/urfave/cli

## api.txt
go run api.go --template "templates/agent/api.tmpl" --min 0 --max 1000 > $PATH_AGENT/api.go; go fmt $PATH_AGENT/api.go
go run api.go --min 1001 --max 32767 > $PATH_GAME/api.go; go fmt $PATH_GAME/api.go
go run api.go --template "templates/client/api.tmpl"  --min 0 --max 32767 > $PATH_CLIENT/NetApi.cs
go run api.go --template "templates/simulate/api.tmpl"  --min 0 --max 1000 > $PATH_SIMULATE/api.go; go fmt $PATH_SIMULATE/api.go


## proto.txt
go run proto.go --template "templates/agent/proto.tmpl" > $PATH_AGENT/proto.go; go fmt $PATH_AGENT/proto.go
go run proto.go > $PATH_GAME/proto.go; go fmt $PATH_GAME/proto.go
go run proto.go --template "templates/client/proto.tmpl" --binding "cs" > $PATH_CLIENT/NetProto.cs
go run proto.go --template "templates/simulate/proto.tmpl"> $PATH_SIMULATE/proto.go; go fmt $PATH_SIMULATE/proto.go

export PROJECTDIR=$HOME/code
cp $PATH_AGENT/*.go $PROJECTDIR/agent/src/app/client_handler/
cp $PATH_SIMULATE/*.go $PROJECTDIR/tools/simulate/src/simulate/
cp $PATH_GAME/proto.go $PROJECTDIR/game/src/app/client_proto/
cp $PATH_GAME/api.go $PROJECTDIR/game/src/app/client_handler/
cp $PATH_CLIENT/*.cs $PROJECTDIR/NetProto
cp $PROJECTDIR/tools/proto_scripts/*.txt $PROJECTDIR/NetProto

sleep 10
