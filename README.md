# Unnatended test

API to provide the functionalities to create a deck that provides three endpoints:

---
- `GET /deck/:deckId`

Open a deck by your id. Example:
```console
$ curl localhost:3000/deck/6ab9cd50-24b3-45d2-93df-76b627a27c5e

  {
    "deck_id": "6ab9cd50-24b3-45d2-93df-76b627a27c5e",
    "shuffled": false,
    "remaining": 3,
    "cards": [
      {
        "value": "ACE",
        "suit": "SPADES",
        "code": "AS",
      },
      {
        "value": "KING",
        "suit": "HEARTS",
        "code": "KH",
      },
      {
        "value": "8",
        "suit": "CLUBS",
        "code": "8C",
      },
    ]
  }
```
---
- `POST /deck` or `POST /deck?cards=AS,KD,AC,2C,KH` or `POST /deck?cards=AS,KD&shuffle=true`

Create a deck passing or not the cards you would like in the deck and if the deck is shuffled or not, by default the deck is not shuffled. Example:

> Create a standard deck with 52 playing cards:
```console
$ curl -X POST localhost:3000/deck

 {
    "deck_id": "6ab9cd50-24b3-45d2-93df-76b627a27c5e",
    "shuffled": false,
    "remaining": 52
 }
```

> Create a custom deck:
```console
$ curl -X POST --url "localhost:3000/deck?cards=AS,KD,QH"

 {
    "deck_id": "db940885-274c-4d63-a1db-17088d4f73d4",
    "shuffled": false,
    "remaining": 3
 }
```

> Create a shuffled custom deck: 
```console
$ curl -X POST --url "localhost:3000/deck?cards=AS,KD,QH&shuffle=true"

 {
    "deck_id": "e7cb2982-6e76-49c0-a87c-1f95dffb5e2c",
    "shuffled": true,
    "remaining": 3
 }
```
---
- `POST /deck/:deckId/:count`

Draw card(s) from a deck. Example:

```console
$ curl -X POST --url "localhost:3000/deck/e7cb2982-6e76-49c0-a87c-1f95dffb5e2c/1"

 {
    "cards": [
      {
        "value": "ACE",
        "suit": "SPADES",
        "code": "AS",
      },
 }
```
---

## Dependencies
* Go 1.17.9

---

##How to run the project

#### Build project
```console
$ make build
```
#### Start the server
```console
$ make start
```

_The API will start in the port 3000_

---

###_If you would like to run the tests_
#### Run tests
```console
$ make test
```