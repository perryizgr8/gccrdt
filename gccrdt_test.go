package gccrdt_test

import (
	gccrdt "gc-crdt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	node1 = "node-1"
	node2 = "node-2"
)

func TestGetCountMap_Empty(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	c := gc.GetCountMap()
	assert.Empty(t, c)
}

func TestGetCountMap_One(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	gc.Increment(node1)
	c := gc.GetCountMap()
	assert.Equal(t, 1, c[node1])
}

func TestGetCountMap_OneWithTwoNodes(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	gc.Increment(node1)
	gc.Increment(node2)
	c := gc.GetCountMap()
	assert.Equal(t, 1, c[node1])
	assert.Equal(t, 1, c[node2])
}

func TestGetCountMap_HundredWithTwoNodes(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	for range 100 {
		gc.Increment(node1)
		gc.Increment(node2)
	}
	c := gc.GetCountMap()
	assert.Equal(t, 100, c[node1])
	assert.Equal(t, 100, c[node2])
}

func TestValue_Zero(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	assert.Equal(t, 0, gc.Value())
}

func TestValue_One(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	gc.Increment(node1)
	assert.Equal(t, 1, gc.Value())
}

func TestValue_Hundred(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	for range 100 {
		gc.Increment(node1)
	}
	assert.Equal(t, 100, gc.Value())
}

func TestValue_OneWithTwoNodes(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	gc.Increment(node1)
	gc.Increment(node2)
	assert.Equal(t, 2, gc.Value())
}

func TestValue_HundredWithTwoNodes(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	for range 100 {
		gc.Increment(node1)
		gc.Increment(node2)
	}
	assert.Equal(t, 200, gc.Value())
}

func TestMerge(t *testing.T) {
	gc := gccrdt.NewGccrdt()
	for range 100 {
		gc.Increment(node1)
	}
	for range 50 {
		gc.Increment(node2)
	}
	others := [1000]map[string]int{}
	for i := range others {
		others[i] = make(map[string]int)
	}
	node1Max := 100
	node2Max := 50
	for _, other := range others {
		other[node1] = rand.IntN(100)
		node1Max = max(node1Max, other[node1])
		other[node2] = rand.IntN(100)
		node2Max = max(node2Max, other[node2])
	}
	gc.Merge(others[:]...)
	assert.Equal(t, node1Max+node2Max, gc.Value())
}
