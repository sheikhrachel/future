package rds

import (
	"fmt"
	"strconv"

	"github.com/sheikhrachel/future/api_common/call"
	"github.com/sheikhrachel/future/api_common/util/errutil"
	"github.com/sheikhrachel/future/model"
)

func (p *Postgres) GetAllBookings(
	cc call.Call,
	opts AppointmentQueryOpts,
) (appointments []model.Appointment, err error) {
	query := fmt.Sprintf("select * from %s "+
		"where trainer_id=%s",
		Appointments,
		strconv.Itoa(opts.TrainerId),
	)
	err = p.Psql.Select(&appointments, query)
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("GetAllBookings Select: %+v, err: %+v", query, err))
		return nil, err
	}
	return appointments, err
}

func (p *Postgres) GetBookingStartsInRange(
	cc call.Call,
	opts AppointmentQueryOpts,
) (appointments []model.Appointment, err error) {
	query := fmt.Sprintf("select * from %s "+
		"where trainer_id=%s "+
		"and (%s<=started_at and started_at<%s)",
		Appointments,
		strconv.Itoa(opts.TrainerId),
		strconv.FormatInt(opts.StartUnix, 10),
		strconv.FormatInt(opts.EndUnix, 10),
	)
	err = p.Psql.Select(&appointments, query)
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("GetBookingStartsInRange Select: %+v, err: %+v", query, err))
		return nil, err
	}
	return appointments, err
}

func (p *Postgres) AddNewAppointment(
	cc call.Call,
	opts AppointmentQueryOpts,
) (err error) {
	query := fmt.Sprintf("insert into %s "+
		"(appointment_id, user_id, trainer_id, started_at, ended_at) "+
		"values(%s, %s, %s, %s, %s)",
		Appointments,
		strconv.Itoa(opts.AppointmentId),
		strconv.Itoa(opts.UserId),
		strconv.Itoa(opts.TrainerId),
		strconv.FormatInt(opts.StartUnix, 10),
		strconv.FormatInt(opts.EndUnix, 10),
	)
	res, err := p.Psql.Exec(query)
	if !errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("AddNewAppointment Select: %+v, err: %+v", query, err))
		return err
	}
	cc.TraceF(fmt.Sprintf("res: %#v", res))
	return err
}

func (p *Postgres) DeleteAppointment(
	cc call.Call,
	opts AppointmentQueryOpts,
) (err error) {
	query := fmt.Sprintf("delete from %s "+
		"where appointment_id=%s",
		Appointments,
		strconv.Itoa(opts.AppointmentId),
	)
	res, err := p.Psql.Exec(query)
	if !errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("DeleteAppointment Select: %+v, err: %+v", query, err))
		return err
	}
	cc.TraceF(fmt.Sprintf("res: %#v", res))
	return err
}

func (p *Postgres) TrainerAlreadyBooked(
	cc call.Call,
	opts AppointmentQueryOpts,
) (exists bool, err error) {
	query := fmt.Sprintf("select count(1) from %s "+
		"where trainer_id=%s "+
		"and started_at=%s",
		Appointments,
		strconv.Itoa(opts.TrainerId),
		strconv.FormatInt(opts.StartUnix, 10),
	)
	var appointmentCount []interface{}
	err = p.Psql.Select(&appointmentCount, query)
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("GetAppointmentCount Select: %+v, err: %+v", query, err))
		return false, err
	}
	return appointmentCount[0].(int64) > 0, err
}

func (p *Postgres) AppointmentExists(
	cc call.Call,
	opts AppointmentQueryOpts,
) (exists bool, err error) {
	query := fmt.Sprintf("select count(1) from %s "+
		"where appointment_id=%s",
		Appointments,
		strconv.Itoa(opts.AppointmentId),
	)
	var appointmentCount []interface{}
	err = p.Psql.Select(&appointmentCount, query)
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("GetAppointmentCount Select: %+v, err: %+v", query, err))
		return false, err
	}
	return appointmentCount[0].(int64) > 0, err
}

func (p *Postgres) TrainerExists(
	cc call.Call,
	opts AppointmentQueryOpts,
) (exists bool, err error) {
	query := fmt.Sprintf("select count(1) from %s "+
		"where trainer_id=%s",
		Appointments,
		strconv.Itoa(opts.TrainerId),
	)
	var trainerCount []interface{}
	err = p.Psql.Select(&trainerCount, query)
	if errutil.HandleError(err) {
		cc.InfoF(fmt.Sprintf("TrainerExists Select: %+v, err: %+v", query, err))
		return false, err
	}
	return trainerCount[0].(int64) > 0, err
}
