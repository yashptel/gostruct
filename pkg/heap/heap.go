package heap


type heap []int

func NewHeap(arr *[]int) (*heap){
	a := &heap{}
	*a = heap(*arr)
	n := a.Len()
	for i := n/2 - 1; i >= 0; i-- {
		a.down(i, n)
	}
	return a
}

func (h heap) Len() int           { return len(h) }
func (h heap) Less(i, j int) bool { return h[i] < h[j] }
func (h heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heap) Push(x any) {
	*h = append(*h, x.(int))
	h.up(h.Len()-1)
}

func (h *heap) Pop() any {
	no := h.Len() - 1
	h.Swap(0, no)
	h.down(0, no)

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func  (h *heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
