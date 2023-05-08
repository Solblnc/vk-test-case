package models

import "time"

// Button - struct for VK button
type Button struct {
	OneTime bool `json:"one_time"`
	Inline  bool `json:"inline"`
	Buttons [][]struct {
		Action struct {
			Type    string `json:"type"`
			Label   string `json:"label"`
			Payload string `json:"payload"`
		} `json:"action"`
		Color string `json:"color"`
	} `json:"buttons"`
}

// Floor - to get beautiful data
type Floor struct {
	Value float64
	Time  time.Time
}

// Stats - statistic of nft token
type Stats struct {
	Symbol      string  `json:"symbol"`
	FloorPrice  float64 `json:"floorPrice"`
	ListedCount int     `json:"listedCount"`
	VolumeAll   float64 `json:"volumeAll"`
}

// Token - statistic of token
type Token struct {
	Type           string  `json:"type"`
	Timestamp      int64   `json:"timestamp"`
	BlockTimestamp int64   `json:"blockTimestamp"`
	MintAddress    string  `json:"mintAddress"`
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	FloorPrice     float64 `json:"floorPrice"`
	RarityStr      string  `json:"rarityStr"`
	Rank           int     `json:"rank"`
	Supply         int     `json:"supply"`
	TokenAddress   string  `json:"tokenAddress"`
	Seller         string  `json:"seller"`
}
