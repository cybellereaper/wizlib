package wizlib

import (
	"fmt"
	"strconv"
	"time"
)

const (
	timeFormat = "01/02/2006 03:04:05 PM MST"
	timeDesc   = "Set a time and date format: " + timeFormat
)

type RaidMember struct {
	RaidPosition string `bson:"raid_position"`
	Backup       bool   `bson:"backup"`
}

func (r *RaidMember) IsBackup() bool {
	return r.Backup
}

type Raid struct {
	GuildID string `bson:"guild_id"`
	Gates   []Gate `bson:"gates"`
}

func (r *Raid) GetGate(date string) *Gate {
	for i := range r.Gates {
		if r.Gates[i].Date == date {
			return &r.Gates[i]
		}
	}
	return nil
}

func (r *Raid) AddGate(date string) {
	if r.GetGate(date) == nil {
		r.Gates = append(r.Gates, Gate{Date: date, Status: 0xFFD700, Members: make(map[string]*RaidMember)})
	}
}

func GetGate(raid *Raid, gateNum int) (*Gate, error) {
	if gateNum < 1 || gateNum > 3 {
		return nil, fmt.Errorf("invalid gate number: %d. Expected a number between 1 and 3", gateNum)
	}
	gateIndex := gateNum - 1
	if len(raid.Gates) <= gateIndex {
		return nil, fmt.Errorf("Gate %d not found in raid", gateNum)
	}
	return &raid.Gates[gateIndex], nil
}

type Gate struct {
	Status  int64                  `bson:"status"`
	Date    string                 `bson:"date"`
	Members map[string]*RaidMember `bson:"members"`
}

func (g *Gate) GetMember(userID string) *RaidMember {
	member, ok := g.Members[userID]
	if !ok {
		return nil
	}
	return member
}

func (g *Gate) AddMember(userID string, raidPosition string, backup bool) {
	g.Members[userID] = &RaidMember{RaidPosition: raidPosition, Backup: backup}
}

func (g *Gate) RemoveMember(userID string) {
	delete(g.Members, userID)
}

func TimeLayout(tim string) (string, error) {
	parsedTime, err := time.Parse(timeFormat, tim)
	if err != nil {
		return "", fmt.Errorf("invalid time format: %s", timeFormat)
	}
	return strconv.FormatInt(parsedTime.Unix(), 10), nil
}
