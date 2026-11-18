package models

import "time"

// Studio
type Studio struct {
	ID          string    `gorm:"primaryKey;type:uuid" json:"id"`
	OwnerID     string    `gorm:"type:uuid;index" json:"owner_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Type        string    `gorm:"type:varchar(50)" json:"type"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Photos      []Photo   `gorm:"foreignKey:StudioID" json:"photos,omitempty"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

	Resources []Resource `gorm:"foreignKey:StudioID" json:"resources,omitempty"`
}

// Resource — sala / stanowisko / sprzęt
type Resource struct {
	ID            string    `gorm:"primaryKey;type:uuid" json:"id"`
	StudioID      string    `gorm:"type:uuid;index" json:"studio_id"`
	Name          string    `gorm:"not null" json:"name"`
	Type          string    `gorm:"type:varchar(50)" json:"type"`
	Capacity      int       `json:"capacity"`
	PricePerHour  float64   `json:"price_per_hour"`
	Photos        []Photo   `gorm:"foreignKey:ResourceID" json:"photos,omitempty"`
	Description   string    `gorm:"type:text" json:"description"`
}

// Schedule — harmonogram tygodniowy
type Schedule struct {
	ID                  string    `gorm:"primaryKey;type:uuid" json:"id"`
	ResourceID          string    `gorm:"type:uuid;index" json:"resource_id"`
	DayOfWeek           int       `gorm:"not null" json:"day_of_week"` // 0=pon, 6=niedz
	OpenTime            string    `gorm:"type:varchar(5)" json:"open_time"`   // "09:00"
	CloseTime           string    `gorm:"type:varchar(5)" json:"close_time"`  // "18:00"
	SlotDurationMinutes int       `gorm:"default:60" json:"slot_duration_minutes"`
}

// ScheduleException — blokady / urlopy
type ScheduleException struct {
	ID                string    `gorm:"primaryKey;type:uuid" json:"id"`
	ResourceID        string    `gorm:"type:uuid;index" json:"resource_id"`
	DateFrom          time.Time `json:"date_from"`
	DateTo            time.Time `json:"date_to"`
	IsBlocked         bool      `json:"is_blocked"`
	OverrideOpenTime  *string   `gorm:"type:varchar(5)" json:"override_open_time,omitempty"`
	OverrideCloseTime *string   `gorm:"type:varchar(5)" json:"override_close_time,omitempty"`
	Reason            string    `gorm:"type:text" json:"reason"`
}

// Booking
type Booking struct {
	ID              string    `gorm:"primaryKey;type:uuid" json:"id"`
	ResourceID      string    `gorm:"type:uuid;index" json:"resource_id"`
	ClientID        string    `gorm:"type:uuid;index" json:"client_id"`
	StartAt         time.Time `gorm:"index" json:"start_at"`
	EndAt           time.Time `gorm:"index" json:"end_at"`
	Status          string    `gorm:"type:varchar(20);default:pending" json:"status"`
	TotalPrice      float64   `json:"total_price"`
	StripePaymentID string    `json:"stripe_payment_id,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`

	Resource Resource `gorm:"foreignKey:ResourceID" json:"resource,omitempty"`
}

// Review
type Review struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	BookingID string    `gorm:"type:uuid;index" json:"booking_id"`
	ClientID  string    `gorm:"type:uuid;index" json:"client_id"`
	StudioID  string    `gorm:"type:uuid;index" json:"studio_id"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Photo — wspólna struktura na zdjęcia
type Photo struct {
	ID         string    `gorm:"primaryKey;type:uuid" json:"id"`
	StudioID   string    `gorm:"type:uuid;index" json:"studio_id,omitempty"`
	ResourceID string    `gorm:"type:uuid;index" json:"resource_id,omitempty"`
	URL        string    `gorm:"not null" json:"url"`
	UploadedAt time.Time `gorm:"autoCreateTime" json:"uploaded_at"`
}

// User — mirror z Appwrite
type User struct {
	ID       string `gorm:"primaryKey;type:uuid" json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `gorm:"type:varchar(10);default:client" json:"role"` // "owner" | "client"
	Avatar   string `json:"avatar,omitempty"`
	Verified bool   `json:"verified"`
}

// TimeSlot — wygenerowany slot dostępności
type TimeSlot struct {
	ResourceID string    `json:"resource_id"`
	Date       string    `json:"date"`
	StartAt    time.Time `json:"start_at"`
	EndAt      time.Time `json:"end_at"`
	Available  bool      `json:"available"`
}