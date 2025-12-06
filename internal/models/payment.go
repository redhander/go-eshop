package models

import (
	"github.com/redhander/go-eshop/internal/db/repository"
)

type PaymentModel struct {
	OrderID         string `json:"orderId" binding:"required,uuid"`
	PaymentMethodID string `json:"paymentMethodId" binding:"required,uuid"`
}

type UpdatePaymentStatusModel struct {
	Status repository.PaymentStatus `json:"status" binding:"required"`
}
