package obmsdk

import "regexp"

var reStringID = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func validStringID(v *string) bool {
	return v != nil && reStringID.MatchString(*v)
}
