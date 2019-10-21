package mocks

//REPOSITORY
//go:generate mockgen -destination=./mock-data-repository.go -package=mocks github.com/Oleg-Skalozub/testtask/src/domain/repository DataRepository

//SERVICES
//go:generate mockgen -destination=./mock-fetch.go -package=mocks github.com/Oleg-Skalozub/testtask/src/domain/services Fetcher

//INFRASTRUCTURE
//go:generate mockgen -destination=./mock-pgsql.go -package=mocks github.com/Oleg-Skalozub/testtask/src/infrastructure/db SQLBDer
//go:generate mockgen -destination=./mock-client.go -package=mocks github.com/Oleg-Skalozub/testtask/src/infrastructure/client ClientInterface
//go:generate mockgen -destination=./mock-logger.go -package=mocks github.com/Oleg-Skalozub/testtask/src/infrastructure/logger Logger
