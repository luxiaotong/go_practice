docker run -itd --name "test-informix" \
-v $PWD/../../:/app -w /app \
-v $GOPATH:/go \
golang:1.18.1

IBM_DB_HOME=/app go run /go/pkg/mod/github.com/ibmdb/go_ibm_db\@v0.4.2/installer/setup.go

export IBM_DB_HOME=/go/pkg/mod/github.com/ibmdb/clidriver
export CGO_CFLAGS=-I$IBM_DB_HOME/include
export CGO_LDFLAGS=-L$IBM_DB_HOME/lib 
export LD_LIBRARY_PATH=/go/pkg/mod/github.com/ibmdb/clidriver/lib

IBM_DB_HOME=/go/pkg/mod/github.com/ibmdb/clidriver \
CGO_CFLAGS=-I$IBM_DB_HOME/include \
CGO_LDFLAGS=-L$IBM_DB_HOME/lib \
LD_LIBRARY_PATH=/go/pkg/mod/github.com/ibmdb/clidriver/lib \
go run main.go

apt-get purge -f libxml2-dev
apt-get clean
apt-get update
apt-get install libxml2