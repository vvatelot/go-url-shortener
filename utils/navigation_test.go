package utils

import (
	"testing"
)

func TestGetLastPage(t *testing.T) {
	type args struct {
		countAll int64
		size     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test last page of 10 elements with page size 10",
			args: args{
				countAll: 10,
				size:     10,
			},
			want: 1,
		},
		{
			name: "Test last page of 10 elements with page size 5",
			args: args{
				countAll: 10,
				size:     5,
			},
			want: 2,
		},
		{
			name: "Test last page of 0 elements with page size 10",
			args: args{
				countAll: 0,
				size:     10,
			},
			want: 0,
		},
		{
			name: "Test last page of 100 elements with page size 10",
			args: args{
				countAll: 100,
				size:     10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastPage(tt.args.countAll, tt.args.size); got != tt.want {
				t.Errorf("GetLastPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPreviousPage(t *testing.T) {
	type args struct {
		currentPage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test previous page of 1",
			args: args{
				currentPage: 1,
			},
			want: 0,
		},
		{
			name: "Test previous page of 2",
			args: args{
				currentPage: 2,
			},
			want: 1,
		},
		{
			name: "Test previous page of 10",
			args: args{
				currentPage: 10,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPreviousPage(tt.args.currentPage); got != tt.want {
				t.Errorf("GetPreviousPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextPage(t *testing.T) {
	type args struct {
		currentPage int
		lastPage    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test next page of page 1 of 2",
			args: args{
				currentPage: 1,
				lastPage:    2,
			},
			want: 2,
		},
		{
			name: "Test next page of page 2 of 2",
			args: args{
				currentPage: 2,
				lastPage:    2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNextPage(tt.args.currentPage, tt.args.lastPage); got != tt.want {
				t.Errorf("GetNextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
