# server

## Install

- create certs with mkcert
- dbcli all

## Running

You can decide to run it just using docker-compose, then run `docker-compose up`.
Don't forget to pass `--build` option when modifying code.

To mix between local server and dockerized db, run:

- `docker-compose up db`
- `source alias.sh; server_local`

Mock users are `test/test` and `testN/testN` with N in 1-2-3.

## Databases

### users

Stores users' informations (account relative)

- sessions
  - id
  - user_id
  - token
- informations: users information
  - id
  - username
  - password
  - firstname
  - lastname
  - mail
  - birthdate
  - address
  - zip_code
  - city
- configurations: client configurations
  - TODO: could be good to store this
  - eg. light/dark display, notifications frequency..
