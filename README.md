# Withdraw Deposit Microservice
For api url, api-key use main.yml as configuration file

## For first installation
run make install

## For data migration 
run make migrate-up  or run make migrate-down

## Start service
run make local

## Run unit test
rum make test


## Folder usage

### /config/app
Configuration files for development and production
 
### /config/i18n
Language translator

### /docs
Swagger api documentation, auto generated

### /entity
Define structs

### /init
Run inits process

### /internal
internal usage, not for public

### /logs
auto generated

### /module/handler/http
set endpoint

### /module/mocks
auto generated

### /module/store
query to database

### /module/usecase
business logic

### iAuthRepository.go 
contract for repository

### iAuthUsecase.go
contract for usecase

### /pkg
addtional package

### /schema
schema for request parameters and response

### /scripts/migrations
database migration script

### /vendor
3rd party apps
