package model

import (
	"time"

	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type EventComment struct {
	Id               int64          `json:"id"`
	UserId           int64          `json:"user_id"`
	EventId          int64          `json:"event_id"`
	Comment          string         `json:"comment"`
	ParentId         null.Int       `json:"parent_id"`
	MentionedUserIds pq.Int64Array  `gorm:"type:int[];default:[]" json:"mentioned_user_ids"`
	CreatedAt        time.Time      `json:"created_at"`
	CreatedBy        int64          `json:"created_by"`
	UpdatedAt        time.Time      `json:"updated_at"`
	UpdatedBy        int64          `json:"updated_by"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	DeletedBy        null.Int       `json:"deleted_by"`
}
