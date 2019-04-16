package model

type Schedule struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TimeSpan  string `json:"timeSpan"`
}

type Job struct {
	ID             string   `json:"id"`
	ScheduleConfig Schedule `json:"schduleConfig"`
}
