package gorsk

// Company represents company model
type Company struct {
	Base
	Name      string     `json:"name"`
	Active    bool       `json:"active" gorm:"default:true"`
	Locations []Location `json:"locations,omitempty"`
	Owner     User       `json:"owner"`
}
