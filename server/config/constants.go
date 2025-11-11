package config

type Enrironment string

const (
	Development Enrironment = "dev"
	Staging     Enrironment = "stage" // used for testing before production
	Production  Enrironment = "prod"
)
