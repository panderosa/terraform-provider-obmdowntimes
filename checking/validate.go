package checking

import (
	"fmt"
	"time"
)

// List of Downtime Categories
func Categories() map[string]string {
	return map[string]string{
		"OTHER":            "1",
		"OS_CONFIGURATION": "2",
		"APP_MAINTENANCE":  "3",
		"APP_INSTALLATION": "4",
		"NW_MAINTENANCE":   "5",
		"HW_MAINTENANCE":   "6",
		"HW_INSTALLATION":  "7",
		"SECURITY":         "8",
	}
}

func DowntimeActions() []string {
	return []string{
		"REMINDER",
		"SUPPRESS_NOTIFICATIONS",
		"ENFORCE_ON_KPI_CALCULATION",
		"ENFORCE_ON_REPORTS",
		"STOP_MONITORING",
		"ENFORCE_ON_REPORTS",
	}
}

//
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

func ValidateCategory(v interface{}, k string) (warnings []string, errors []error) {
	category, ok := v.(string)

	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}
	keys := make([]string, 0, len(Categories()))
	for e := range Categories() {
		if category == e {
			return warnings, errors
		}
		keys = append(keys, e)
	}

	errors = append(errors, fmt.Errorf("expected %s to be one of %v, got %s", k, keys, category))
	return warnings, errors
}

func ValidateTimezone(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errors
	}

	_, err := time.LoadLocation(v)
	if err != nil {
		errors = append(errors, err)
	}
	return warnings, errors
}

/*func ValidateCategory(v interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	category := v.(string)

	for k := range Categories() {
		if category == k {
			return diags
		}
	}

	diag := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "wrong category",
		Detail:   fmt.Sprintf("%q is not in the list of valid categories", category),
	}
	diags = append(diags, diag)

	return diags
}*/
