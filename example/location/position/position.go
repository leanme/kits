package position

import (
	"context"
	"github.com/leanme/kits/example/location/model"
	"github.com/leanme/kits/service/restful/code"
)

func GetCitizenPosition(ctx context.Context, id string) (model.Location, code.Error) {
	return model.GetRedisSession().GetCitizenLocation(id)
}
