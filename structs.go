package main

type Response struct {
	Transactid string `json:"transactid"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}
