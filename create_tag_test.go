package tag

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		insert     TagInsert
		createdTag Tag
		wantErr    bool
		errorMsg   string
	}{
		{
			name: "ok",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
				Name:             map[string]string{"it": "B"},
				Parent:           TagID(uuid.New()),
			},
			wantErr: false,
		},
		{
			name: "should fail if classification id is missing",
			insert: TagInsert{
				Name: map[string]string{"it": "B"},
			},
			wantErr:  true,
			errorMsg: "missing classification id",
		},
		{
			name: "should fail if name is missing",
			insert: TagInsert{
				ClassificationID: ClassificationID("fake-classification"),
			},
			wantErr:  true,
			errorMsg: "missing name",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res, err := createTag(tt.insert)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)

				return
			}

			assert.Equal(t, res.Status, Ready)
			assert.Equal(t, res.Parent, tt.insert.Parent)
			assert.Equal(t, res.Name, tt.insert.Name)
			assert.Equal(t, res.ClassificationID, tt.insert.ClassificationID)
		})
	}
}
