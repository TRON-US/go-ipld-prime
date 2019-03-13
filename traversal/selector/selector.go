package selector

import ipld "github.com/ipld/go-ipld-prime"

type Selector interface {
	Explore(ipld.Node) (ipld.KeyIterator, Selector)
	Decide(ipld.Node) bool
}

func ReifySelector(cidRootedSelector ipld.Node) (Selector, error) {
	return nil, nil
}
