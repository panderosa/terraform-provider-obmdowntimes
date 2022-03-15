package downtime

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/panderosa/obmprovider/obmsdk"
)

func GenerateIdByUuid() (*string, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	id = strings.Replace(id, "-", "_", -1)
	id = strings.ToLower(id)
	return &id, nil
}

func GenerateIdByHash(ids []string) string {
	var id string
	if len(ids) > 0 {
		id = strings.Join(ids, "")
	} else {

	}
	id = fmt.Sprintf("%d", schema.HashString(id))
	return id
}

func Flatten2CIs(data []obmsdk.Ci) []interface{} {
	array := make([]string, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return []interface{}{array}
}
func Flatten3CIs(data []obmsdk.Ci) []interface{} {
	array := make([]interface{}, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return array
}

func Flatten4CIs(data []obmsdk.Ci) []interface{} {
	array := make([]interface{}, 0, len(data))
	for i := range data {
		array = append(array, data[i].ID)
	}
	return array
}
