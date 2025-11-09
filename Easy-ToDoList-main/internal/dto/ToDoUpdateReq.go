package dto

import "time"

type ToDoUpdateReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	EndTime     time.Time `json:"end_time"`
}
