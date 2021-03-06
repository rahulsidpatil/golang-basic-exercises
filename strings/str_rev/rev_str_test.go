package main

import "testing"

func Test_revstr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "string_reverse_test",
			args:       args{str: "pqrs"},
			wantResult: "srqp",
		},
		{
			name:       "string_reverse_test1",
			args:       args{str: "pqrst"},
			wantResult: "srqp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := revstr(tt.args.str); gotResult != tt.wantResult {
				t.Errorf("revstr() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// func Test_revstr(t *testing.T) {
// 	type args struct {
// 		str string
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantResult string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotResult := revstr(tt.args.str); gotResult != tt.wantResult {
// 				t.Errorf("revstr() = %v, want %v", gotResult, tt.wantResult)
// 			}
// 		})
// 	}
// }
