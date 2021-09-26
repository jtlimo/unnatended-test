**Yoke Cards**

You will need to provide the following methods to your API to handle cards
and decks:

- Create a new Deck
- Open a Deck
- Draw a Card

**Background** 

> Create a new Deck

It would create the standard 52-card deck of French playing cards, It includes
all thirteen ranks in each of the four suits: clubs (♣), diamonds (♦), hearts (♥)
and spades (♠). You don't need to worry about Joker cards for this
assignment.

You should allow the following options to the request:

- the deck to be shuffled or not — by default the deck is sequential: A
  spades, 2-spades, 3-spades... followed by diamonds, clubs, then hearts.

- the deck to be full or partial — by default it returns the standard 52
cards, otherwise the request would accept the wanted cards like this example:

`?cards=AS,KD,AC,2C,KH`

_The response needs to return a JSON that would include:_

- the deck id (UUID)
- the deck properties like shuffled (boolean) and total cards remaining in this deck (integer)

```
  {
    "deck_id": "7dd13273-fabb-4223-9df6-9646c9473880",
    "shuffled": false,
    "remaining": 30
  }
```

> Open a Deck

It would return a given deck by its UUID. If the deck was not passed over or is
invalid it should return an error. This method will "open the deck", meaning that
it will list all cards by the order it was created.

_The response needs to return a JSON that would include:_

- the deck id (UUID)
- the deck properties like shuffled (boolean) and total cards remaining in this deck (integer)
- all the remaining cards (card object)

```
  {
    "deck_id": "7dd13273-fabb-4223-9df6-9646c9473880",
    "shuffled": false,
    "remaining": 30,
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

> Draw a card

I would draw a card(s) of a given Deck. If the deck was not passed over or
invalid it should return an error. A count parameter needs to be provided to
define how many cards to draw from the deck.

_The response needs to return a JSON that would include:_

- all the drawn cards (card object)

```
  {
    "cards": [
      {
        "value": "QUEEN",
        "suit": "HEARTS",
        "code": "QH",
      },
      {
        "value": "4",
        "suit": "DIAMONDS",
        "code": "4D",
      },
    ] 
  }
```