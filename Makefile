test: 
	go test ./...

docker-run-interactive: build
	docker run -p 8080:8080 -it silo-agent 

docker-run-daemon: build
	docker run -d -p 8080:8080 -it silo-agent 

build:
	docker build -t silo-agent .
