package subset

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/eminano/greenmask/internal/db/postgres/entries"
)

func TestComponent_findCycles(t *testing.T) {
	c := &Component{
		cyclesIdents: make(map[string]struct{}),
		componentGraph: map[int][]*Edge{
			1: {
				{
					id:  1,
					idx: 2,
					from: &TableLink{
						idx: 1,
					},
					to: &TableLink{
						idx: 2,
					},
				},
			},
			2: {
				{
					id:  2,
					idx: 3,
					from: &TableLink{
						idx: 2,
					},
					to: &TableLink{
						idx: 3,
					},
				},
			},
			3: {
				{
					id:  3,
					idx: 1,
					from: &TableLink{
						idx: 3,
					},
					to: &TableLink{
						idx: 1,
					},
				},
				{
					id:  4,
					idx: 1,
					from: &TableLink{
						idx: 3,
					},
					to: &TableLink{
						idx: 1,
					},
				},
				{
					id:  5,
					idx: 4,
					from: &TableLink{
						idx: 3,
					},
					to: &TableLink{
						idx: 4,
					},
				},
			},
			4: {
				{
					id:  6,
					idx: 3,
					from: &TableLink{
						idx: 4,
					},
					to: &TableLink{
						idx: 3,
					},
				},
				{
					id:  7,
					idx: 1,
					from: &TableLink{
						idx: 4,
					},
					to: &TableLink{
						idx: 1,
					},
				},
			},
		},
		tables: map[int]*entries.Table{},
	}

	c.findCycles()
	require.Len(t, c.cycles, 4)
}

func TestComponent_findCycles_pt2(t *testing.T) {
	c := &Component{
		componentGraph: map[int][]*Edge{
			1: {
				{
					id:  1,
					idx: 2,
					from: &TableLink{
						idx: 1,
					},
					to: &TableLink{
						idx: 2,
					},
				},
			},
			2: {
				{
					id:  2,
					idx: 1,
					from: &TableLink{
						idx: 2,
					},
					to: &TableLink{
						idx: 1,
					},
				},
				{
					id:  3,
					idx: 1,
					from: &TableLink{
						idx: 2,
					},
					to: &TableLink{
						idx: 1,
					},
				},
			},
		},
		tables:       map[int]*entries.Table{},
		cyclesIdents: make(map[string]struct{}),
	}

	c.findCycles()
	require.Len(t, c.cycles, 2)
}

func BenchmarkComponent_findCycles(b *testing.B) {
	c := &Component{
		cyclesIdents: make(map[string]struct{}),
		componentGraph: map[int][]*Edge{
			1: {
				{
					id:  1,
					idx: 2,
					from: &TableLink{
						idx: 1,
					},
					to: &TableLink{
						idx: 2,
					},
				},
			},
			2: {
				{
					id:  2,
					idx: 1,
					from: &TableLink{
						idx: 2,
					},
					to: &TableLink{
						idx: 1,
					},
				},
				{
					id:  3,
					idx: 1,
					from: &TableLink{
						idx: 2,
					},
					to: &TableLink{
						idx: 1,
					},
				},
			},
		},
		tables: map[int]*entries.Table{},
	}

	// Reset the timer to exclude the setup time from the benchmark
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.findCycles()
	}
}
