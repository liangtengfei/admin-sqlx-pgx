package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

func NewSQLNullString(s string) sql.NullString {
	if len(strings.TrimSpace(s)) > 0 {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	} else {
		return sql.NullString{
			Valid:  false,
			String: s,
		}
	}
}

func SQLNullStringDefault(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func SQLNullInt64Default(s sql.NullInt64) int64 {
	if s.Valid {
		return s.Int64
	}
	return int64(0)
}

func StringConcat(str ...string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}

func Int64Join(a []int64) string {
	if len(a) > 0 {
		var b []string
		for _, s := range a {
			b = append(b, fmt.Sprintf("%d", s))
		}
		return strings.Join(b, ",")
	}
	return ""
}

func Int2String(in interface{}) string {
	return fmt.Sprintf("%d", in)
}

// TernaryOperation 三元运算
func TernaryOperation(ok bool, yes interface{}, no interface{}) interface{} {
	if ok {
		return yes
	}
	return no
}
