APP_BIN = build/app

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/app/main.go

clean:
	rm -rf ./app/build || true

swagger:
	swag init -g ./cmd/app/main.go -o ./docs

migrate:
	$(APP_BIN) 	migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up