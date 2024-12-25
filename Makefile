
hello:
	echo "hello"

build: 
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0  go build ./main.go

build-sysagent-cli: 
	go build -o ./bin/sysagent-cli/cli ./bin/sysagent-cli/

build-sysagent:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0  go build -o ./bin/sysagent/sysagent ./bin/sysagent/

push-sysagent:
	rsync -avz -e "ssh -o StrictHostKeyChecking=no" --progress ./bin/sysagent/sysagent drguru@192.168.68.119:~/
	rsync -avz -e "ssh -o StrictHostKeyChecking=no" --progress virtualclinic.conf.toml drguru@192.168.68.119:~/

push:
	rsync -avz -e "ssh -o StrictHostKeyChecking=no" --progress main drguru@192.168.68.119:~/
	rm main

all: build push

bprint:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build ../prototypes/printer.go

pprint:
	rsync -avz -e "ssh -o StrictHostKeyChecking=no" --progress printer drguru@192.168.68.119:~/
	rm printer

print: bprint pprint
