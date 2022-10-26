package tag

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TagInsert struct {
	ClassificationID ClassificationID
	Name             LocalizedValue
	Parent           TagID
	Ancestors        Path // should keep?
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
		Ancestors:        t.Ancestors,
		Status:           Ready,
	}

	return tag, nil
}

func renameTag(tag Tag, name LocalizedValue) (Tag, error) {
	err := validateName(name)
	if err != nil {
		return Tag{}, err
	}

	tag.Name = name
	return tag, nil
}

func moveTag(tag Tag, p Path, cf ChildrenFinder) (Tag, error) {
	if cf.HasChildren(tag) {
		return tag, errors.New("can't move tag, is not leaf")
	}

	tag.Ancestors = p

	return tag, nil
}
