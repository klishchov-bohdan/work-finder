package models

type Vacation struct {
	Title          string `json:"title"`
	Salary         string `json:"salary"`
	SalaryDetail   string `json:"salary_detail"`
	Customer       string `json:"customer"`
	CustomerDetail string `json:"customer_detail"`
	Address        string `json:"address"`
	Conditions     string `json:"conditions"`
	Logo           string `json:"logo"`
}
