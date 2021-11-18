# Danale Golang Makefile

# variable
bizLine=demo
binaryName=dartman-demo-client
packageName=go.danale.net/be/be/biz/demo/dartman-demo-client
versionPath=go.danale.net/be/be/go/version
version=v0.1.0
outputPath=_output

all: build

init:
	go mod init ${packageName};go mod tidy;
 
build:
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	go build -ldflags " -X '${versionPath}.Version=${version}' \
						-X '${versionPath}.BuildTime=$$buildTime' \
						-X '${versionPath}.GoVersion=`go version`' \
						-X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/${binaryName}; \

pandora: build
	cd ${outputPath}; \
	tar zcvf ${binaryName}.tar.gz ${binaryName}; \
	name=${bizLine}@${binaryName}:${version}; \
	pandora -pd $$name ${binaryName}.tar.gz; \

clean:
	rm -rf _output
	rm -rf logs

.PHONY: all build pandora clean
