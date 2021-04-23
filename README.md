<!-- 
Some problems that I could not resolve in time and improvements I saw in the code

- I would like to make some refactoring to cleanup the method NewDeck, I see some duplication, and I could extract to a method to be more clear;

- About project architecture I would like to separate better the code in main.go to another packages like handlers and so on.

- The draw endpoint has a bug and not fully implemented, it's missing this requirement: If the deck was not passed over should return an error; The bug on this endpoint it's not implemented the upsert to replace the old deck, to the new one with the fields updated. Because of that when you get all the decks, you see the same deck in different states.
 -->

# Unnantended Test

API to provide the functionalities to create a playing card games with these specifics endpoints

- `GET /api/v1/deck/:deckId`
- `PUT /api/v1/deck` or `PUT /api/v1/deck?cards=AS,KD,AC,2C,KH`
- `PUT /api/v1/deck/:deckId/:count`

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

>You will see in your terminal something like that:

![Alt text](./assets/start-server.png?raw=true "Initializing the server")

>And now open the browser on `localhost:3000`
