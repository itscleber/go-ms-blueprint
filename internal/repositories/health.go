package repositories

type HealthRepository interface {
	Status() string
}

type StaticHealthRepository struct{}

func (StaticHealthRepository) Status() string {
	return "healthy"
}
