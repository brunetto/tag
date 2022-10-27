package tag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestTag() Tag {
	return Tag{
		ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
		Name: LocalizedValue{"it": "B"}, Status: Ready,
	}
}

func newFixedTestTag() Tag {
	return Tag{
		ID: TagID(NewFakeFixedUUID()), ClassificationID: ClassificationID("fake-classification"),
		Name: LocalizedValue{"it": "B"}, Status: Ready,
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
			newName:  LocalizedValue{"it": "C"},
			tag:      newTestTag(),
			wantName: LocalizedValue{"it": "C"},
			wantErr:  false,
		},
		{
			testName: "fail if new name is empty",
			newName:  nil,
			tag:      newTestTag(),
			wantName: LocalizedValue{"it": "C"},
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
			ancestors: Path{newTagID(), newTagID()},
			wantErr:   false,
		},
		{
			testName: "fail tag is not leaf",
			tag: Tag{
				ID: TagID(NewFakeFixedUUID()), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{newTagID()},
			},
			tags: []Tag{{
				ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{TagID(NewFakeFixedUUID())},
			},
				newTestTag()},
			ancestors: Path{newTagID(), newTagID()},
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

func TestMakeAliasOf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		tagA     Tag
		tagB     Tag
		tags     []Tag
		wantErr  bool
		errorMsg string
	}{
		{
			testName: "ok",
			tagA:     newTestTag(),
			tagB:     newTestTag(),
			tags:     []Tag{newTestTag(), newTestTag()},
			wantErr:  false,
		},
		{
			testName: "fail tagA is not leaf",
			tagA: Tag{
				ID: TagID(NewFakeFixedUUID()), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{newTagID()},
			},
			tagB: Tag{
				ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{newTagID()},
			},
			tags: []Tag{{
				ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{TagID(NewFakeFixedUUID())},
			},
				newTestTag()},
			wantErr:  true,
			errorMsg: fmt.Sprintf("can't alias tag, source tag %v is not leaf", TagID(NewFakeFixedUUID())),
		},
		{
			testName: "fail tagB is not leaf",
			tagA: Tag{
				ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{newTagID()},
			},
			tagB: Tag{
				ID: TagID(NewFakeFixedUUID()), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{newTagID()},
			},
			tags: []Tag{{
				ID: newTagID(), ClassificationID: ClassificationID("fake-classification"),
				Name: LocalizedValue{"it": "B"}, Status: Ready,
				Ancestors: Path{TagID(NewFakeFixedUUID())},
			},
				newTestTag()},
			wantErr:  true,
			errorMsg: fmt.Sprintf("can't alias tag, destination tag %v is not leaf", TagID(NewFakeFixedUUID())),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			aliasedTag, err := makeAliasOf(tt.tagA, tt.tagB, ChildrenFinder{tt.tags})
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Equal(t, tt.tagA, aliasedTag)
				return
			}

			assert.True(t, samePayload(tt.tagB, aliasedTag))
		})
	}
}

func Test_samePayload(t *testing.T) {
	tests := []struct {
		name string
		tagA Tag
		tagB Tag
		want bool
	}{
		{
			name: "same tags",
			tagA: newTestTag(),
			tagB: newTestTag(),
			want: true,
		},
		{
			name: "different tags",
			tagA: Tag{},
			tagB: newTestTag(),
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			same := samePayload(tt.tagA, tt.tagB)

			assert.Equal(t, tt.want, same)
		})
	}
}
