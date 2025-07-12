package handler

import "fmt"

func errParamIsRequired(param, param_type string) error {
	return fmt.Errorf("param: %s(type %s) is required", param, param_type)
}

type CreateOppeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOppeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Remote == nil{
		return errParamIsRequired("remote", "boolean")
	}
	
	if r.Salary <= 0{
		return errParamIsRequired("salary","int64")
	}

	return nil
}

type CreateOppeningUpdate struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
} 

func (r *CreateOppeningUpdate) Validate() error {
	if r.Company != "" || r.Role != "" || r.Link != "" || r.Link != "" || r.Remote != nil || r.Salary < 0 {
		return nil
	}

	return fmt.Errorf("at least one field valid must be provied")
}


