package defangingAnIPAddress

import "testing"

func Test_defangIPAddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "255.255.255.0 => 255[.]255[.]255[.]0",
			args: args{address: "255.255.255.0"},
			want: "255[.]255[.]255[.]0",
		},
		{
			name: "255.100.50.0 => 255[.]100[.]50[.]0",
			args: args{address: "255.100.50.0"},
			want: "255[.]100[.]50[.]0",
		},
		{
			name: "1.1.1.1 => 1[.]1[.]1[.]1",
			args: args{address: "1.1.1.1"},
			want: "1[.]1[.]1[.]1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPAddr(tt.args.address); got != tt.want {
				t.Errorf("defangIPAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
