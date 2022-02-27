package checking

import "fmt"

func IsOnce(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	if v != "ONCE" {
		errors = append(errors, fmt.Errorf("expected %q to be a \"ONCE\", got %q", k, v))
	}
	return warnings, errors
}
