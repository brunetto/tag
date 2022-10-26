package tag

import (
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
	Parent           TagID
	Status           Status
}

type TagInsert struct {
	ClassificationID ClassificationID
	Name             LocalizedValue
	Parent           TagID
}

func validateName(name LocalizedValue) error {
	if len(name) == 0 {
		return fmt.Errorf("missing name")
	}
	return nil
}

func createTag(t TagInsert) (Tag, error) {
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
		Parent:           t.Parent,
		Status:           Ready,
	}
	return tag, nil
}

func updateTagName(tag Tag, name LocalizedValue) (Tag, error) {
	err := validateName(name)
	if err != nil {
		return Tag{}, err
	}
	tag.Name = name
	return tag, nil
}
