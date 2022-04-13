# godeck
Sample web api to manipulate a deck of cards

## Simple how-to
Clone the repo
`go build`
`./godeck`


### Endpoints
#### /deck/create
Creates a deck of cards, storing it in memory.
Available params:
| Param | Type | Description | Example
| ----------- |----------- | ----------- |----------- |
| cards | []string | Deck will only be generated with given cards | AS,AD
| shuffle | string | If deck should be shuffled upon creation | Any string will work, key only needs to exist

Example request:
```
curl --request POST 'localhost:8080/deck/create?cards=AS,AD&shuffle'
```
Return:
```
{
    "id": "0af572af-537b-416b-966d-43842d58a1bd",
    "remaining": 2,
    "shuffled": false
}
```
#### /deck/open
Opens a deck of cards
Available params:
| Param | Type | Description
| ----------- |----------- | ----------- 
| id | string | Deck UUID, returned on creation
Example request:
```
curl --request GET 'localhost:8080/deck/open?id=0e03f268-ba94-4842-af0e-68678720282d'
```
Return:
```
{
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "ACE",
            "suit": "DIAMONDS",
            "code": "AD"
        }
    ],
    "id": "0e03f268-ba94-4842-af0e-68678720282d",
    "remaining": 2,
    "shuffled": true
}
```

#### /deck/draw
Draws cards from given deck
Available params:
| Param | Type | Description
| ----------- |----------- | ----------- 
| id | string | Deck UUID, returned on creation
| amount| int | Number of cards to draw

Example request:
```
curl --request POST 'localhost:8080/deck/draw?id=b929c16b-a13e-495c-a570-2e108c99267d&amount=2'
```
Return:
```
{
    "cards": [
        {
            "value": "FOUR",
            "suit": "SPADES",
            "code": "FS"
        },
        {
            "value": "SIX",
            "suit": "HEARTS",
            "code": "SH"
        }
    ]
}
```


