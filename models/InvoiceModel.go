package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceModel struct {
	ID               primitive.ObjectID `bson:"_id"`
	InvoiceId        string             `json:"invoice_id"`
	OrderId          string             `json:"order_id"`
	Payment_Method   *string            `json:"payment_method validate:"required,eq=CARD|eq=CASH"`
	Payment_Status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
