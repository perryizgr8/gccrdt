package gccrdt

type Gccrdt struct {
	count map[string]int
}

func NewGccrdt() Gccrdt {
	return Gccrdt{
		count: make(map[string]int),
	}
}

func (gc Gccrdt) Increment(node string) {
	gc.count[node]++
}

func (gc Gccrdt) Merge(others ...map[string]int) {
	for _, count := range others {
		for node, theirVal := range count {
			ourVal := gc.count[node]
			gc.count[node] = max(ourVal, theirVal)
		}
	}
}

func (gc Gccrdt) Value() int {
	total := 0
	for _, c := range gc.count {
		total += c
	}
	return total
}

func (gc Gccrdt) GetCountMap() map[string]int {
	return gc.count
}
