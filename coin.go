package main

type coin struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	PriceUSD float64 `json:"price_usd,string"`
	Delta1H  float64 `json:"percent_change_1h,string"`
	Delta1D  float64 `json:"percent_change_24h,string"`
	Delta7D  float64 `json:"percent_change_7d,string"`
	Holdings float64 `json:"holdings"`
	Worth    float64
}
