package indefart

import (
	"testing"
)

func TestFindIndefiniteArticle(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "word <empty>",
			args: args{
				word: "",
			},
			want: "a",
		},
		{
			name: "word unanticipated result",
			args: args{
				word: "unanticipated result",
			},
			want: "an",
		},
		{
			name: "word unanimous vote",
			args: args{
				word: "unanimous vote",
			},
			want: "a",
		},
		{
			name: "word honest decision",
			args: args{
				word: "honest decision",
			},
			want: "an",
		},
		{
			name: "word honeysuckle shrub",
			args: args{
				word: "honeysuckle shrub",
			},
			want: "a",
		},
		{
			name: "word 0800 number",
			args: args{
				word: "0800 number",
			},
			want: "an",
		},
		{
			name: "word xmas tree",
			args: args{
				word: "xmas tree",
			},
			want: "a",
		},
		{
			name: "word unidirectional beam",
			args: args{
				word: "unidirectional beam",
			},
			want: "a",
		},
		{
			name: "word unidiomatic phrase",
			args: args{
				word: "unidiomatic phrase",
			},
			want: "an",
		},
		{
			name: "word NASA scientist",
			args: args{
				word: "NASA scientist",
			},
			want: "a",
		},
		{
			name: "word NSA analyst",
			args: args{
				word: "NSA analyst",
			},
			want: "an",
		},
		{
			name: "word FIAT car",
			args: args{
				word: "FIAT car",
			},
			want: "a",
		},
		{
			name: "word FAA policy",
			args: args{
				word: "FAA policy",
			},
			want: "an",
		},
		{
			name: `word with all ignored in prefix runes "‘’“”$\'-(0800 number`,
			args: args{
				word: `"‘’“”$\'-(0800 number`,
			},
			want: "an",
		},
		{
			name: "word Embedded Software Engineer",
			args: args{
				word: "Embedded Software Engineer",
			},
			want: "an",
		},
		{
			name: "word Staff Simulation Engineer",
			args: args{
				word: "Staff Simulation Engineer",
			},
			want: "a",
		},
		{
			name: "word Senior Infrastructure Engineer",
			args: args{
				word: "Senior Infrastructure Engineer",
			},
			want: "a",
		},
		{
			name: "word Senior Go Engineer",
			args: args{
				word: "Senior Go Engineer",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.word); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
