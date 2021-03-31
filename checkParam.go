// (C) 2021 GON Y YI.
// https://gonyyi.com/copyright.txt

package mutt

func CheckParamString(fields ...string) bool {
	for i := 0; i < len(fields); i++ {
		if fields[i] == "" {
			return false
		}
	}
	return true
}
