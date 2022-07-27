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
	assists                        int
	baronKills                     int
	bountyLevel                    int
	champExperience                int
	champLevel                     int
	championId                     int
	championName                   string
	championTransform              int
	consumablesPurchased           int
	damageDealtToBuildings         int
	damageDealtToObjectives        int
	damageDealthToTurrets          int
	damageSelfMitigated            int
	deaths                         int
	detectorWardsPlaced            int
	doubleKills                    int
	dragonKills                    int
	firstBloodAssist               bool
	firstBloodKill                 bool
	firstTowerAssist               bool
	firstTowerKill                 bool
	gameEndedInEarlySurrender      bool
	gameEndedInSurrender           bool
	goldEarned                     int
	individualPosition             string
	inhibitorKills                 int
	inhibitorTakedowns             int
	inhibitorsLost                 int
	item0                          int
	item1                          int
	item2                          int
	item3                          int
	item4                          int
	item5                          int
	item6                          int
	itemsPurchased                 int
	killingSprees                  int
	kills                          int
	lane                           string
	largestCriticalStrike          int
	largestKillingSpree            int
	largestMultiKill               int
	longestTimeSpentLiving         int
	magicDamageDealt               int
	magicDamageDealtToChampions    int
	magicDamageTaken               int
	neutralMinionsKilled           int
	nexusKills                     int
	nexusTakedowns                 int
	nexusLost                      int
	objectivesStolen               int
	objectivesStolenAssists        int
	participantId                  int
	pentaKills                     int
	perks                          perksDTO
	physicalDamageDealt            int
	physicalDamageDealtToChampions int
	physicalDamageTaken            int
	profileIcon                    int
	puuid                          string
	quadraKills                    int
	riotIdName                     string
	riotIdTagline                  string
	role                           string
	sightWardsBoughtInGame         int
	spell1Casts                    int
	spell2Casts                    int
	spell3Casts                    int
	spell4Casts                    int
	summoner1Casts                 int
	summoners1Id                   int
	Summoner2Casts                 int
	summoner2Id                    int
	summonerId                     string
	summonerLevel                  int
	summonerName                   string
	teamEarlySurrendered           bool
	teamId                         int
	teamPosition                   string
	timeCCingOthers                int
	timePlayed                     int
	totalDamageDealt               int
	totalDamageDealtToChampions    int
	totalDamageShieldedOnTeammates int
	totalDamageTaken               int
	totalHeal                      int
	totalHealsOnTeammates          int
	totalMinionsKilled             int
	totalTimeCCDealt               int
	totalTimeSpentDead             int
	totalUnitsHealed               int
	tripleKills                    int
	trueDamageDealt                int
	trueDamageDealtToChampions     int
	trueDamageTaken                int
	turretKills                    int
	turretTakedowns                int
	turretsLost                    int
	unrealKills                    int
	visionScore                    int
	visionWardsBoughtInGame        int
	wardsKilled                    int
	wardsPlaced                    int
	win                            bool
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

type teamDTO struct {
	bans       []banDTO
	objectives objectivesDTO
	teamId     int
	win        bool
}

type banDTO struct {
	championId int
	pickTurn   int
}

type objectivesDTO struct {
	baron      objectiveDTO
	champion   objectiveDTO
	dragon     objectiveDTO
	inhibitor  objectiveDTO
	riftHerald objectiveDTO
	tower      objectiveDTO
}
type objectiveDTO struct {
	first bool
	kills int
}
