package models

type Project struct {
    CompanyName  string `json:"company_name"`
    Terms       string `json:"terms"`
    Scope       string `json:"scope"`
    ValidVulns  string `json:"valid_vulnerabilities"`
    InvalidVulns string `json:"invalid_vulnerabilities"`
}

func NewProject() *Project {
    return &Project{}
}

