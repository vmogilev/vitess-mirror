/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by ASTHelperGen. DO NOT EDIT.

package integration

import (
	vtrpc "vitess.io/vitess/go/vt/proto/vtrpc"
	vterrors "vitess.io/vitess/go/vt/vterrors"
)

// EqualsAST does deep equals between the two objects.
func EqualsAST(inA, inB AST) bool {
	if inA == nil && inB == nil {
		return true
	}
	if inA == nil || inB == nil {
		return false
	}
	switch a := inA.(type) {
	case BasicType:
		b, ok := inB.(BasicType)
		if !ok {
			return false
		}
		return a == b
	case Bytes:
		b, ok := inB.(Bytes)
		if !ok {
			return false
		}
		return EqualsBytes(a, b)
	case InterfaceContainer:
		b, ok := inB.(InterfaceContainer)
		if !ok {
			return false
		}
		return EqualsInterfaceContainer(a, b)
	case InterfaceSlice:
		b, ok := inB.(InterfaceSlice)
		if !ok {
			return false
		}
		return EqualsInterfaceSlice(a, b)
	case *Leaf:
		b, ok := inB.(*Leaf)
		if !ok {
			return false
		}
		return EqualsRefOfLeaf(a, b)
	case LeafSlice:
		b, ok := inB.(LeafSlice)
		if !ok {
			return false
		}
		return EqualsLeafSlice(a, b)
	case *NoCloneType:
		b, ok := inB.(*NoCloneType)
		if !ok {
			return false
		}
		return EqualsRefOfNoCloneType(a, b)
	case *RefContainer:
		b, ok := inB.(*RefContainer)
		if !ok {
			return false
		}
		return EqualsRefOfRefContainer(a, b)
	case *RefSliceContainer:
		b, ok := inB.(*RefSliceContainer)
		if !ok {
			return false
		}
		return EqualsRefOfRefSliceContainer(a, b)
	case *SubImpl:
		b, ok := inB.(*SubImpl)
		if !ok {
			return false
		}
		return EqualsRefOfSubImpl(a, b)
	case ValueContainer:
		b, ok := inB.(ValueContainer)
		if !ok {
			return false
		}
		return EqualsValueContainer(a, b)
	case ValueSliceContainer:
		b, ok := inB.(ValueSliceContainer)
		if !ok {
			return false
		}
		return EqualsValueSliceContainer(a, b)
	default:
		// this should never happen
		return false
	}
}

// CloneAST creates a deep clone of the input.
func CloneAST(in AST) AST {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case BasicType:
		return in
	case Bytes:
		return CloneBytes(in)
	case InterfaceContainer:
		return CloneInterfaceContainer(in)
	case InterfaceSlice:
		return CloneInterfaceSlice(in)
	case *Leaf:
		return CloneRefOfLeaf(in)
	case LeafSlice:
		return CloneLeafSlice(in)
	case *NoCloneType:
		return CloneRefOfNoCloneType(in)
	case *RefContainer:
		return CloneRefOfRefContainer(in)
	case *RefSliceContainer:
		return CloneRefOfRefSliceContainer(in)
	case *SubImpl:
		return CloneRefOfSubImpl(in)
	case ValueContainer:
		return CloneValueContainer(in)
	case ValueSliceContainer:
		return CloneValueSliceContainer(in)
	default:
		// this should never happen
		return nil
	}
}

// VisitAST will visit all parts of the AST
func VisitAST(in AST, f Visit) error {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case BasicType:
		return VisitBasicType(in, f)
	case Bytes:
		return VisitBytes(in, f)
	case InterfaceContainer:
		return VisitInterfaceContainer(in, f)
	case InterfaceSlice:
		return VisitInterfaceSlice(in, f)
	case *Leaf:
		return VisitRefOfLeaf(in, f)
	case LeafSlice:
		return VisitLeafSlice(in, f)
	case *NoCloneType:
		return VisitRefOfNoCloneType(in, f)
	case *RefContainer:
		return VisitRefOfRefContainer(in, f)
	case *RefSliceContainer:
		return VisitRefOfRefSliceContainer(in, f)
	case *SubImpl:
		return VisitRefOfSubImpl(in, f)
	case ValueContainer:
		return VisitValueContainer(in, f)
	case ValueSliceContainer:
		return VisitValueSliceContainer(in, f)
	default:
		// this should never happen
		return nil
	}
}

