package sortstruct

import (
	"reflect"
	"sort"
)

func Prop(field string, asc bool) func(p1, p2 interface{}) bool {
	return func(p1, p2 interface{}) bool {

		v1 := reflect.Indirect(reflect.ValueOf(p1)).FieldByName(field)
		v2 := reflect.Indirect(reflect.ValueOf(p2)).FieldByName(field)

		ret := false

		switch v1.Kind() {
		case reflect.Int64:
			ret = int64(v1.Int()) < int64(v2.Int())
		case reflect.Int:
			ret = v1.Int() < v2.Int()
		case reflect.Float64:
			ret = float64(v1.Float()) < float64(v2.Float())
		case reflect.Float32:
			ret = float32(v1.Float()) < float32(v2.Float())
		case reflect.String:
			ret = string(v1.String()) < string(v2.String())
		}

		if asc {
			return ret
		}
		return !ret
	}
}

type By func(p1, p2 interface{}) bool

func (by By) Sort(entries []interface{}) {
	ps := &entriesSort{
		entries: entries,
		by:      by,
	}
	sort.Sort(ps)
}

type entriesSort struct {
	entries []interface{}
	by      func(p1, p2 interface{}) bool
}

func (s *entriesSort) Len() int { return len(s.entries) }

func (s *entriesSort) Swap(i, j int) {
	s.entries[i], s.entries[j] = s.entries[j], s.entries[i]
}

func (s *entriesSort) Less(i, j int) bool {
	return s.by(s.entries[i], s.entries[j])
}
