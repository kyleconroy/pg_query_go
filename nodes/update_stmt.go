// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type UpdateStmt struct {
	Relation      *RangeVar   `json:"relation"`      /* relation to update */
	TargetList    []Node      `json:"targetList"`    /* the target list (of ResTarget) */
	WhereClause   Node        `json:"whereClause"`   /* qualifications */
	FromClause    []Node      `json:"fromClause"`    /* optional from clause for more tables */
	ReturningList []Node      `json:"returningList"` /* list of expressions to return */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */
}

func (node UpdateStmt) MarshalJSON() ([]byte, error) {
	type UpdateStmtMarshalAlias UpdateStmt
	return json.Marshal(map[string]interface{}{
		"UPDATE": (*UpdateStmtMarshalAlias)(&node),
	})
}

func (node *UpdateStmt) UnmarshalJSON(input []byte) (err error) {
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

	if fields["targetList"] != nil {
		node.TargetList, err = UnmarshalNodeArrayJSON(fields["targetList"])
		if err != nil {
			return
		}
	}

	if fields["whereClause"] != nil {
		node.WhereClause, err = UnmarshalNodeJSON(fields["whereClause"])
		if err != nil {
			return
		}
	}

	if fields["fromClause"] != nil {
		node.FromClause, err = UnmarshalNodeArrayJSON(fields["fromClause"])
		if err != nil {
			return
		}
	}

	if fields["returningList"] != nil {
		node.ReturningList, err = UnmarshalNodeArrayJSON(fields["returningList"])
		if err != nil {
			return
		}
	}

	if fields["withClause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["withClause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(WithClause)
			node.WithClause = &val
		}
	}

	return
}
