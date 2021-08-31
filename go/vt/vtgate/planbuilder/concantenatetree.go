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

package planbuilder

import (
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vtgate/semantics"
)

type (
	concatenateTree struct {
		sources []queryTree
	}
)

var _ queryTree = (*concatenateTree)(nil)

func (c *concatenateTree) tableID() semantics.TableSet {
	var tableSet semantics.TableSet
	for _, source := range c.sources {
		tableSet |= source.tableID()
	}
	return tableSet
}

func (c *concatenateTree) clone() queryTree {
	var sourcesCopy []queryTree
	for _, source := range c.sources {
		sourcesCopy = append(sourcesCopy, source.clone())
	}
	return &concatenateTree{sources: sourcesCopy}
}

func (c *concatenateTree) cost() int {
	var totalCost int
	for _, source := range c.sources {
		totalCost += source.cost()
	}
	return totalCost
}

func (c *concatenateTree) pushOutputColumns(columns []*sqlparser.ColName, semTable *semantics.SemTable) ([]int, error) {
	panic("implement me")
	//var indexPushed []int
	//sourceColumns := make([][]*sqlparser.ColName, len(c.sources))
	//for _, col := range columns {
	//	deps := semTable.BaseTableDependencies(col)
	//	isSolved := false
	//	for i, source := range c.sources {
	//		if deps.IsSolvedBy(source.tableID()) {
	//			isSolved = true
	//			sourceColumns[i] = append(sourceColumns[i], col)
	//			indexPushed = append(indexPushed, i)
	//		}
	//	}
	//	if !isSolved {
	//		return nil, vterrors.Errorf(vtrpcpb.Code_INTERNAL, "[BUG]: unable to push output column to any of the sources in concatenate")
	//	}
	//}
	//
	//offSetColumns := make([][]int, len(c.sources))
	//for i, source := range c.sources {
	//	offsets, err := source.pushOutputColumns(sourceColumns[i], semTable)
	//	if err != nil {
	//		return nil, err
	//	}
	//	offSetColumns[i] = offsets
	//}
	//
	//outputColumns := make([]int, len(indexPushed))
	//return outputColumns, nil
}
