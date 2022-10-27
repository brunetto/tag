package tag

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
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

func makeAliasOf(tagA Tag, tagB Tag, cf ChildrenFinder) (Tag, error) {
	if cf.HasChildren(tagA) {
		return tagA, errors.Errorf("can't alias tag, source tag %v is not leaf", tagA.ID)
	}

	if cf.HasChildren(tagB) {
		return tagA, errors.Errorf("can't alias tag, destination tag %v is not leaf", tagB.ID)
	}
	newTag := tagB.Clone()
	newTag.ID = tagA.ID

	return newTag, nil
}

func samePayload(tagA Tag, tagB Tag) bool {
	return reflect.DeepEqual(tagA.Ancestors, tagB.Ancestors) &&
		tagA.ClassificationID == tagB.ClassificationID &&
		reflect.DeepEqual(tagA.Name, tagB.Name) &&
		tagA.Status == tagB.Status
}
