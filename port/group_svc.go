package port

import (
	"context"

	"github.com/w-woong/product/dto"
)

type GroupSvc interface {
	ReadGroup(ctx context.Context, id string) (dto.Group, error)
}
