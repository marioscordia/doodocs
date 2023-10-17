start:
	go run .
build:
	docker build -t file-handler .    
run:
	docker run -d -p 8080:8080 --name file-handler file-handler