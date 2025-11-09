package vo

import "time"

type ToDoVO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	EndTime     time.Time `json:"end_time"`
}
