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

- authentication:
  - create and parse tokens (`jwt`)
  - Implement `webauthn` login and register
  - add route `/auth/delete` to delete an account
- cli:
  - fix foreign keys order bug (order tables when creating/filling (accounts, then session and items))
- routes:
  - set proper request handlers type (post, get, put, delete...)
  - switch to `fasthttp`
    - https://github.com/valyala/fasthttp
    - https://github.com/julienschmidt/httprouter
- testing:
  - write unit tests, functionnal tests
  - set tests steps
    - init: create an user, add a note
    - end: delete note, delete user
- ci-cd:
  - on merge: trigger build on circle-ci
