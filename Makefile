
build:
	go build -o bin/api-center ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	rm -rf bin

test:
	go test -v ./...

migrate:
	go run ./cmd/main.go migrate

migrate-up:
	go run ./cmd/main.go migrate -m "up"

migrate-down:
	go run ./cmd/main.go migrate -m "down"

migrate-redo:
	go run ./cmd/main.go migrate -m "redo"

lint:
	golangci-lint run --verbose ./...

docker-build:
	docker build -t api-center .

docker-run:
	docker run -p 9200:9200 -d api-center

docker-stop:
	docker stop $(docker ps -a -q)

docker-rm:
	docker rm $(docker ps -a -q)

docker-clean:
	docker system prune -af

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down
