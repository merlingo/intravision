package types

import (
	"time"
)

func (task Task) Validate() (err error) {
	_, err = task.GetDaysUntilDeadline()
	if err != nil {
		return err
	}
	if task.State == 2 {
		_, err = task.GetAllDays()
		if err != nil {
			return err
		}
	}
	return
}
func (task Task) GetDaysUntilDeadline() (days int, err error) {

	if task.GetBeginTask() >= task.GetDeadline() {
		return -1, ErrDeadlineConfliction
	}
	diff := time.UnixMilli(int64(task.Deadline)).Sub(time.UnixMilli(int64(task.BeginTask)))

	days = int(diff.Hours() / 24)
	return
}
func (task Task) GetAllDays() (days int, err error) {

	if task.GetBeginTask() >= task.GetFinishTask() {
		return -1, ErrDeadlineConfliction
	}
	diff := time.UnixMilli(int64(task.Deadline)).Sub(time.UnixMilli(int64(task.BeginTask)))

	days = int(diff.Hours() / 24)
	return
}
