package tag

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestTag() Tag {
	return Tag{
		ID: TagID(uuid.New()), ClassificationID: ClassificationID("fake-classification"),
		Name: map[string]string{"it": "B"}, Status: Ready,
	}
}

func TestRenameTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		newName  LocalizedValue
		tag      Tag
		wantName LocalizedValue
		wantErr  bool
		errorMsg string
	}{
		{
			testName: "ok",
			newName:  map[string]string{"it": "C"},
			tag:      newTestTag(),
			wantName: map[string]string{"it": "C"},
			wantErr:  false,
		},
		{
			testName: "fail if new name is empty",
			newName:  nil,
			tag:      newTestTag(),
			wantName: map[string]string{"it": "C"},
			wantErr:  true, errorMsg: "missing name",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			renamedTag, err := renameTag(tt.tag, tt.newName)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)

				return
			}

			assert.Equal(t, tt.wantName, renamedTag.Name)
		})
	}
}

func TestMoveTag(t *testing.T) {
	t.Parallel()

	// TODO: fix this
	secondTestUUID := uuid.New()

	tests := []struct {
		testName  string
		tag       Tag
		tags      []Tag
		ancestors Path
		wantErr   bool
		errorMsg  string
	}{
		{
			testName:  "ok",
			tag:       newTestTag(),
			tags:      []Tag{newTestTag(), newTestTag()},
			ancestors: Path{TagID(uuid.New()), TagID(uuid.New())},
			wantErr:   false,
		},
		{
			testName: "fail tag is not leaf",
			tag: Tag{
				ID: TagID(secondTestUUID), ClassificationID: ClassificationID("fake-classification"),
				Name: map[string]string{"it": "B"}, Status: Ready,
				Ancestors: Path{TagID(uuid.New())},
			},
			tags: []Tag{{
				ID: TagID(uuid.New()), ClassificationID: ClassificationID("fake-classification"),
				Name: map[string]string{"it": "B"}, Status: Ready,
				Ancestors: Path{TagID(secondTestUUID)},
			},
				newTestTag()},
			ancestors: Path{TagID(uuid.New()), TagID(uuid.New())},
			wantErr:   true, errorMsg: "can't move tag, is not leaf",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			movedTag, err := moveTag(tt.tag, tt.ancestors, ChildrenFinder{tt.tags})
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)

				return
			}

			assert.Equal(t, tt.ancestors, movedTag.GetAncestors())
		})
	}
}
