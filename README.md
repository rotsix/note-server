# server

## Install

- install and trust certs: `mkcert localhost; mkcert -install`
- retrieve dependencies: `go get ./...`
- set up databases: `docker-compose up -d db; source alias.sh; db_local`

## Running

You can run it using docker-compose, (`docker-compose up`).
Don't forget to pass `--build` option when modifying code.

To mix between local server and dockerized db, run:

- `docker-compose up -d db`
- `source alias.sh; server_local`

Mock users are `test/test` and `testN/testN` with N in 1-2-3.

## Database, and tables

db note:
- sessions
  - id
  - user_id
  - token
- accounts
  - id
  - username
  - password


## Todo

### Authentication

Creating and parsing tokens
Note: JWT

Managing other authentication protocols
Note: fido2, webauthn, oauth2, sqrl

### cli

Fix foreign keys order bug
Order tables when creating/filling (accounts, then session and items)
