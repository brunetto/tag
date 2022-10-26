package tag

import (
	"testing"

	"github.com/google/uuid"
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
				Parent:           TagID(uuid.New()),
			},
			isRoot:  true,
			wantErr: false,
		},
		{
			testName: "ok is not root",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
				Name:             LocalizedValue{"it": "B"},
				Parent:           TagID(uuid.New()),
				Ancestors:        Path{TagID(uuid.New()), TagID(uuid.New())},
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
		tt := tt

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
