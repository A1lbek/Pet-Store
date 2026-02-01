package pet

import "sync"

type Repository struct {
	mu     sync.Mutex
	pets   map[int]Pet
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		pets:   make(map[int]Pet),
		nextID: 1,
	}
}

func (r *Repository) Create(p Pet) Pet {
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	p.Status = "available"
	r.pets[r.nextID] = p
	r.nextID++

	return p
}

func (r *Repository) GetAll() []Pet {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]Pet, 0, len(r.pets))
	for _, p := range r.pets {
		result = append(result, p)
	}
	return result
}

func (r *Repository) GetByID(id int) (Pet, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.pets[id]
	return p, ok
}
