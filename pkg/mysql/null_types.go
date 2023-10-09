package mysql

import "database/sql"

func NewNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: v, Valid: true}
}

func NewNullInt64(v int64) sql.NullInt64 {
	if v == 0 {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: v, Valid: true}
}

func NewNullFloat64(v float64) sql.NullFloat64 {
	if v == 0 {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: v, Valid: true}
}
