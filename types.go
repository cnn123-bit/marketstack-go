package marketstack

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

type EODData struct {
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      float64 `json:"volume"`
	AdjHigh     float64 `json:"adj_high"`
	AdjLow      float64 `json:"adj_low"`
	AdjClose    float64 `json:"adj_close"`
	AdjOpen     float64 `json:"adj_open"`
	AdjVolume   float64 `json:"adj_volume"`
	SplitFactor float64 `json:"split_factor"`
	Dividend    float64 `json:"dividend"`
	Symbol      string  `json:"symbol"`
	Exchange    string  `json:"exchange"`
	Date        string  `json:"date"`
}

type IntradayData struct {
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Last     float64 `json:"last"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Symbol   string  `json:"symbol"`
	Exchange string  `json:"exchange"`
	Date     string  `json:"date"`
}

type StockExchange struct {
	Name        string `json:"name"`
	Acronym     string `json:"acronym"`
	MIC         string `json:"mic"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	Website     string `json:"website"`
	Timezone    struct {
		Timezone string `json:"timezone"`
		Abbr     string `json:"abbr"`
		AbbrDST  string `json:"abbr_dst"`
	} `json:"timezone,omitempty"`
}

type Ticker struct {
	Name          string         `json:"name"`
	Symbol        string         `json:"symbol"`
	HasIntraday   bool           `json:"has_intraday"`
	HasEOD        bool           `json:"has_eod"`
	Country       string         `json:"country"`
	StockExchange *StockExchange `json:"stock_exchange"`
}

type Currency struct {
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type Timezone struct {
	Timezone string `json:"timezone"`
	Abbr     string `json:"abbr"`
	AbbrDST  string `json:"abbr_dst"`
}
