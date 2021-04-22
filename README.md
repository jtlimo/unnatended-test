<!-- 
Some problems that I could not resolve in time and improvements I saw in the code

- I would like to make some refactoring to cleanup the method NewDeck, I see some duplication, and I could extract to a method to be more clear;

- About project architecture I would like to separate better the code in main.go to another packages like handlers and so on.
 -->

# Unnantended Test

API to provide the functionalities to create a playing card games with these specifics endpoints

- `GET /api/v1/deck/:deckId`
- `PUT /api/v1/deck` or `PUT /api/v1/deck?cards=AS,KD,AC,2C,KH`
- `/draw-card`

## Dependencies

* Go 1.16.3

## Build project
```console
$ unzip <project_folder>
$ cd <project_folder>
$ sh build.sh
```

## Run tests
```console
$ cd <project_folder>
$ sh test.sh
```

## Start the server
```console
$ cd <project_folder>
$ sh start.sh
```
* Open the browser on `localhost:3000`
