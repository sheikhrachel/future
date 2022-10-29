package rds

type FutureTable string

const (
	Appointments FutureTable = "future.appointments"
)

type AppointmentQueryOpts struct {
	StartUnix     int64
	EndUnix       int64
	TrainerId     int
	AppointmentId int
	UserId        int
}
