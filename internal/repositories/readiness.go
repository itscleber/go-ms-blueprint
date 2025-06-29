package repositories

type ReadinessRepository interface {
	IsReady() bool
}

type StaticReadinessRepository struct{}

func (StaticReadinessRepository) IsReady() bool {
	// stub: add logic to check if the service is ready
	return true
}
