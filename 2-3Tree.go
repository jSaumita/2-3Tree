//Arbol 2-3 en GO
// Noviembre 2015 , Javier Sauma
//Basado en  2_3_tree/ttTree.cpp /ajaxon

package main

import (
	"fmt"
	"sort"
)

type nodo23 struct {
	left, mid, right, parent *nodo23
	smallData, bigData       int
	dataFull                 bool
}
type T23 struct {
	root *nodo23
}

func newT23() *T23 {
	return new(T23)
}
func (t *T23) Insert(x int) bool {
	fmt.Println("insertar", x)
	if t.root == nil {
		t.root = &nodo23{smallData: x}
		return true
	}
	if t.Search(x) {
		return false
	}
	return t.searchInsert(t.root, x)
}
func (t *T23) searchInsert(nodo *nodo23, x int) bool {
	if nodo.isLeaf() {
		var aux, aux2 *nodo23
		if nodo.isThreeTree() {
			t.insert3Tree(nodo, aux2, aux, x)
		} else {
			nodo.insert2Tree(x, aux2, aux)
		}
	} else {
		if x < nodo.smallData {
			t.searchInsert(nodo.left, x)
		} else if nodo.isThreeTree() && x > nodo.bigData {
			t.searchInsert(nodo.right, x)
		} else {
			t.searchInsert(nodo.mid, x)
		}
	}
	return true
}
func (t *T23) Search(x int) bool {
	tmp := t.root
	for tmp != nil {
		if tmp.smallData == x || tmp.bigData == x {
			return true
		}
		if x < tmp.smallData {
			tmp = tmp.left
		} else if tmp.isThreeTree() && x > tmp.bigData {
			tmp = tmp.right
		} else {
			tmp = tmp.mid
		}
	}
	return false
}
func (n *nodo23) insert2Tree(x int, leftC, rightC *nodo23) {
	if n.smallData > x {
		n.bigData = n.smallData
		n.smallData = x
		n.left = leftC
		n.right = n.mid
		n.mid = rightC
	} else {
		n.bigData = x
		n.mid = leftC
		n.right = rightC
	}
	n.dataFull = true
}
func (n *nodo23) isThreeTree() bool {
	return n.bigData != 0
}
func (n *nodo23) isLeaf() bool {
	return n.left == nil && n.mid == nil && n.right == nil
}
func (t *T23) insert3Tree(nodo, newLeft, newRight *nodo23, x int) {
	p := nodo
	p = nil
	nodoLC := nodo.left
	nodoMC := nodo.mid
	nodoRC := nodo.right
	cont := 0
	if x < nodo.smallData {
		cont = 1
	} else if x > nodo.bigData {
		cont = 2
	} else {
		cont = 3
	}
	data := []int{nodo.smallData, nodo.bigData, x}
	sort.Ints(data)
	//caso 1 sin padres o nodo raiz
	if nodo.parent == nil {
		p = &nodo23{}
	} else {
		p = nodo.parent
	}
	leaf := nodo.isLeaf()
	mid := data[1]
	n1 := &nodo23{
		smallData: data[0],
		parent:    p,
	}
	nodo.nodoEdit(data[2], p)
	n2 := nodo
	if !leaf {
		if cont == 1 {
			newLeft.parent = n1
			newRight.parent = n1
			n1.left = newLeft
			n1.mid = newRight
			n2.mid = nodoRC
			n2.left = nodoMC
			nodoMC.parent = n2
			nodoRC.parent = n2
		} else if cont == 2 {
			newLeft.parent = n2
			newRight.parent = n2
			n2.mid = newRight
			n2.left = newLeft
			n1.left = nodoLC
			n1.mid = nodoMC
			nodoLC.parent = n1
			nodoMC.parent = n1
		} else {
			n1.left = nodoLC
			nodoLC.parent = n1
			n1.mid = newLeft
			newLeft.parent = n1
			n2.mid = nodoRC
			nodoRC.parent = n2
			n2.left = newRight
			newRight.parent = n2
		}
	}
	if p.left == nil {
		p.smallData = mid
		p.left = n1
		p.mid = n2
		t.root = p
	} else if !p.isThreeTree() {
		p.insert2Tree(mid, n1, n2)
	} else {
		t.insert3Tree(p, n1, n2, mid)
	}
}
func (n *nodo23) nodoEdit(data int, p *nodo23) {
	n.smallData = data
	n.bigData = 0
	n.parent = p
	n.left = nil
	n.right = nil
	n.mid = nil
}
func (t *T23) tPrint() {
	if t.root == nil {
		fmt.Println("vacio")
		return
	}
	t.inorderTraverse(t.root)
	fmt.Println()

}
func (t *T23) inorderTraverse(nodo *nodo23) {
	if nodo.isLeaf() {
		if nodo.isThreeTree() {
			fmt.Print("(", nodo.smallData, ",", nodo.bigData, ")")
		} else {
			fmt.Print("(", nodo.smallData, ")")
		}
	} else if !nodo.isThreeTree() {
		fmt.Print("(")
		t.inorderTraverse(nodo.left)
		fmt.Print(",", nodo.smallData, ",")
		t.inorderTraverse(nodo.mid)
		fmt.Print(")")
	} else {
		fmt.Print("(")
		t.inorderTraverse(nodo.left)
		fmt.Print(",", nodo.smallData, ",")
		t.inorderTraverse(nodo.mid)
		fmt.Print(",", nodo.bigData, ",")
		t.inorderTraverse(nodo.right)
		fmt.Print(")")
	}
}
func main() {
	mytree := newT23()
	// mytree.t23Insert(37)
	// mytree.t23Insert(50)
	// mytree.t23Insert(30)
	// mytree.t23Insert(39)
	// mytree.t23Insert(70)
	// mytree.t23Insert(90)
	// mytree.t23Insert(10)
	// mytree.t23Insert(36)
	// mytree.t23Insert(20)
	// mytree.t23Insert(38)
	// mytree.t23Insert(40)
	// mytree.t23Insert(60)
	// mytree.t23Insert(80)
	// mytree.t23Insert(100)
	// mytree.t23Insert(35)
	// mytree.t23Insert(34)
	// mytree.t23Insert(33)
	// mytree.t23Insert(32)
	for i := 1; i < 11; i++ {
		mytree.t23Insert(i * 10)
	}
	mytree.tPrint()

}
