package gccrdt_test

import (
	gccrdt "gc-crdt"
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
