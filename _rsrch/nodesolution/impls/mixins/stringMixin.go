package mixins

import (
	ipld "github.com/ipld/go-ipld-prime/_rsrch/nodesolution"
)

// String can be embedded in a struct to provide all the methods that
// have fixed output for any string-kinded nodes.
// (Mostly this includes all the methods which simply return ErrWrongKind.)
// Other methods will still need to be implemented to finish conforming to Node.
//
// To conserve memory and get a TypeName in errors without embedding,
// write methods on your type with a body that simply initializes this struct
// and immediately uses the relevant method;
// this is more verbose in source, but compiles to a tighter result:
// in memory, there's no embed; and in runtime, the calls will be inlined
// and thus have no cost in execution time.
type String struct {
	TypeName string
}

func (String) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_String
}
func (x String) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (x String) Lookup(key ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (x String) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}
}
func (x String) LookupSegment(ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_String}
}
func (String) MapIterator() ipld.MapIterator {
	return nil
}
func (String) ListIterator() ipld.ListIterator {
	return nil
}
func (String) Length() int {
	return -1
}
func (String) IsUndefined() bool {
	return false
}
func (String) IsNull() bool {
	return false
}
func (x String) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_String}
}
func (x String) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_String}
}
func (x String) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_String}
}
func (x String) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_String}
}
func (x String) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: x.TypeName, MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_String}
}
