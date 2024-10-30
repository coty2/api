package domain

import (
	v "suffgo/internal/settingRoom/domain/valueObjects"
	sv "suffgo/internal/shared/domain/valueObjects"
)

type (
	SettingRoom struct {
		id            *sv.ID
		privacy       v.Privacy
		proposalTimer v.ProposalTimer
		quorum        v.Quorum
		timeAndDate   v.TimeAndDate
		voterLimit    v.VoterLimit
	}

	SettingRoomDTO struct {
		ID            uint `json:"id"`
		Privacy       bool `json:"privacy"`
		ProposalTimer int  `json:"proposaltimer"`
		Quorum        int  `json:"quorum"`
		Time          int  `json:"time"`
		Date          int  `json:"date"`
		VoterLimit    int  `json:"voterlimit"`
	}

	SettingRoomCreateRequest struct {
		Privacy       bool `json:"privacy"`
		ProposalTimer int  `json:"proposaltimer"`
		Quorum        int  `json:"quorum"`
		Time          int  `json:"time"`
		Date          int  `json:"date"`
		VoterLimit    int  `json:"voterlimit"`
	}
)

func NewSettingRoom(
	id *sv.ID,
	privacy v.Privacy,
	proposalTimer v.ProposalTimer,
	quorum v.Quorum,
	timeAndDate v.TimeAndDate,
	voterLimit v.VoterLimit,
) *SettingRoom {
	return &SettingRoom{
		id:            id,
		privacy:       privacy,
		proposalTimer: proposalTimer,
		quorum:        quorum,
		timeAndDate:   timeAndDate,
		voterLimit:    voterLimit,
	}
}

func (s *SettingRoom) ID() sv.ID {
	return *s.id
}

func (s *SettingRoom) Privacy() v.Privacy {
	return s.privacy
}

func (s *SettingRoom) ProposalTimer() v.ProposalTimer {
	return s.proposalTimer
}

func (s *SettingRoom) Quorum() v.Quorum {
	return s.quorum
}

func (s *SettingRoom) TimeAndDate() v.TimeAndDate {
	return s.timeAndDate
}

func (s *SettingRoom) VoterLimit() v.VoterLimit {
	return s.voterLimit
}
