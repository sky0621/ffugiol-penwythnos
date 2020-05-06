package util

type StringSet struct {
	set map[string]struct{}
}

func (s *StringSet) Add(v string) {
	s.set[v] = struct{}{}
}

func (s *StringSet) Values() []string {
	var values []string
	for v := range s.set {
		values = append(values, v)
	}
	return values
}

func (s *StringSet) Contains(v string) bool {
	for _, value := range s.Values() {
		if value == v {
			return true
		}
	}
	return false
}

func NewBlankStringSet() *StringSet {
	set := make(map[string]struct{})
	return &StringSet{set: set}
}

func NewStringSet(s []string) *StringSet {
	stringSet := NewBlankStringSet()
	for _, v := range s {
		stringSet.Add(v)
	}
	return stringSet
}
