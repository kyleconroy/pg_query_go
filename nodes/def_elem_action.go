// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type DefElemAction uint

const (
	DEFELEM_UNSPEC = iota /* no action given */
	DEFELEM_SET
	DEFELEM_ADD
	DEFELEM_DROP
)
