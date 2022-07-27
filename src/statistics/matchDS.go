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
type perksDTO struct {
	statPerks perkStatsDTO
	styles    []perkStyleDTO
}

type perkStatsDTO struct {
	defense int
	flex    int
	offense int
}

type perkStyleDTO struct {
	description string
	selections  []perkStyleSelectionDTO
	style       int
}

type perkStyleSelectionDTO struct {
	perk int
	var1 int
	var2 int
	var3 int
}
