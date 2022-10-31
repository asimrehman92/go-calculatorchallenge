package cal_telnet

import (
	"net"
	"testing"
)

var connection net.Conn

func TestHandleConnection(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "handleConnectFunc",
			args: args{connection},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleConnection(tt.args.c)
		})
	}
}
