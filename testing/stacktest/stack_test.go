package main

import (
	"fmt"
	"testing"
)

func Test_stack_push(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    *stack
		args args
	}{
		// s=> s{}
		// s.push(3)=> s{3}
		// s.push(3)->push(4)->push(6)->push(9)-> => s{3,3,4,6,9}
		{
			name: "Case 1: pushing an element onto an empty stack",
			s:    &stack{},
			args: args{n: 3},
		},
		{
			name: "Case 2: pushing an element onto non empty stack",
			s:    &stack{3},
			args: args{n: 4},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.push(tt.args.n)
			if len(*tt.s) == 0 {
				t.Errorf("stack push failed!!")
			}
			if i == 1 && len(*tt.s) == 2 && (*tt.s)[1] != 4 {
				t.Errorf("stack push failed!!")
			}
		})
	}
}

// func Test_stack_pop(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		s    *stack
// 	}{
// 		{
// 			name: "Case 1: poping an element onto an empty stack",
// 			s:    &stack{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.s.pop()

// 		})
// 	}
// }

// func Test_stack_pop(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		s       *stack
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Case 1: poping an element from an empty stack",
// 			s:       &stack{},
// 			wantErr: true,
// 		},
// 		{
// 			name:    "Case 2: poping an element from a non empty stack",
// 			s:       &stack{3, 4},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.s.pop()
// 			flag := (err != nil)
// 			fmt.Printf("***Error:%v\n", err)
// 			fmt.Printf("**flag:%v\n", flag)
// 			if flag != tt.wantErr {
// 				t.Errorf("stack.pop() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func Test_stack_pop(t *testing.T) {
	tests := []struct {
		name      string
		s         *stack
		wantPoped int
		wantErr   bool
	}{
		{
			name:      "Case 1: poping an element from an empty stack",
			s:         &stack{},
			wantPoped: 0,
			wantErr:   true,
		},
		{
			name:      "Case 2: poping an element from a non empty stack",
			s:         &stack{3, 4},
			wantPoped: 4,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPoped, err := tt.s.pop()
			fmt.Printf("Poped stack:%v\n", tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("stack.pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPoped != tt.wantPoped {
				t.Errorf("stack.pop() = %v, want %v", gotPoped, tt.wantPoped)
			}
		})
	}
}
