// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateStmt struct {
	Relation     *RangeVar `json:"relation"`     /* relation to create */
	TableElts    []Node    `json:"tableElts"`    /* column definitions (list of ColumnDef) */
	InhRelations []Node    `json:"inhRelations"` /* relations to inherit from (list of
	 * inhRelation) */

	OfTypename     *TypeName      `json:"ofTypename"`     /* OF typename */
	Constraints    []Node         `json:"constraints"`    /* constraints (list of Constraint nodes) */
	Options        []Node         `json:"options"`        /* options from WITH clause */
	Oncommit       OnCommitAction `json:"oncommit"`       /* what do we do at COMMIT? */
	Tablespacename *string        `json:"tablespacename"` /* table space to use, or NULL */
	IfNotExists    bool           `json:"if_not_exists"`  /* just do nothing if it already exists? */
}

func (node CreateStmt) MarshalJSON() ([]byte, error) {
	type CreateStmtMarshalAlias CreateStmt
	return json.Marshal(map[string]interface{}{
		"CREATESTMT": (*CreateStmtMarshalAlias)(&node),
	})
}

func (node *CreateStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["relation"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["relation"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeVar)
			node.Relation = &val
		}
	}

	if fields["tableElts"] != nil {
		node.TableElts, err = UnmarshalNodeArrayJSON(fields["tableElts"])
		if err != nil {
			return
		}
	}

	if fields["inhRelations"] != nil {
		node.InhRelations, err = UnmarshalNodeArrayJSON(fields["inhRelations"])
		if err != nil {
			return
		}
	}

	if fields["ofTypename"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["ofTypename"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(TypeName)
			node.OfTypename = &val
		}
	}

	if fields["constraints"] != nil {
		node.Constraints, err = UnmarshalNodeArrayJSON(fields["constraints"])
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	if fields["oncommit"] != nil {
		err = json.Unmarshal(fields["oncommit"], &node.Oncommit)
		if err != nil {
			return
		}
	}

	if fields["tablespacename"] != nil {
		err = json.Unmarshal(fields["tablespacename"], &node.Tablespacename)
		if err != nil {
			return
		}
	}

	if fields["if_not_exists"] != nil {
		err = json.Unmarshal(fields["if_not_exists"], &node.IfNotExists)
		if err != nil {
			return
		}
	}

	return
}
