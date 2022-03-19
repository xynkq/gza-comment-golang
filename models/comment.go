package models

import "time"

type Comment struct {
	Id      string    `json:"id" binding:"required" gorm:"primary_key"`
	UserId  string    `json:"UserId" binding:"required" gorm:"ForeignKey:UserId"`
	BoardId int       `json:"boardId" binding:"required" gorm:"ForeignKey:BoardId"`
	Content string    `json:"content" binding:"required"`
	Created time.Time `json:"created" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	Updated time.Time `json:"updated" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
