package model

type Appointment struct {
	AppointmentId int    `json:"id" db:"appointment_id"`
	TrainerId     int    `json:"trainer_id" db:"trainer_id"`
	UserId        int    `json:"user_id" db:"user_id"`
	StartsAt      string `json:"started_at" db:"started_at"`
	EndsAt        string `json:"ended_at" db:"ended_at"`
}
