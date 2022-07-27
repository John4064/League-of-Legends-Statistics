package main

//Match Data structures

type matchDTO struct {
	metadata metadataDTO
	info     infoDTO
}

type infoDTO struct {
	gameCreation       int64
	gameDuration       int64
	gameEndTimestamp   int64
	gameId             int64
	gameMode           string
	gameName           string
	gameStartTimestamp int64
	gameType           string
	gameVersion        string
	mapId              int
	participants       []participantDTO
	platformId         string
	queueId            int
	teams              []teamDTO
	tournamentCode     string
}

type metadataDTO struct {
	dataVersion  string
	matchId      string
	participants []string
}

type participantDTO struct {
}
type teamDTO struct {
}
