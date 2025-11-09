package dto

import "time"

type ToDoAddReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EndTime     time.Time `json:"end_time"`
}
