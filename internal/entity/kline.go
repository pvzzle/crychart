package entity

type Kline struct {
	OpenTime  int64  `json:"openTime" bson:"openTime"`
	Open      string `json:"open" bson:"open"`
	High      string `json:"high" bson:"high"`
	Low       string `json:"low" bson:"low"`
	Close     string `json:"close" bson:"close"`
	Volume    string `json:"volume" bson:"volume"`
	CloseTime int64  `json:"closeTime" bson:"closeTime"`
}

type Klines []Kline

type KlineOptions struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Limit    int    `json:"limit"`
}
