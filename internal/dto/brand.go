package dto

import (
	"github.com/redhander/go-eshop/internal/db/repository"
)

type BrandDetail struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	ImageUrl     *string `json:"imageUrl,omitempty"`
	ImageID      *string `json:"imageId,omitempty"`
	Description  *string `json:"description,omitempty"`
	Slug         string  `json:"slug"`
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	Published    bool    `json:"published"`
}

func MapBrandDetail(brand repository.Brand) BrandDetail {
	return BrandDetail{
		ID:           brand.ID.String(),
		Name:         brand.Name,
		ImageUrl:     brand.ImageUrl,
		ImageID:      brand.ImageID,
		Description:  brand.Description,
		Slug:         brand.Slug,
		DisplayOrder: brand.DisplayOrder,
		Published:    brand.Published,
	}
}
