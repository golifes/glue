package auth

import (
	"github.com/xwinie/glue/core"
)

func findDictByTypeService(t string) (responseEntity core.ResponseEntity) {
	u, err := getDictByType(t)
	if err != nil {
		return *responseEntity.BuildError(core.BuildEntity(QueryError, getMsg(QueryError)))
	}
	return *responseEntity.Build(u)
}
