package tag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName   string
		insert     TagInsert
		createdTag Tag
		isRoot     bool
		wantErr    bool
		errorMsg   string
	}{
		{
			testName: "ok is root",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
				Name:             LocalizedValue{"it": "B"},
				Parent:           newTagID(),
			},
			isRoot:  true,
			wantErr: false,
		},
		{
			testName: "ok is not root",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
				Name:             LocalizedValue{"it": "B"},
				Parent:           newTagID(),
				Ancestors:        Path{newTagID(), newTagID()},
			},
			isRoot:  false,
			wantErr: false,
		},
		{
			testName: "should fail if classification id is missing",
			insert: TagInsert{
				Name: LocalizedValue{"it": "B"},
			},
			wantErr:  true,
			errorMsg: "missing classification id",
		},
		{
			testName: "should fail if name is missing",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
			},
			wantErr:  true,
			errorMsg: "missing name",
		},
	}

	for _, tt := range tests {
		tt := tt // see https://github.com/golang/go/discussions/56010

		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			res, err := NewTag(tt.insert)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)

				return
			}

			assert.Equal(t, Ready, res.Status)
			assert.Equal(t, tt.insert.Name, res.Name)
			assert.Equal(t, tt.insert.ClassificationID, res.ClassificationID)
			assert.Equal(t, tt.isRoot, res.IsRoot())
		})
	}
}

func TestTag_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		tag  Tag
		want Tag
	}{
		{
			name: "clone successful",
			tag:  newTestTag(),
		},
	}
	for _, tt := range tests {
		tt := tt // see https://github.com/golang/go/discussions/56010
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cloned := tt.tag.Clone()

			assert.Equal(t, tt.tag, cloned)
		})
	}
}
