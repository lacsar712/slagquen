package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("slagquen: invalid identifier")
	ErrNotFound        = errors.New("slagquen: entity not found")
	ErrConflict        = errors.New("slagquen: state conflict")
	ErrInterlock       = errors.New("slagquen: interlock denied")
	ErrMoistureHold    = errors.New("slagquen: moisture hold active")
	ErrAirflowSetpoint = errors.New("slagquen: airflow setpoint violation")
	ErrFanFault        = errors.New("slagquen: fan fault")
	ErrScheduleEmpty   = errors.New("slagquen: schedule empty")
	ErrGradient        = errors.New("slagquen: moisture gradient violation")
	ErrSlagDrift   = errors.New("slagquen: moisture drift exceeded")
	ErrQuenchTrip    = errors.New("slagquen: heat overtemperature")
	ErrGranHold    = errors.New("slagquen: gradient hold not satisfied")
	ErrContextCanceled = errors.New("slagquen: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("slagquen %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("slagquen %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
