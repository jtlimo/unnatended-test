<!-- 
Some problems that I could not resolve in time and improvements I saw in the code

- The tests fail when running all package tests but individually they pass, I thinking it's about a concurrency or because the struct only grow up when I use the constructor method to build them.

- When I build a deck with the constructor method NewDeck and next I create other deck with the same constructor method the Card field grow up and duplicate. I try to research about that but I don't know why this occurs.

- I would like to improve my tests, with setup and cleanup the state between tests;


 -->

# Unnantended Test

API to provide the functionalities to create a playing card games with these specifics endpoints

- `/create-deck`
- `/open-deck`
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
* Open the browser on `localhost:9091`
