package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Schedule struct {
	bun.BaseModel `bun:"table:schedule"`
	ID            int64     `json:"id"`
	RelationID    int64     `json:"relation_id"`
	SpecificKey   string    `json:"specific_key"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Content       string    `json:"content"`
	IsActive      bool      `json:"is_active"`
	PublishDate   time.Time `json:"published_at"`
	ExpiredDate   time.Time `json:"expired_date"`
}

type ScheduleResponse struct {
	bun.BaseModel `bun:"table:schedule"`
	ID            int64     `json:"id"`
	RelationID    int64     `json:"relation_id"`
	SpecificKey   string    `json:"specific_key"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Content       string    `json:"content"`
	IsActive      bool      `json:"is_active"`
	PublishDate   time.Time `json:"published_at"`
	ExpiredDate   time.Time `json:"expired_date"`
}

func NewScheduleResponse(schedule *Schedule) ScheduleResponse {
	return ScheduleResponse{
		ID:          schedule.ID,
		RelationID:  schedule.RelationID,
		SpecificKey: schedule.SpecificKey,
		Type:        schedule.Type,
		Name:        schedule.Name,
		Image:       schedule.Image,
		Content:     schedule.Content,
		IsActive:    schedule.IsActive,
		PublishDate: schedule.PublishDate,
		ExpiredDate: schedule.ExpiredDate,
	}
}
