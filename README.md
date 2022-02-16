<h1> Deck Handler API </h1>

Simple Restful API implemented in Go to simulate a deck of cards with the functions of creating a deck, opening a deck, and drawing from a deck.

<br>

How to run
=====
1) If you haven't installed go, install go https://go.dev/dl/
1) Open terminal and clone this repo using git clone
2) cd into the cloned repo
3) type `go run main.go`
4) API is running under `localhost:8080`

**Running tests**

1) Follow above methods until line 4
2) Run the following `go test -v ./...`

**Building**

1) If you haven't installed go, install go https://go.dev/dl/
1) Open terminal and clone this repo using git clone
2) cd into the cloned repo
3) type `go build -o DeckAPI`
4) Now you can run an executable called DeckAPI

Methods
======

## Create Deck
Creates the standard 52-card deck. <br> Can choose the option to have it shuffled and/or select specific cards.

## Endpoint

**{url}/deck/create**

| Params        | values          
| ------------- |:-------------:|
| shuffle      | true/false |
| cards      | string[]  |

### Returns

- uuid - unique specifier for deck id as a string
- shuffled - deck property specifying the deck was shuffled or not as a boolean
- remaining - total cards remaining in the deck

### Examples

*{url}/deck/create?shuffle=false*

Response
```json
{
    "uuid":"5bc933f1-6736-4e2a-903c-1d8360795821",
    "shuffled":false,
    "remaining":52
}
```
<br>

*{url}/deck/create?shuffle=false&cards=AS,KD,AC,2C,KH*

Response
```json
{
    "uuid":"760e1c03-f52f-47f2-a85f-ad8613755db0",
    "shuffled":false,
    "remaining":5
}
```

### Open Deck

Opens the deck with the specified deck id as uuid.

**{url}/deck/open**

| Params        | values          
| ------------- |:-------------:|
| uuid      | string  |

### Returns
- uuid - unique specifier for deck id as a string
- shuffled - deck property specifying the deck was shuffled or not as a boolean
- remaining - total cards remaining in the deck as an integer
- cards - all remaining cards as card objects
### Examples

{url}/deck/open?uuid=760e1c03-f52f-47f2-a85f-ad8613755db0

Response
```json
{
    "uuid":"760e1c03-f52f-47f2-a85f-ad8613755db0",
    "shuffled":false,
    "remaining":5,
    "cards":[
        {
            "value":"ACE",
            "suit":"SPADE",
            "code":"AS"
        },
        {
            "value":"KING",
            "suit":"DIAMOND",
            "code":"KD"
        },
        {
            "value":"ACE",
            "suit":"CLUB",
            "code":"AC"
        },
        {
            "value":"2",
            "suit":"CLUB",
            "code":"2C"
        },
        {
            "value":"KING",
            "suit":"HEART",
            "code":"KH"
        }
    ]
}
```


### Draw Card

**{url}/deck/draw**

| Params        | values          
| ------------- |:-------------:|
| uuid      | string  |
| count      | int  |

### Returns

- all the cards drawn as a cards object

### Examples

{url}/deck/draw?uuid=760e1c03-f52f-47f2-a85f-ad8613755db0&count=5

Response
```json
{ 
    "cards": [
        {
            "value":"KING",
            "suit":"HEART",
            "code":"KH"
        },
        {
            "value":"2",
            "suit":"CLUB",
            "code":"2C"
        },
        {
            "value":"ACE",
            "suit":"CLUB",
            "code":"AC"
        },
        {
            "value":"KING",
            "suit":"DIAMOND",
            "code":"KD"
        },
        {
            "value":"ACE",
            "suit":"SPADE",
            "code":"AS"
        }
    ]
}
```
