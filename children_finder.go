package tag

type ChildrenFinder struct {
	tags []Tag
}

func (cf ChildrenFinder) HasChildren(tag Tag) bool {
	for _, t := range cf.tags {
		p, err := t.GetParent()
		if err != nil {
			// t is root
			continue
		}

		if p == tag.ID {
			return true
		}
	}

	return false
}
