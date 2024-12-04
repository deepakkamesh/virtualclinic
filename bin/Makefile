
hello:
	echo "hello"

build: 
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0  go build ./main.go

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
