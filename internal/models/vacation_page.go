package models

type VacationPage struct {
	Vacations []*Vacation `json:"vacations"`
	PageNum   uint64      `json:"page_num"`
}
