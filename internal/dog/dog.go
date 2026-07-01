package dog

import (
	"math/rand"
	"strings"
	"time"
)

// Dog describes one dog profile exposed by the CLI and HTTP API.
type Dog struct {
	Name        string   `json:"name"`
	Breed       string   `json:"breed"`
	Age         int      `json:"age"`
	Personality []string `json:"personality"`
}

// Repository keeps a small in-memory dog catalogue.
type Repository struct {
	dogs []Dog
	rnd  *rand.Rand
}

// NewRepository creates a repository with friendly default data.
func NewRepository() *Repository {
	return &Repository{
		dogs: []Dog{
			{Name: "Mochi", Breed: "Shiba Inu", Age: 3, Personality: []string{"curious", "independent", "loyal"}},
			{Name: "Bao", Breed: "Golden Retriever", Age: 5, Personality: []string{"gentle", "playful", "patient"}},
			{Name: "Nori", Breed: "Border Collie", Age: 2, Personality: []string{"smart", "energetic", "focused"}},
			{Name: "Tofu", Breed: "Corgi", Age: 4, Personality: []string{"cheerful", "bold", "food-motivated"}},
		},
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// All returns a copy of every dog profile.
func (r *Repository) All() []Dog {
	dogs := make([]Dog, len(r.dogs))
	copy(dogs, r.dogs)
	return dogs
}

// Random returns a random dog profile.
func (r *Repository) Random() Dog {
	return r.dogs[r.rnd.Intn(len(r.dogs))]
}

// FindByName searches for a dog by name, case-insensitively.
func (r *Repository) FindByName(name string) (Dog, bool) {
	for _, dog := range r.dogs {
		if strings.EqualFold(dog.Name, name) {
			return dog, true
		}
	}
	return Dog{}, false
}
