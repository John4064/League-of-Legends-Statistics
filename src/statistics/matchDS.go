package main

//Match Data structures

type matchDTO struct {
	metadata metadataDTO
	info     infoDTO
}

type infoDTO struct {
}

type metadataDTO struct {
	dataVersion  string
	matchId      string
	participants []string
}
