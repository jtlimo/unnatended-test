# Yoke Cards

API to provide the functionalities to create a playing card games with these specifics endpoints

- `GET /api/v1/deck/:deckId`
- `PUT /api/v1/deck` or `PUT /api/v1/deck?cards=AS,KD,AC,2C,KH`
- `PUT /api/v1/deck/:deckId/:count`

## Dependencies
* Go 1.16.3

## Build project
```console
$ make build
```

## Run tests
```console
$ make test
```

## Start the server
```console
$ make start
```

>You will see in your terminal something like that:

![Alt text](./assets/start-server.png?raw=true "Initializing the server")

>And now open the browser on `localhost:3000`
