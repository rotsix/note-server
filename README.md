# server

`server` is a note taking Rest API.

This project is made up of a Go (golang) application interacting with a Postgres database.

## Installation

Before installing and running this application, set up configuration file.
A sample configuration is located at `./configs/sample_config.json`.

> Note: in order to initialize database, set its name to `localhost` in configuration file.

Type following commands to prepare running environment.

- install and trust certs: `mkcert localhost; mkcert -install`
- get dependencies: `go mod download`
- set up databases: `docker-compose up -d db; source alias.sh; db_local`

## Usage

Application can be launched using `docker-compose up`.
Don't forget to pass `--build` option when modifying code.

If running all in docker containers, set the database name to docker-compose service (default `db`) name in configuration file.
Or if the app is running using plain Go commands, set database name to `localhost`.

Or mix between local server and dockerized database, so run:

- `docker-compose up -d db`
- `source alias.sh; server_local`

> Don't forget to set the `KEY_LOCATION`, `CERT_LOCATION` and `CONF_LOCATION` environment variables.

Mock users are `test/test` and `testN/testN` with N in 1-2-3.
The `test` user already have some notes associated.

## Database and tables

For more details, see [config](./configs/sample_config.json).

db note:
- session
  - id
  - user_id
  - token
  - expiration
  - last_seen
- user
  - id
  - username
  - password
- notes
  - id
  - user_id
  - title
  - description
  - creation_date
  - modification_date


## TODO

- authentication:
  - create and parse tokens (`jwt`)
  - Implement `webauthn` login and register
  - add route `/auth/delete` to delete an account
- cli:
  - fix foreign keys order bug (order tables when creating/filling (accounts, then session and items))
- routes:
  - set proper request handlers type (`post`, `get`, `put`, `delete`...)
  - set full rest end-points (id in url, and co)
  - switch to `fasthttp`
    - https://github.com/valyala/fasthttp
    - https://github.com/julienschmidt/httprouter
- testing:
  - write unit tests, functionnal tests [0]
  - steps (simplier: create a new db, and run `db all`)
    - signin, login
    - creating account, creating note
    - get all notes, get note, modify note
    - delete note, delete account
  - set-up testing using `postman` or `advanced-rest-client`
- ci-cd:
  - on commit on master: trigger build on `circle-ci`

[0] : https://stackoverflow.com/questions/31201858/how-to-run-golang-tests-sequentially
