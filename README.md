# MagaCoin API

This application is responsible for provide wallet informations to MagaCoin clients and integrate to MagaCoin Payments API.

## Table of contents

- [MagaCoin API](#MagaCoin API)
- [Table of contents](#Table-of-contents)
- [Installing](#Installing)
- [Environment Variables](#Environment-Variables)
- [Building](#Building)
- [Running](#Running)
- [Deploying](#Deploying)

## Installing

Assuming that you have already cloned the project and the
[Go](https://golang.org/doc/install) is installed, the first
step is install the [dep](https://github.com/golang/dep) and
ensure that all dependencies are vendored in the project:

```sh
$ dep ensure
```

## Environment Variables

```
Variable                      | Type    | Description                                              
----------------------------- | ------- | ---------------------------------------------------------
DATABASE_HOST                 | string  | Database host
DATABASE_PORT                 | string  | Database port
DATABASE_NAME                 | string  | Database name
DATABASE_PASSWORD             | string  | Database Password
DATABASE_USERNAME             | string  | Database username
PRIVATE_KEY                   | string  | Hedere private key of current user account
TARGET_ACCOUNT                | integer | Hedera account ID which will receive the transfer ammount
DATABASE_MAX_IDLE             | integer | Database max idle value
DATABASE_IDLE_TIMEOUT         | integer | Database idle timeout value
MAGACOIN_PAYMENT_URL          | integer | MagaCoin Payment URL
MAGACOIN_PAYMENT_TIMEOUT      | integer | MagaCoin Payment Timeout
DATABASE_MAX_CON              | integer | Database max. connections
```

## Building

Build project:

```sh
$ make build
```

## Running

Start the app consumer:

```sh
$ make start
```

## Deploying

Deploy application to Teresa:

```sh
$ bash ./scripts/deploy.sh -f <env_file_name>
```
