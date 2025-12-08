package utils

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return
	}
	if d.size[px] < d.size[py] {
		px, py = py, px // keeps tree shallow
	}
	d.parent[py] = px
	d.size[px] += d.size[py]
}

func (d *DSU) Size(x int) int {
	return d.size[d.Find(x)]
}

func (d *DSU) Roots() []int {
	roots := make([]int, 0)
	for i, p := range d.parent {
		if i == p {
			roots = append(roots, i)
		}
	}
	return roots
}

func (d *DSU) RootSizes() []int {
	roots := d.Roots()
	sizes := make([]int, 0)
	for _, r := range roots {
		sizes = append(sizes, d.size[r])
	}
	return sizes
}
