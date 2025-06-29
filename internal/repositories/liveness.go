package repositories

type LivenessRepository interface {
	IsAlive() bool
}

type StaticLivenessRepository struct{}

func (StaticLivenessRepository) IsAlive() bool {
	return true
}
