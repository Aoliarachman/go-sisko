package service

import (
	"context"
	"go-sisko/model/web"
)

type ElementaryService interface {
	Create(ctx context.Context, request web.ElementaryCreateRequest) web.ElementaryResponse
	Update(ctx context.Context, request web.ElementaryUpdateRequest) web.ElementaryResponse
	Delete(ctx context.Context, elementaryId int)
	FindById(ctx context.Context, elementaryId int) web.ElementaryResponse
	FindAll(ctx context.Context) []web.ElementaryResponse
}
