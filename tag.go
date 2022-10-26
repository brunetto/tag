package tag

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TagID uuid.UUID
type ClassificationID string
type LocalizedValue map[string]string

type Status int64

const (
	Ready Status = iota
	Combined
	Merged
	Merging
	Locked
)

type Tag struct {
	ID               TagID
	ClassificationID ClassificationID
	Name             LocalizedValue
	Ancestors        Path // TODO: should keep?
	Status           Status
}

func NewTag(t TagInsert) (Tag, error) {
	if len(t.ClassificationID) == 0 {
		return Tag{}, fmt.Errorf("missing classification id")
	}

	err := validateName(t.Name)
	if err != nil {
		return Tag{}, err
	}

	tag := Tag{
		ID:               TagID(uuid.New()),
		ClassificationID: t.ClassificationID,
		Name:             t.Name,
		Ancestors:        t.Ancestors,
		Status:           Ready,
	}

	return tag, nil
}

func (tag Tag) IsRoot() bool {
	return tag.Ancestors.IsEmpty()
}

func (tag Tag) GetParent() (TagID, error) {
	if tag.IsRoot() {
		return TagID{}, errors.New("can't get parent, I'm root")
	}

	return tag.Ancestors.GetParent()
}

// TODO: Do we need this?
func (tag Tag) GetAncestors() Path {
	return tag.Ancestors
}
