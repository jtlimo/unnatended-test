# Unnatended test

API to provide the functionalities to create a playing card games with these specifics endpoints

- `GET /deck/:deckId`
- `POST /deck` or `POST /deck?cards=AS,KD,AC,2C,KH`
- `POST /deck/:deckId/:count`

## Dependencies
* Go 1.17.9

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

## How to create a new deck

> Open insomnia / postman or simply make a request with Curl to this endpoint:
- `POST /deck` or `POST /deck?cards=AS,KD,AC,2C,KH`

> Example using insomnia

![example of creating a deck](assets/create-deck.png "a title")

> Example using Curl

```console
curl --request POST \
  --url 'http://localhost:8080/deck?card=AS' \
  --header 'content-type: application/json'```
```

>And now open the browser on `localhost:8080`
