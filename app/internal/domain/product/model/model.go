package model

import (
	"time"
)

type Product struct {
	ID            string                 `mapstructure:"id"`
	Name          string                 `mapstructure:"name"`
	Description   string                 `mapstructure:"description"`
	ImageID       *string                `mapstructure:"image_id"`
	Price         uint64                 `mapstructure:"price"`
	CurrencyID    uint32                 `mapstructure:"currency_id"`
	Rating        uint32                 `mapstructure:"rating"`
	CategoryID    uint32                 `mapstructure:"category_id"`
	Specification map[string]interface{} `mapstructure:"specification"`
	CreatedAt     time.Time              `mapstructure:"created_at"`
	UpdatedAt     *time.Time             `mapstructure:"updated_at"`
}
