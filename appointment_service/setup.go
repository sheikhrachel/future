package appointment_service

import (
	"github.com/sheikhrachel/future/api_common/aws/rds"
	"github.com/sheikhrachel/future/api_common/call"
)

type AppointmentService struct {
	Postgres *rds.Postgres
}

func NewAppointmentService(cc call.Call) *AppointmentService {
	return &AppointmentService{
		Postgres: rds.InitPostgres(cc),
	}
}
