package mutt

func NewStringMap() StringMap {
	return make(StringMap)
}

type StringMap map[string]string

func (t StringMap) Add(key, val string) {
	t[key] = val
}

func (t StringMap) Get(key string) (string, bool) {
	if v, ok := t[key]; ok {
		return v, true
	}
	return "", false
}

func (t StringMap) Del(key string) {
	delete(t, key)
}

func (t StringMap) Reset() {
	for k, _ := range t {
		delete(t, k)
	}
}
