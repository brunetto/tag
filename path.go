package tag

import (
	"errors"
)

type Path []TagID

func (p Path) IsEmpty() bool {
	return len(p) == 0
}

// If I only use GetParent I don't need a path
// func (p Path) GetFirst() (TagID, error) {
// 	if p.IsEmpty() {
// 		return TagID{}, errors.New("can't get first, path is empty")
// 	}

// 	return p[0], nil
// }

func (p Path) GetParent() (TagID, error) {
	if p.IsEmpty() {
		return TagID{}, errors.New("can't get parent, path is empty")
	}

	return p[len(p)-1], nil
}

func (p Path) ToTagIDs() []TagID {
	return p
}
