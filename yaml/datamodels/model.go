package datamodels

type Developer struct {
	ID         string             `yaml:"id"`
	Sector     int                `yaml:"sector"`
	Salary     float32            `yaml:"salary"`
	Tasks      []string           `yaml:"tasks"`
	DailyHours []int              `yaml:"dailyHours"`
	Languages  map[string]float32 `yaml:"languages"`
}

type Company struct {
	ID      string             `yaml:"id"`
	Active  bool               `yaml:"active"`
	People  []string           `yaml:"people"`
	Budget  int64              `yaml:"budget"`
	Sectors map[string]float32 `yaml:"sectors"`
}

type CompanyGroup struct {
	Companies []Company `yaml:"companies"`
}
