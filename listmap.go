package listmap

import (
	log "github.com/sirupsen/logrus"
	"sort"
)

type MyMap struct {
	keys, vals []string
}

func (m *MyMap) ToMap() map[string]string {
	out := make(map[string]string)
	for i := range m.keys {
		k := m.keys[i]
		v := m.vals[i]
		out[k] = v
	}
	return out
}

func (m *MyMap) IsEmpty() bool {
	return len(m.keys) == 0
}

func (m *MyMap) Contains(k string) bool {
	if len(m.keys) == 0 {
		return false
	}
	i := sort.SearchStrings(m.keys, k)
	if i == len(m.keys) {
		return false
	}
	return m.keys[i] == k
}

func (m *MyMap) containsKey(k string, i int) bool {
	if len(m.keys) == 0 {
		return false
	}
	if i == len(m.keys) {
		return false
	}
	return m.keys[i] == k
}

func (m *MyMap) Put(k, v string) {
	keyIndex := sort.SearchStrings(m.keys, k)
	valIndex := sort.SearchStrings(m.vals, v)

	log.Infof("keyIndex = %d, valIndex = %d", keyIndex, valIndex)

	if m.containsKey(k, keyIndex) {
		log.Infof("Map contains key: %s", k)
		if valIndex == len(m.vals) {
			m.vals[valIndex-1] = v
		} else {
			m.vals[valIndex] = v
		}
	} else {
		log.Infof("Map does not contain key: %s. Setting %s=%s", k, k, v)

		k1 := m.keys[0:keyIndex]
		k2 := []string{k}
		k3 := m.keys[keyIndex:]
		m.keys = append(k1, append(k2, k3...)...)

		v1 := m.vals[0:valIndex]
		v2 := []string{v}
		v3 := m.vals[valIndex:]
		m.vals = append(v1, append(v2, v3...)...)
	}
}

func (m *MyMap) Get(k string) string {
	i := sort.SearchStrings(m.keys, k)
	contained := len(m.keys) > i
	if contained {
		return m.vals[i]
	}
	return ""
}