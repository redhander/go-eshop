package dto

import (
	"time"

	"github.com/google/uuid"
)

type CartDetail struct {
	ID             uuid.UUID        `json:"id"`
	TotalPrice     float64          `json:"totalPrice"`
	DiscountAmount float64          `json:"discountAmount"`
	UpdatedAt      *time.Time       `json:"updatedAt"`
	CreatedAt      time.Time        `json:"createdAt"`
	CartItems      []CartItemDetail `json:"cartItems"`
}

type CartDiscount struct {
	ID                string     `json:"id"`
	Code              string     `json:"code"`
	Description       *string    `json:"description,omitempty"`
	DiscountType      string     `json:"discountType"`
	DiscountValue     float64    `json:"discountValue"`
	MinOrderValue     *float64   `json:"minOrderValue,omitempty"`
	MaxDiscountAmount *float64   `json:"maxDiscountAmount,omitempty"`
	ValidFrom         time.Time  `json:"validFrom"`
	ValidUntil        *time.Time `json:"validUntil,omitempty"`
	IsStackable       bool       `json:"isStackable"`
}
