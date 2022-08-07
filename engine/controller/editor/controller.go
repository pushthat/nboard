package editor

import (
	context "context"

	"github.com/pushthat/nboard-engine/internal/kv"
)

// Controller for editors
type Controller struct {
	kv kv.API
}

func New(kv kv.API) *Controller {
	return &Controller{
		kv: kv,
	}
}

type TaskType int

const (
	UNKNOWN TaskType = iota
	FREEFORM
)

type Task struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        TaskType `json:"type"`
}

// Onboarding struct
type Onboarding struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
}

// Index of editors
// GET /editor
func (c *Controller) Index(ctx context.Context) (editors []*Onboarding, err error) {
	return []*Onboarding{
		{
			ID:          1,
			Name:        "Marketing onboarding",
			Description: "Natalia's team onboarding",
		},
	}, nil
}

// New returns a view for creating a new editor
// GET /editor/new
func (c *Controller) New(ctx context.Context) {
}

// Create editor
// POST /editor
func (c *Controller) Create(ctx context.Context) (editor *Onboarding, err error) {
	return &Onboarding{
		ID: 0,
	}, nil
}

// Show editor
// GET /editor/:id
func (c *Controller) Show(ctx context.Context, id int) (editor *Onboarding, err error) {
	return &Onboarding{
		ID:          id,
		Name:        "Marketing onboarding",
		Description: "Natalia's team onboarding",
	}, nil
}

// Edit returns a view for editing a editor
// GET /editor/:id/edit
func (c *Controller) Edit(ctx context.Context, id int) (editor *Onboarding, err error) {
	return &Onboarding{
		ID:          id,
		Name:        "Marketing onboarding",
		Description: "Natalia's team onboarding",
		Tasks: []Task{
			{
				ID:          1,
				Name:        "play games",
				Description: "play games with me",
				Type:        FREEFORM,
			},
			{
				ID:          2,
				Name:        "run",
				Description: "run with me",
				Type:        FREEFORM,
			},
		},
	}, nil
}

// Update editor
// PATCH /editor/:id
func (c *Controller) Update(ctx context.Context, id int) (editor *Onboarding, err error) {
	return &Onboarding{
		ID: id,
	}, nil
}

// UpdateTask editor
// PATCH /editor/:id/task
func (c *Controller) UpdateTask(ctx context.Context, id int) (editor *Onboarding, err error) {
	return &Onboarding{
		ID: id,
	}, nil
}

// Delete editor
// DELETE /editor/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
