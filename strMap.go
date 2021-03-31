package mutt

func NewStrMap() StrMap {
	return make(StrMap)
}

type StrMap map[string]string

func (m StrMap) Add(key, val string) {
	m[key] = val
}

func (m StrMap) Get(key string) (string, bool) {
	if v, ok := m[key]; ok {
		return v, true
	}
	return "", false
}

func (m StrMap) Del(key string) {
	delete(m, key)
}

func (m StrMap) Reset() {
	for k, _ := range m {
		delete(m, k)
	}
}
