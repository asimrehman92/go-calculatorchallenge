package stores

import "testing"

func TestDataMigration(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "dataMigrationFunc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DataMigration()
		})
	}
}