// rewriteAST is part of the Rewrite implementation
func rewriteAST(parent AST, node AST, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	switch node := node.(type) {
	case BasicType:
		return rewriteBasicType(parent, node, replacer, pre, post)
	case Bytes:
		return rewriteBytes(parent, node, replacer, pre, post)
	case InterfaceContainer:
		return rewriteInterfaceContainer(parent, node, replacer, pre, post)
	case InterfaceSlice:
		return rewriteInterfaceSlice(parent, node, replacer, pre, post)
	case *Leaf:
		return rewriteRefOfLeaf(parent, node, replacer, pre, post)
	case LeafSlice:
		return rewriteLeafSlice(parent, node, replacer, pre, post)
	case *NoCloneType:
		return rewriteRefOfNoCloneType(parent, node, replacer, pre, post)
	case *RefContainer:
		return rewriteRefOfRefContainer(parent, node, replacer, pre, post)
	case *RefSliceContainer:
		return rewriteRefOfRefSliceContainer(parent, node, replacer, pre, post)
	case *SubImpl:
		return rewriteRefOfSubImpl(parent, node, replacer, pre, post)
	case ValueContainer:
		return rewriteValueContainer(parent, node, replacer, pre, post)
	case ValueSliceContainer:
		return rewriteValueSliceContainer(parent, node, replacer, pre, post)
	default:
		// this should never happen
		return nil
	}
}

// EqualsBytes does deep equals between the two objects.
func EqualsBytes(a, b Bytes) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CloneBytes creates a deep clone of the input.
func CloneBytes(n Bytes) Bytes {
	res := make(Bytes, 0, len(n))
	copy(res, n)
	return res
}

// VisitBytes will visit all parts of the AST
func VisitBytes(in Bytes, f Visit) error {
	_, err := f(in)
	return err
}

// rewriteBytes is part of the Rewrite implementation
func rewriteBytes(parent AST, node Bytes, replacer replacerFunc, pre, post ApplyFunc) error {
	// ptrToStructMethod
}

// EqualsInterfaceContainer does deep equals between the two objects.
func EqualsInterfaceContainer(a, b InterfaceContainer) bool {
	return true
}

// CloneInterfaceContainer creates a deep clone of the input.
func CloneInterfaceContainer(n InterfaceContainer) InterfaceContainer {
	return *CloneRefOfInterfaceContainer(&n)
}

// VisitInterfaceContainer will visit all parts of the AST
func VisitInterfaceContainer(in InterfaceContainer, f Visit) error {
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	return nil
}

// rewriteInterfaceContainer is part of the Rewrite implementation
func rewriteInterfaceContainer(parent AST, node InterfaceContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	var err error
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if err != nil {
		return err
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsInterfaceSlice does deep equals between the two objects.
func EqualsInterfaceSlice(a, b InterfaceSlice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !EqualsAST(a[i], b[i]) {
			return false
		}
	}
	return true
}

// CloneInterfaceSlice creates a deep clone of the input.
func CloneInterfaceSlice(n InterfaceSlice) InterfaceSlice {
	res := make(InterfaceSlice, 0, len(n))
	for _, x := range n {
		res = append(res, CloneAST(x))
	}
	return res
}

// VisitInterfaceSlice will visit all parts of the AST
func VisitInterfaceSlice(in InterfaceSlice, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	for _, el := range in {
		if err := VisitAST(el, f); err != nil {
			return err
		}
	}
	return nil
}

// rewriteInterfaceSlice is part of the Rewrite implementation
func rewriteInterfaceSlice(parent AST, node InterfaceSlice, replacer replacerFunc, pre, post ApplyFunc) error {
	// ptrToStructMethod
}

// EqualsRefOfLeaf does deep equals between the two objects.
func EqualsRefOfLeaf(a, b *Leaf) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.v == b.v
}

// CloneRefOfLeaf creates a deep clone of the input.
func CloneRefOfLeaf(n *Leaf) *Leaf {
	if n == nil {
		return nil
	}
	out := *n
	return &out
}

// VisitRefOfLeaf will visit all parts of the AST
func VisitRefOfLeaf(in *Leaf, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	return nil
}

// rewriteRefOfLeaf is part of the Rewrite implementation
func rewriteRefOfLeaf(parent AST, node *Leaf, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsLeafSlice does deep equals between the two objects.
func EqualsLeafSlice(a, b LeafSlice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !EqualsRefOfLeaf(a[i], b[i]) {
			return false
		}
	}
	return true
}

