// (C) 2021 GON Y YI.
// https://gonyyi.com/copyright.txt

package mutt

func CheckParamString(fields ...string) error {
	for i := 0; i < len(fields); i++ {
		if fields[i] == "" {
			return ERR_MISSING_REQUIRED_FIELDS
		}
	}
	return nil
}
