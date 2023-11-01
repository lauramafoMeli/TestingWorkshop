package hunt

// NewTuna creates a new Tuna
func NewTuna(name string, speed float64) (t *Tuna) {
	t = &Tuna{
		name:  name,
		speed: speed,
	}
	return
}

// Tuna is a type of prey
type Tuna struct {
	// name of the tuna
	name string
	// speed of the tuna
	speed float64
}