// CloneLeafSlice creates a deep clone of the input.
func CloneLeafSlice(n LeafSlice) LeafSlice {
	res := make(LeafSlice, 0, len(n))
	for _, x := range n {
		res = append(res, CloneRefOfLeaf(x))
	}
	return res
}

// VisitLeafSlice will visit all parts of the AST
func VisitLeafSlice(in LeafSlice, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	for _, el := range in {
		if err := VisitRefOfLeaf(el, f); err != nil {
			return err
		}
	}
	return nil
}

// rewriteLeafSlice is part of the Rewrite implementation
func rewriteLeafSlice(parent AST, node LeafSlice, replacer replacerFunc, pre, post ApplyFunc) error {
	// ptrToStructMethod
}

// EqualsRefOfNoCloneType does deep equals between the two objects.
func EqualsRefOfNoCloneType(a, b *NoCloneType) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.v == b.v
}

// CloneRefOfNoCloneType creates a deep clone of the input.
func CloneRefOfNoCloneType(n *NoCloneType) *NoCloneType {
	return n
}

// VisitRefOfNoCloneType will visit all parts of the AST
func VisitRefOfNoCloneType(in *NoCloneType, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	return nil
}

// rewriteRefOfNoCloneType is part of the Rewrite implementation
func rewriteRefOfNoCloneType(parent AST, node *NoCloneType, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsRefOfRefContainer does deep equals between the two objects.
func EqualsRefOfRefContainer(a, b *RefContainer) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.NotASTType == b.NotASTType &&
		EqualsAST(a.ASTType, b.ASTType) &&
		EqualsRefOfLeaf(a.ASTImplementationType, b.ASTImplementationType)
}

// CloneRefOfRefContainer creates a deep clone of the input.
func CloneRefOfRefContainer(n *RefContainer) *RefContainer {
	if n == nil {
		return nil
	}
	out := *n
	out.ASTType = CloneAST(n.ASTType)
	out.ASTImplementationType = CloneRefOfLeaf(n.ASTImplementationType)
	return &out
}

// VisitRefOfRefContainer will visit all parts of the AST
func VisitRefOfRefContainer(in *RefContainer, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	if err := VisitAST(in.ASTType, f); err != nil {
		return err
	}
	if err := VisitRefOfLeaf(in.ASTImplementationType, f); err != nil {
		return err
	}
	return nil
}

