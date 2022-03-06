package downtime

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
