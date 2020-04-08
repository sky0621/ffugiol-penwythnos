package sys

import "github.com/volatiletech/null"

func ToInt64(v null.Int64) int64 {
	if v.Valid {
		return v.Int64
	}
	return 0
}

func ToInt(v null.Int) int {
	if v.Valid {
		return v.Int
	}
	return 0
}

func ToString(v null.String) string {
	if v.Valid {
		return v.String
	}
	return ""
}

func ToBool(v null.Bool) bool {
	if v.Valid {
		return v.Bool
	}
	return false
}
