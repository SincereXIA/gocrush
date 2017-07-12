package gocrush

import (
	"math/rand"
	"strconv"
	"testing"
)

// nolint
func BenchmarkCrushStraw(b *testing.B) {
	b.StopTimer()
	tree := makeBenchStrawTree()
	r := rand.New(rand.NewSource(544564))
	b.StartTimer()

	for x := 0; x < b.N; x++ {
		Select(tree, r.Int63(), 3, NODE, nil)
	}
	b.StopTimer()
}

// nolint
func BenchmarkCrushTree(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(544564))
	tree := makeBenchTreeTree()
	b.StartTimer()
	for x := 0; x < b.N; x++ {
		Select(tree, r.Int63(), 3, NODE, nil)

	}
	b.StopTimer()
}

// nolint
func makeBenchStrawTree() *TestingNode {
	var parent = new(TestingNode)
	parent.Id = "ROOT"
	parent.Type = ROOT
	parent.Weight = 0
	parent.Children = make([]Node, 5)
	for dc := 0; dc < 5; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = DATA_CENTER
		dcNode.Id = parent.Id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 50)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 50; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = RACK
			rkNode.Weight = 1
			rkNode.Id = dcNode.Id + ":Rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 5000)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 5000; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = NODE
				ndNode.Weight = 1
				ndNode.Id = rkNode.Id + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 2)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 2; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = DISK
					dskNode.Weight = 1
					dskNode.Id = ndNode.Id + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewStrawSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewStrawSelector(ndNode)
			}
			rkNode.Selector = NewStrawSelector(rkNode)
		}
		dcNode.Selector = NewStrawSelector(dcNode)
	}
	parent.Selector = NewStrawSelector(parent)
	return parent
}

// nolint
func makeBenchSimpleStrawTree() *TestingNode {
	var parent = new(TestingNode)
	parent.Id = "ROOT"
	parent.Type = ROOT
	parent.Weight = 0
	parent.Children = make([]Node, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = DATA_CENTER
		dcNode.Id = parent.Id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 1)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = RACK
			rkNode.Weight = 1
			rkNode.Id = dcNode.Id + ":Rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 3)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = NODE
				ndNode.Weight = 1
				ndNode.Id = rkNode.Id + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 1)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = DISK
					dskNode.Weight = 1
					dskNode.Id = ndNode.Id + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewStrawSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewStrawSelector(ndNode)
			}
			rkNode.Selector = NewStrawSelector(rkNode)
		}
		dcNode.Selector = NewStrawSelector(dcNode)
	}
	parent.Selector = NewStrawSelector(parent)
	return parent
}

// nolint
func makeBenchTreeTree() *TestingNode {
	var parent = new(TestingNode)
	parent.Id = "ROOT"
	parent.Type = ROOT
	parent.Weight = 0
	parent.Children = make([]Node, 5)
	for dc := 0; dc < 5; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = DATA_CENTER
		dcNode.Id = parent.Id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 50)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 50; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = RACK
			rkNode.Weight = 1
			rkNode.Id = dcNode.Id + ":Rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 5000)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 5000; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = NODE
				ndNode.Weight = 1
				ndNode.Id = rkNode.Id + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 2)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 2; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = DISK
					dskNode.Weight = 1
					dskNode.Id = ndNode.Id + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewTreeSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewTreeSelector(ndNode)
			}
			rkNode.Selector = NewTreeSelector(rkNode)
		}
		dcNode.Selector = NewTreeSelector(dcNode)
	}
	parent.Selector = NewTreeSelector(parent)
	return parent
}

// nolint
func makeBenchSimpleTreeTree() *TestingNode {
	var parent = new(TestingNode)
	parent.Id = "ROOT"
	parent.Type = ROOT
	parent.Weight = 0
	parent.Children = make([]Node, 1)
	for dc := 0; dc < 1; dc++ {
		var dcNode = new(TestingNode)
		dcNode.Parent = parent
		dcNode.Weight = 1
		dcNode.Type = DATA_CENTER
		dcNode.Id = parent.Id + ":DataCenter" + strconv.Itoa(dc)
		dcNode.Children = make([]Node, 1)

		parent.Children[dc] = dcNode

		for rk := 0; rk < 1; rk++ {
			var rkNode = new(TestingNode)
			rkNode.Parent = dcNode
			rkNode.Type = RACK
			rkNode.Weight = 1
			rkNode.Id = dcNode.Id + ":Rack" + strconv.Itoa(rk)
			rkNode.Children = make([]Node, 3)

			dcNode.Children[rk] = rkNode
			for nd := 0; nd < 3; nd++ {
				var ndNode = new(TestingNode)
				ndNode.Parent = rkNode
				ndNode.Type = NODE
				ndNode.Weight = 1
				ndNode.Id = rkNode.Id + ":Node" + strconv.Itoa(nd)
				ndNode.Children = make([]Node, 1)

				rkNode.Children[nd] = ndNode
				for dsk := 0; dsk < 1; dsk++ {
					var dskNode = new(TestingNode)
					dskNode.Parent = ndNode
					dskNode.Type = DISK
					dskNode.Weight = 1
					dskNode.Id = ndNode.Id + ":Disk" + strconv.Itoa(dsk)
					dskNode.Selector = NewTreeSelector(dskNode)
					ndNode.Children[dsk] = dskNode
				}
				ndNode.Selector = NewTreeSelector(ndNode)
			}
			rkNode.Selector = NewTreeSelector(rkNode)
		}
		dcNode.Selector = NewTreeSelector(dcNode)
	}
	parent.Selector = NewTreeSelector(parent)
	return parent
}
