package react

import (

	"github.com/google/uuid"
)

type Cel struct {
	name uuid.UUID
	rtr *RTR
}

func (c Cel) Value() int {
	return c.rtr.cellVals[c.name]
}

type InpCell struct {
	Cel
}

// RTR is Reactor
type RTR struct {
	cellDeps map[uuid.UUID][]*ComCell
	cellVals map[uuid.UUID]int
}


type ComCell struct {
	Cel
	compute func(int) int
	compute2 func(int, int) int
	parentName []uuid.UUID
	callbacks map[uuid.UUID]func(int)
}

func (ic *InpCell) SetValue(input int) {
	rtr := ic.rtr
	if rtr.cellVals[ic.name] == input {
		return
	}
	rtr.cellVals[ic.name] = input
	// propagate
	if _, ok := rtr.cellDeps[ic.name]; ok {
		propagate(rtr.cellDeps[ic.name], input)
	} 
}

type Cancelr struct {
	name uuid.UUID
	cell *ComCell
}

func (clr Cancelr) Cancel() {
	delete(clr.cell.callbacks, clr.name)
}

// propagate will be responsible for propagating changes to dependents and invoke callbacks
func propagate(comps []*ComCell, input int) {
	q := comps
	visited := make(map[uuid.UUID]bool)
	for ; len(q) > 0; {
		nq := make([]*ComCell, 0)
		for _, comp := range(q) {
			changed := false
			if len(comp.parentName) == 2 {
				one, another := comp.rtr.cellVals[comp.parentName[0]], comp.rtr.cellVals[comp.parentName[1]]
				newVal := comp.compute2(one, another)
				if newVal != comp.rtr.cellVals[comp.name] {
					comp.rtr.cellVals[comp.name] = newVal
					changed = true
				}
			}
			if len(comp.parentName) == 1 {
				newVal := comp.compute(comp.rtr.cellVals[comp.parentName[0]])
				if newVal != comp.rtr.cellVals[comp.name] {
					comp.rtr.cellVals[comp.name] = comp.compute(comp.rtr.cellVals[comp.parentName[0]])
					changed = true
				}
			}

			if !changed {
				continue
			}
			if _, ok := visited[comp.name]; ok && visited[comp.name] {
				continue
			}
			visited[comp.name] = true

			// callback
			for _, callback := range(comp.callbacks) {
				callback(comp.rtr.cellVals[comp.name])
			}
			if _, ok := comp.rtr.cellDeps[comp.name]; ok {
				nq = append(nq, comp.rtr.cellDeps[comp.name]...)
			}
		}
		q = nq
	}
}

func (cc *ComCell) AddCallback(f func(int)) Canceler {
	key := uuid.New()
	cc.callbacks[key] = f
	clr := Cancelr{ name: key, cell: cc }
	return clr
}

// CreateInput will create input cells
func (r *RTR) CreateInput(input int) InputCell {
	key := uuid.New()
	c := Cel{
		name: key,
		rtr: r,
	}
	inp := InpCell{ Cel: c }
	r.cellDeps[key] = make([]*ComCell, 0)
	r.cellVals[key] = input
	return &inp
}

// CreateCompute1 will create compute cells from cell with compute function
func (r *RTR) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	var name uuid.UUID
	switch cell.(type) {
	case *InpCell:
		c := cell.(*InpCell)
		name = c.name
	case *ComCell:
		c := cell.(*ComCell)
		name = c.name
	case *Cel:
		c := cell.(*Cel)
		name = c.name
	}
	comp := &ComCell{
		Cel: Cel{},
		compute: compute,
		parentName: []uuid.UUID{name},
		callbacks: make(map[uuid.UUID]func(int), 0),
	}
	comp.rtr = r
	comp.name = uuid.New()
	// register to reactor
	r.cellDeps[name] = append(r.cellDeps[name], comp)
	r.cellVals[comp.name] = compute(r.cellVals[name])
	return comp
}
// CreateCompute2 will create compute cells from two cells with compute function
func (r *RTR) CreateCompute2(cell, cell2 Cell, compute func(int, int) int) ComputeCell {
  var name uuid.UUID
	switch cell.(type) {
	case *InpCell:
		c := cell.(*InpCell)
		name = c.name
	case *ComCell:
		c := cell.(*ComCell)
		name = c.name
	case *Cel:
		c := cell.(*Cel)
		name = c.name
	}
	var name2 uuid.UUID
	switch cell2.(type) {
	case *InpCell:
		c2 := cell2.(*InpCell)
		name2 = c2.name
	case *ComCell:
		c2 := cell2.(*ComCell)
		name2 = c2.name
	case *Cel:
		c2 := cell2.(*Cel)
		name2 = c2.name
	}

	comp := &ComCell{
		Cel: Cel{},
		compute2: compute,
		parentName: []uuid.UUID{name, name2},
		callbacks: make(map[uuid.UUID]func(int), 0),
	}
	comp.name = uuid.New()
	comp.rtr = r
	// register to reactor
	r.cellDeps[name] = append(r.cellDeps[name], comp)
	r.cellDeps[name2] = append(r.cellDeps[name2], comp)
	r.cellVals[comp.name] = compute(r.cellVals[name], r.cellVals[name2])
	return comp
}

// New will create a new Reactor
func New() *RTR {
	return &RTR{
		cellDeps: make(map[uuid.UUID][]*ComCell),
		cellVals: make(map[uuid.UUID]int),
	}
}