build:
	CGO_ENABLED=0 GOOS=linux go build -a ./...
	docker build -t kristenfjacobs/polibot:1.0.0 .

push:
	docker push kristenfjacobs/polibot:1.0.0

run:
	docker run -d -p 8081:8081 --name polibot -t kristenfjacobs/polibot:1.0.0

test:
	while true; do curl localhost:8081; sleep 1; done

clean:
	docker rm -f polibot
	rm -rf polibot
