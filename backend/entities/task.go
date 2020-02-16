package entities

import "time"

type Task struct {
	Name                 string    `json:"name"`
	StartTime            time.Time `json:"start_time"`
	EndTime              time.Time `json:"end_time"`
	WorkLength           int       `json:"work_length"`
	BreakLength          int       `json:"break_length"`
	Category             string    `json:"category"`
	EndedEarly           bool      `json:"ended_early"`
	EarlyTerminationTime time.Time `json:"early_termination_time"`
}

func (t *Task) SetName(name string) *Task {
	t.Name = name
	return t
}

func (t *Task) SetStart(s time.Time) *Task {
	t.StartTime = s
	return t
}

func (t *Task) SetEnd(e time.Time) *Task {
	t.EndTime = e
	return t
}

func (t *Task) SetWorkLength(length int) *Task {
	t.WorkLength = length
	return t
}

func (t *Task) SetBreakLength(length int) *Task {
	t.BreakLength = length
	return t
}
