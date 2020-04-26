package util

import "github.com/volatiletech/null"

func ToNullableString(v null.String) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}
