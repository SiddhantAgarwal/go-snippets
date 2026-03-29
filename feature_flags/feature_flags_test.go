package feature_flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeatureSet_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		want  FeatureSet
		toAdd FeatureFlag
	}{
		{
			name:  "only FeatureA",
			toAdd: FeatureA,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0000_0001,
			},
		},
		{
			name:  "only FeatureB",
			toAdd: FeatureB,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0000_0010,
			},
		},
		{
			name:  "only FeatureC",
			toAdd: FeatureC,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0000_0100,
			},
		},
		{
			name:  "only FeatureD",
			toAdd: FeatureD,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0000_1000,
			},
		},
		{
			name:  "only FeatureE",
			toAdd: FeatureE,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0001_0000,
			},
		},
		{
			name:  "only FeatureF",
			toAdd: FeatureF,
			want: FeatureSet{
				mask: 0b_0000_0000_0000_0000_0000_0000_0010_0000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fs := FeatureSet{}
			fs.Add(tt.toAdd)

			assert.Equal(t, tt.want, fs)
		})
	}
}

func TestFeatureSet_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		setup func() FeatureSet
		get   FeatureFlag
	}{
		{
			name: "only FeatureA",
			setup: func() FeatureSet {
				fs := FeatureSet{}
				fs.Add(FeatureA)

				return fs
			},
			get: FeatureA,
		},
		{
			name: "both FeatureA & FeatureB",
			setup: func() FeatureSet {
				fs := FeatureSet{}
				fs.Add(FeatureA)
				fs.Add(FeatureB)

				return fs
			},
			get: FeatureA | FeatureB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fs := tt.setup()
			assert.True(t, fs.IsEnabled(tt.get))
		})
	}
}

func TestFeatureSet_Remove(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func() FeatureSet
		remove FeatureFlag
	}{
		{
			name: "only FeatureA",
			setup: func() FeatureSet {
				fs := FeatureSet{}
				fs.Add(FeatureA)
				fs.Add(FeatureB)
				fs.Add(FeatureC)
				fs.Add(FeatureD)
				fs.Add(FeatureE)
				fs.Add(FeatureE)

				return fs
			},
			remove: FeatureA,
		},
		{
			name: "FeatureA & FeatureB",
			setup: func() FeatureSet {
				fs := FeatureSet{}
				fs.Add(FeatureA)
				fs.Add(FeatureB)
				fs.Add(FeatureC)
				fs.Add(FeatureD)
				fs.Add(FeatureE)
				fs.Add(FeatureE)

				return fs
			},
			remove: FeatureA | FeatureB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fs := tt.setup()
			fs.Remove(tt.remove)
			assert.False(t, fs.IsEnabled(tt.remove))
		})
	}
}