// rewriteRefOfRefContainer is part of the Rewrite implementation
func rewriteRefOfRefContainer(parent AST, node *RefContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if errF := rewriteAST(node, node.ASTType, func(newNode, parent AST) {
		parent.(*RefContainer).ASTType = newNode.(AST)
	}, pre, post); errF != nil {
		return errF
	}
	if errF := rewriteRefOfLeaf(node, node.ASTImplementationType, func(newNode, parent AST) {
		parent.(*RefContainer).ASTImplementationType = newNode.(*Leaf)
	}, pre, post); errF != nil {
		return errF
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsRefOfRefSliceContainer does deep equals between the two objects.
func EqualsRefOfRefSliceContainer(a, b *RefSliceContainer) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return EqualsSliceOfAST(a.ASTElements, b.ASTElements) &&
		EqualsSliceOfInt(a.NotASTElements, b.NotASTElements) &&
		EqualsSliceOfRefOfLeaf(a.ASTImplementationElements, b.ASTImplementationElements)
}

// CloneRefOfRefSliceContainer creates a deep clone of the input.
func CloneRefOfRefSliceContainer(n *RefSliceContainer) *RefSliceContainer {
	if n == nil {
		return nil
	}
	out := *n
	out.ASTElements = CloneSliceOfAST(n.ASTElements)
	out.NotASTElements = CloneSliceOfInt(n.NotASTElements)
	out.ASTImplementationElements = CloneSliceOfRefOfLeaf(n.ASTImplementationElements)
	return &out
}

// VisitRefOfRefSliceContainer will visit all parts of the AST
func VisitRefOfRefSliceContainer(in *RefSliceContainer, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	for _, el := range in.ASTElements {
		if err := VisitAST(el, f); err != nil {
			return err
		}
	}
	for _, el := range in.ASTImplementationElements {
		if err := VisitRefOfLeaf(el, f); err != nil {
			return err
		}
	}
	return nil
}

// rewriteRefOfRefSliceContainer is part of the Rewrite implementation
func rewriteRefOfRefSliceContainer(parent AST, node *RefSliceContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	for i, el := range node.ASTElements {
		if errF := rewriteAST(node, el, func(newNode, parent AST) {
			parent.(*RefSliceContainer).ASTElements[i] = newNode.(AST)
		}, pre, post); errF != nil {
			return errF
		}
	}
	for i, el := range node.ASTImplementationElements {
		if errF := rewriteRefOfLeaf(node, el, func(newNode, parent AST) {
			parent.(*RefSliceContainer).ASTImplementationElements[i] = newNode.(*Leaf)
		}, pre, post); errF != nil {
			return errF
		}
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsRefOfSubImpl does deep equals between the two objects.
func EqualsRefOfSubImpl(a, b *SubImpl) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return EqualsSubIface(a.inner, b.inner) &&
		EqualsRefOfBool(a.field, b.field)
}

// CloneRefOfSubImpl creates a deep clone of the input.
func CloneRefOfSubImpl(n *SubImpl) *SubImpl {
	if n == nil {
		return nil
	}
	out := *n
	out.inner = CloneSubIface(n.inner)
	out.field = CloneRefOfBool(n.field)
	return &out
}

// VisitRefOfSubImpl will visit all parts of the AST
func VisitRefOfSubImpl(in *SubImpl, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	if err := VisitSubIface(in.inner, f); err != nil {
		return err
	}
	return nil
}

// rewriteRefOfSubImpl is part of the Rewrite implementation
func rewriteRefOfSubImpl(parent AST, node *SubImpl, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if errF := rewriteSubIface(node, node.inner, func(newNode, parent AST) {
		parent.(*SubImpl).inner = newNode.(SubIface)
	}, pre, post); errF != nil {
		return errF
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsValueContainer does deep equals between the two objects.
func EqualsValueContainer(a, b ValueContainer) bool {
	return a.NotASTType == b.NotASTType &&
		EqualsAST(a.ASTType, b.ASTType) &&
		EqualsRefOfLeaf(a.ASTImplementationType, b.ASTImplementationType)
}

// CloneValueContainer creates a deep clone of the input.
func CloneValueContainer(n ValueContainer) ValueContainer {
	return *CloneRefOfValueContainer(&n)
}

// VisitValueContainer will visit all parts of the AST
func VisitValueContainer(in ValueContainer, f Visit) error {
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	if err := VisitAST(in.ASTType, f); err != nil {
		return err
	}
	if err := VisitRefOfLeaf(in.ASTImplementationType, f); err != nil {
		return err
	}
	return nil
}

// rewriteValueContainer is part of the Rewrite implementation
func rewriteValueContainer(parent AST, node ValueContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	var err error
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if errF := rewriteAST(node, node.ASTType, func(newNode, parent AST) {
		err = vterrors.New(vtrpc.Code_INTERNAL, "[BUG] tried to replace 'ASTType' on 'ValueContainer'")
	}, pre, post); errF != nil {
		return errF
	}
	if errF := rewriteRefOfLeaf(node, node.ASTImplementationType, func(newNode, parent AST) {
		err = vterrors.New(vtrpc.Code_INTERNAL, "[BUG] tried to replace 'ASTImplementationType' on 'ValueContainer'")
	}, pre, post); errF != nil {
		return errF
	}
	if err != nil {
		return err
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsValueSliceContainer does deep equals between the two objects.
func EqualsValueSliceContainer(a, b ValueSliceContainer) bool {
	return EqualsSliceOfAST(a.ASTElements, b.ASTElements) &&
		EqualsSliceOfInt(a.NotASTElements, b.NotASTElements) &&
		EqualsSliceOfRefOfLeaf(a.ASTImplementationElements, b.ASTImplementationElements)
}

// CloneValueSliceContainer creates a deep clone of the input.
func CloneValueSliceContainer(n ValueSliceContainer) ValueSliceContainer {
	return *CloneRefOfValueSliceContainer(&n)
}

// VisitValueSliceContainer will visit all parts of the AST
func VisitValueSliceContainer(in ValueSliceContainer, f Visit) error {
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	for _, el := range in.ASTElements {
		if err := VisitAST(el, f); err != nil {
			return err
		}
	}
	for _, el := range in.ASTImplementationElements {
		if err := VisitRefOfLeaf(el, f); err != nil {
			return err
		}
	}
	return nil
}

// rewriteValueSliceContainer is part of the Rewrite implementation
func rewriteValueSliceContainer(parent AST, node ValueSliceContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	var err error
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	for _, el := range node.ASTElements {
		if errF := rewriteAST(node, el, func(newNode, parent AST) {
			err = vterrors.New(vtrpc.Code_INTERNAL, "[BUG] tried to replace 'ASTElements' on 'ValueSliceContainer'")
		}, pre, post); errF != nil {
			return errF
		}
	}
	for _, el := range node.ASTImplementationElements {
		if errF := rewriteRefOfLeaf(node, el, func(newNode, parent AST) {
			err = vterrors.New(vtrpc.Code_INTERNAL, "[BUG] tried to replace 'ASTImplementationElements' on 'ValueSliceContainer'")
		}, pre, post); errF != nil {
			return errF
		}
	}
	if err != nil {
		return err
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsSubIface does deep equals between the two objects.
func EqualsSubIface(inA, inB SubIface) bool {
	if inA == nil && inB == nil {
		return true
	}
	if inA == nil || inB == nil {
		return false
	}
	switch a := inA.(type) {
	case *SubImpl:
		b, ok := inB.(*SubImpl)
		if !ok {
			return false
		}
		return EqualsRefOfSubImpl(a, b)
	default:
		// this should never happen
		return false
	}
}

// CloneSubIface creates a deep clone of the input.
func CloneSubIface(in SubIface) SubIface {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case *SubImpl:
		return CloneRefOfSubImpl(in)
	default:
		// this should never happen
		return nil
	}
}

// VisitSubIface will visit all parts of the AST
func VisitSubIface(in SubIface, f Visit) error {
	if in == nil {
		return nil
	}
	switch in := in.(type) {
	case *SubImpl:
		return VisitRefOfSubImpl(in, f)
	default:
		// this should never happen
		return nil
	}
}

// rewriteSubIface is part of the Rewrite implementation
func rewriteSubIface(parent AST, node SubIface, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	switch node := node.(type) {
	case *SubImpl:
		return rewriteRefOfSubImpl(parent, node, replacer, pre, post)
	default:
		// this should never happen
		return nil
	}
}

// VisitBasicType will visit all parts of the AST
func VisitBasicType(in BasicType, f Visit) error {
	_, err := f(in)
	return err
}

// rewriteBasicType is part of the Rewrite implementation
func rewriteBasicType(parent AST, node BasicType, replacer replacerFunc, pre, post ApplyFunc) error {
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsRefOfInterfaceContainer does deep equals between the two objects.
func EqualsRefOfInterfaceContainer(a, b *InterfaceContainer) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return true
}

// CloneRefOfInterfaceContainer creates a deep clone of the input.
func CloneRefOfInterfaceContainer(n *InterfaceContainer) *InterfaceContainer {
	if n == nil {
		return nil
	}
	out := *n
	out.v = n.v
	return &out
}

// VisitRefOfInterfaceContainer will visit all parts of the AST
func VisitRefOfInterfaceContainer(in *InterfaceContainer, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	return nil
}

// rewriteRefOfInterfaceContainer is part of the Rewrite implementation
func rewriteRefOfInterfaceContainer(parent AST, node *InterfaceContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsSliceOfAST does deep equals between the two objects.
func EqualsSliceOfAST(a, b []AST) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !EqualsAST(a[i], b[i]) {
			return false
		}
	}
	return true
}

// CloneSliceOfAST creates a deep clone of the input.
func CloneSliceOfAST(n []AST) []AST {
	res := make([]AST, 0, len(n))
	for _, x := range n {
		res = append(res, CloneAST(x))
	}
	return res
}

// EqualsSliceOfInt does deep equals between the two objects.
func EqualsSliceOfInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CloneSliceOfInt creates a deep clone of the input.
func CloneSliceOfInt(n []int) []int {
	res := make([]int, 0, len(n))
	copy(res, n)
	return res
}

// EqualsSliceOfRefOfLeaf does deep equals between the two objects.
func EqualsSliceOfRefOfLeaf(a, b []*Leaf) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !EqualsRefOfLeaf(a[i], b[i]) {
			return false
		}
	}
	return true
}

// CloneSliceOfRefOfLeaf creates a deep clone of the input.
func CloneSliceOfRefOfLeaf(n []*Leaf) []*Leaf {
	res := make([]*Leaf, 0, len(n))
	for _, x := range n {
		res = append(res, CloneRefOfLeaf(x))
	}
	return res
}

// EqualsRefOfBool does deep equals between the two objects.
func EqualsRefOfBool(a, b *bool) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// CloneRefOfBool creates a deep clone of the input.
func CloneRefOfBool(n *bool) *bool {
	if n == nil {
		return nil
	}
	out := *n
	return &out
}

// EqualsRefOfValueContainer does deep equals between the two objects.
func EqualsRefOfValueContainer(a, b *ValueContainer) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.NotASTType == b.NotASTType &&
		EqualsAST(a.ASTType, b.ASTType) &&
		EqualsRefOfLeaf(a.ASTImplementationType, b.ASTImplementationType)
}

// CloneRefOfValueContainer creates a deep clone of the input.
func CloneRefOfValueContainer(n *ValueContainer) *ValueContainer {
	if n == nil {
		return nil
	}
	out := *n
	out.ASTType = CloneAST(n.ASTType)
	out.ASTImplementationType = CloneRefOfLeaf(n.ASTImplementationType)
	return &out
}

// VisitRefOfValueContainer will visit all parts of the AST
func VisitRefOfValueContainer(in *ValueContainer, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	if err := VisitAST(in.ASTType, f); err != nil {
		return err
	}
	if err := VisitRefOfLeaf(in.ASTImplementationType, f); err != nil {
		return err
	}
	return nil
}

// rewriteRefOfValueContainer is part of the Rewrite implementation
func rewriteRefOfValueContainer(parent AST, node *ValueContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	if errF := rewriteAST(node, node.ASTType, func(newNode, parent AST) {
		parent.(*ValueContainer).ASTType = newNode.(AST)
	}, pre, post); errF != nil {
		return errF
	}
	if errF := rewriteRefOfLeaf(node, node.ASTImplementationType, func(newNode, parent AST) {
		parent.(*ValueContainer).ASTImplementationType = newNode.(*Leaf)
	}, pre, post); errF != nil {
		return errF
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}

// EqualsRefOfValueSliceContainer does deep equals between the two objects.
func EqualsRefOfValueSliceContainer(a, b *ValueSliceContainer) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return EqualsSliceOfAST(a.ASTElements, b.ASTElements) &&
		EqualsSliceOfInt(a.NotASTElements, b.NotASTElements) &&
		EqualsSliceOfRefOfLeaf(a.ASTImplementationElements, b.ASTImplementationElements)
}

// CloneRefOfValueSliceContainer creates a deep clone of the input.
func CloneRefOfValueSliceContainer(n *ValueSliceContainer) *ValueSliceContainer {
	if n == nil {
		return nil
	}
	out := *n
	out.ASTElements = CloneSliceOfAST(n.ASTElements)
	out.NotASTElements = CloneSliceOfInt(n.NotASTElements)
	out.ASTImplementationElements = CloneSliceOfRefOfLeaf(n.ASTImplementationElements)
	return &out
}

// VisitRefOfValueSliceContainer will visit all parts of the AST
func VisitRefOfValueSliceContainer(in *ValueSliceContainer, f Visit) error {
	if in == nil {
		return nil
	}
	if cont, err := f(in); err != nil || !cont {
		return err
	}
	for _, el := range in.ASTElements {
		if err := VisitAST(el, f); err != nil {
			return err
		}
	}
	for _, el := range in.ASTImplementationElements {
		if err := VisitRefOfLeaf(el, f); err != nil {
			return err
		}
	}
	return nil
}

// rewriteRefOfValueSliceContainer is part of the Rewrite implementation
func rewriteRefOfValueSliceContainer(parent AST, node *ValueSliceContainer, replacer replacerFunc, pre, post ApplyFunc) error {
	if node == nil {
		return nil
	}
	cur := Cursor{
		node:     node,
		parent:   parent,
		replacer: replacer,
	}
	if !pre(&cur) {
		return nil
	}
	for i, el := range node.ASTElements {
		if errF := rewriteAST(node, el, func(newNode, parent AST) {
			parent.(*ValueSliceContainer).ASTElements[i] = newNode.(AST)
		}, pre, post); errF != nil {
			return errF
		}
	}
	for i, el := range node.ASTImplementationElements {
		if errF := rewriteRefOfLeaf(node, el, func(newNode, parent AST) {
			parent.(*ValueSliceContainer).ASTImplementationElements[i] = newNode.(*Leaf)
		}, pre, post); errF != nil {
			return errF
		}
	}
	if !post(&cur) {
		return abortE
	}
	return nil
}
