package main

type Years [12]map[int]Day

type Day struct {
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Day     int    `json:"day"`
	Weekday int    `json:"weekday"`
	Date    string `json:"Date"`
	Work    bool   `json:"work"`
	Remark  string `json:"remark"`
}
