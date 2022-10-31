package stores

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

var u = &Clients{
	Model:     gorm.Model{},
	Name:      "Momo",
	Equation:  "100 + 900 = 1000",
	Timestamp: time.Now(),
}

func init() {
	DataMigration()
}

func TestCreateClient(t *testing.T) {
	type args struct {
		uname, eq string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "createClientFunc",
			args: args{u.Name, u.Equation},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateClient(tt.args.uname, tt.args.eq)
		})
	}
}

func TestUpdateClient(t *testing.T) {
	type args struct {
		uname, eq string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "updateClientFunc",
			args: args{u.Name, u.Equation},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateClient(tt.args.uname, tt.args.eq)
		})
	}
}
