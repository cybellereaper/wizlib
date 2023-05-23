package wizlib

import (
	"fmt"
	"strconv"
	"time"
)

// RaidMember represents a member of a raid.
type RaidMember struct {
	RaidPosition string `json:"raid_position"`
	Backup       bool   `json:"backup"`
}

// IsBackup checks if a raid member is a backup.
func (r *RaidMember) IsBackup() bool {
	return r.Backup
}

// RaidRepository provides methods for accessing raid data.
type RaidRepository interface {
	GetRaid(guildID string) (*Raid, error)
	SaveRaid(raid *Raid) error
}

// RaidService provides methods for performing raid-related operations.
type RaidService struct {
	repository RaidRepository
}

// NewRaidService creates a new instance of RaidService.
func NewRaidService(repository RaidRepository) *RaidService {
	return &RaidService{
		repository: repository,
	}
}

// GetRaid retrieves a raid by guild ID.
func (s *RaidService) GetRaid(guildID string) (*Raid, error) {
	return s.repository.GetRaid(guildID)
}

// SaveRaid saves a raid.
func (s *RaidService) SaveRaid(raid *Raid) error {
	return s.repository.SaveRaid(raid)
}

// Raid represents a raid with multiple gates.
type Raid struct {
	GuildID string `json:"guild_id"`
	Gates   []Gate `json:"gates"`
}

// GetGate retrieves a gate from the raid based on the date.
func (r *Raid) GetGate(date string) *Gate {
	for i := range r.Gates {
		if r.Gates[i].Date == date {
			return &r.Gates[i]
		}
	}
	return nil
}

// AddGate adds a new gate to the raid with the specified date if it doesn't already exist.
func (r *Raid) AddGate(date string) {
	if r.GetGate(date) == nil {
		r.Gates = append(r.Gates, Gate{Date: date, Status: 0xFFD700, Members: make(map[string]*RaidMember)})
	}
}

// GetGate retrieves a specific gate from the raid based on the gate number.
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

// Gate represents a gate in a raid.
type Gate struct {
	Status  int64                  `json:"status"`
	Date    string                 `json:"date"`
	Members map[string]*RaidMember `json:"members"`
}

// GetMember retrieves a raid member from the gate based on the user ID.
func (g *Gate) GetMember(userID string) *RaidMember {
	member, ok := g.Members[userID]
	if !ok {
		return nil
	}
	return member
}

// AddMember adds a new raid member to the gate.
func (g *Gate) AddMember(userID string, raidPosition string, backup bool) {
	g.Members[userID] = &RaidMember{RaidPosition: raidPosition, Backup: backup}
}

// RemoveMember removes a raid member from the gate.
func (g *Gate) RemoveMember(userID string) {
	delete(g.Members, userID)
}

// TimeFormatter provides methods for formatting time.
type TimeFormatter interface {
	ParseTime(timeString string) (string, error)
}

// DefaultTimeFormatter is a default implementation of TimeFormatter.
type DefaultTimeFormatter struct {
	layout string
}

// NewDefaultTimeFormatter creates a new instance of DefaultTimeFormatter.
func NewDefaultTimeFormatter(layout string) *DefaultTimeFormatter {
	return &DefaultTimeFormatter{
		layout: layout,
	}
}

// ParseTime parses a time string into the desired format.
func (f *DefaultTimeFormatter) ParseTime(timeString string) (string, error) {
	parsedTime, err := time.Parse(f.layout, timeString)
	if err != nil {
		return "", fmt.Errorf("invalid time format: %s", f.layout)
	}
	return strconv.FormatInt(parsedTime.Unix(), 10), nil
}

// TimeService provides methods for working with time.
type TimeService struct {
	formatter TimeFormatter
}

// NewTimeService creates a new instance of TimeService.
func NewTimeService(formatter TimeFormatter) *TimeService {
	return &TimeService{
		formatter: formatter,
	}
}

// ParseTime parses a time string into the desired format.
func (s *TimeService) ParseTime(timeString string) (string, error) {
	return s.formatter.ParseTime(timeString)
}
