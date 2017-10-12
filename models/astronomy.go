package models

type (
	AstronomyRepository interface {
		GetAstronomy(string) (*AstronomyDay, error)
	}

	AstronomyDay struct {
		Sunrise string
		Sunset  string
	}
)
