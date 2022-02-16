package card

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func New(value string, suit string) Card {
	return Card{value, suit, string(value[0]) + string(suit[0])}
}
