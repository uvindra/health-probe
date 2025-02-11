package generator

import (
	"testing"
)

func TestGetRandomName(t *testing.T) {
	var names []string
	indexes := make(map[int]struct{})

	for i := 0; i < 10; i++ {
		names = append(names, GetRandomName())
		indexes[i] = struct{}{}
	}

	for i := 0; i < len(names); i++ {
		tmp := make(map[int]struct{})

		for key, value := range indexes {
			tmp[key] = value
		}

		delete(tmp, i)

		for key := range tmp {
			if names[i] == names[key] {
				t.Errorf("Name %s is not duplicated", names[i])
			}
		}
	}

}
