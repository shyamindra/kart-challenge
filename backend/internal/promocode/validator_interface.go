package promocode

// PromoCodeValidator defines the interface for promo code validation.
type PromoCodeValidator interface {
	IsValid(code string) bool
}
