package template

import (
	"github.com/ervinismu/go-server-wall/models"
)

// Struct for generate Response error or Success
func Response(res *models.Response) map[string]interface{}{
	data := map[string]interface{}{
		"code": res.Code,
		"message": res.Message,
		"token": res.Token,
		"status": res.Status,
	}
	return data
}