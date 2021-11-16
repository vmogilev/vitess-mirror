/*
   Copyright 2016 GitHub Inc.
	 See https://github.com/github/gh-ost/blob/master/LICENSE
*/

package vrepl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	columnsA = &ColumnList{
		columns: []Column{
			{
				Name: "id",
			},
			{
				Name: "cint",
			},
			{
				Name: "cgen1",
			},
			{
				Name: "cgen2",
			},
			{
				Name: "cchar",
			},
			{
				Name: "cremoved",
			},
			{
				Name:       "cnullable",
				IsNullable: true,
			},
			{
				Name:          "cnodefault",
				IsNullable:    false,
				IsDefaultNull: true,
			},
		},
		Ordinals: ColumnsMap{},
	}
	columnsB = &ColumnList{
		columns: []Column{
			{
				Name: "id",
			},
			{
				Name: "cint",
			},
			{
				Name: "cgen1",
			},
			{
				Name: "cchar_alternate",
			},
			{
				Name:       "cnullable",
				IsNullable: true,
			},
			{
				Name:          "cnodefault",
				IsNullable:    false,
				IsDefaultNull: true,
			},
		},
		Ordinals: ColumnsMap{},
	}
	columnsVirtual = ParseColumnList("cgen1,cgen2")
)

func TestGetSharedColumns(t *testing.T) {
	tt := []struct {
		name                                    string
		sourceCols                              *ColumnList
		targetCols                              *ColumnList
		renameMap                               map[string]string
		expectSourceSharedColNames              []string
		expectTargetSharedColNames              []string
		expectDroppedSourceNonGeneratedColNames []string
	}{
		{
			name:                                    "rename map empty",
			sourceCols:                              columnsA,
			targetCols:                              columnsB,
			renameMap:                               map[string]string{},
			expectSourceSharedColNames:              []string{"id", "cint", "cnullable", "cnodefault"},
			expectTargetSharedColNames:              []string{"id", "cint", "cnullable", "cnodefault"},
			expectDroppedSourceNonGeneratedColNames: []string{"cchar", "cremoved"},
		},
		{
			name:                                    "renamed column",
			sourceCols:                              columnsA,
			targetCols:                              columnsB,
			renameMap:                               map[string]string{"cchar": "cchar_alternate"},
			expectSourceSharedColNames:              []string{"id", "cint", "cchar", "cnullable", "cnodefault"},
			expectTargetSharedColNames:              []string{"id", "cint", "cchar_alternate", "cnullable", "cnodefault"},
			expectDroppedSourceNonGeneratedColNames: []string{"cremoved"},
		},
	}

	parser := NewAlterTableParser()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			parser.columnRenameMap = tc.renameMap
			sourceSharedCols, targetSharedCols, droppedNonGeneratedCols, _ := GetSharedColumns(
				tc.sourceCols, tc.targetCols,
				columnsVirtual, columnsVirtual,
				parser,
			)
			assert.Equal(t, tc.expectSourceSharedColNames, sourceSharedCols.Names())
			assert.Equal(t, tc.expectTargetSharedColNames, targetSharedCols.Names())
			assert.Equal(t, tc.expectDroppedSourceNonGeneratedColNames, droppedNonGeneratedCols.Names())
		})
	}
}
