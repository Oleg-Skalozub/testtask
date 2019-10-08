LOCALCONFIGFILE=config.json

.PHONY: run
run:
	rm -f *.log
	go run -race src/main.go -config $(LOCALCONFIGFILE)

.PHONY: clean
clean:
	@echo ===== Stop and remove containers
	@cd ../dev ; docker-compose stop ; docker-compose rm -f

.PHONY: start
start:
	@echo ===== Start docker compose
	@docker-compose -f docker-compose.yml up -d

.PHONY: migrate
migrate:
	rm -f *.log
	go run -race src/infrastructure/db/persist/main.go -config $(LOCALCONFIGFILE)

.PHONY: mocks
mocks:
	go get -u github.com/golang/mock/mockgen
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: code-quality
code-quality:
	golangci-lint --exclude-use-default=false --out-format checkstyle run ./... > static-analysis.xml

.PHONY: code-quality-print
code-quality-print:
	golangci-lint --exclude-use-default=false  --out-format tab run ./...
