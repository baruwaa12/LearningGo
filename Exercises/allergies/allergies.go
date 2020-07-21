package allergies



// Key type integer
type Key uint

const (
	// Eggs Key for allergies
	Eggs Key = iota
	// Peanuts Key for allergies
	Peanuts
	// Shellfish Key for allergies
	Shellfish
	// Strawberries Key for allergies
	Strawberries
	// Tomatoes Key for allergies
	Tomatoes
	// Chocolate Key for allergies
	Chocolate
	// Pollen key for allergies
	Pollen
	// Cats key for allergies
	Cats
)

var allergies = []string{
	"Eggs",
	"Peanuts",
	"Shellfish", 
	"Strawberries",
	"Tomatoes", 
	"Chocolate",
	"Pollen", 
	"Cats", 
}

// Names function 
func Names(k Key) []string {
	var result []string
	for i := 0; i < len(allergies); i++ {
		if k&(1<<uint(i)) != 0 {
			result = append(result, allergies[i])
		}
	}
	return result
}

