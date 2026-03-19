package usecase

import "atlas.chr/vaa/hw/cart/internal/entity"

func ProductToEntity(dto *ProductDTO) *entity.Product {
	return &entity.Product{
		Sku:   dto.Sku,
		Name:  dto.Name,
		Price: dto.Price,
	}
}
