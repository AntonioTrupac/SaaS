// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type Actor interface {
	IsActor()
}

type Node interface {
	IsNode()
}

type Address struct {
	ID          int    `json:"id"`
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	PostalCode  int    `json:"postalCode"`
	Country     string `json:"country"`
	UserID      int    `json:"userId"`
}

type AddressInput struct {
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	PostalCode  int    `json:"postalCode"`
	Country     string `json:"country"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryInput struct {
	Name string `json:"name"`
}

type HabitSettings struct {
	ID            int `json:"id"`
	Complete      int `json:"complete"`
	Failed        int `json:"failed"`
	Skipped       int `json:"skipped"`
	Total         int `json:"total"`
	CurrentStreak int `json:"currentStreak"`
	HabitID       int `json:"habitId"`
}

type HabitSettingsInput struct {
	Complete      int `json:"complete"`
	Failed        int `json:"failed"`
	Skipped       int `json:"skipped"`
	Total         int `json:"total"`
	CurrentStreak int `json:"currentStreak"`
}

type Habits struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	UserID       string         `json:"userId"`
	HabitSetting *HabitSettings `json:"habitSetting"`
}

type HabitsInput struct {
	Name         string              `json:"name"`
	HabitSetting *HabitSettingsInput `json:"habitSetting"`
}

type Image struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	ProductID int    `json:"productId"`
}

type ImageInput struct {
	URL string `json:"url"`
}

type Moods struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Notes  string `json:"notes"`
	UserID int    `json:"userId"`
}

type MoodsInput struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Rating      float64     `json:"rating"`
	Stock       int         `json:"stock"`
	Image       []*Image    `json:"image"`
	Category    []*Category `json:"category"`
}

type ProductInput struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Rating      float64          `json:"rating"`
	Stock       int              `json:"stock"`
	Images      []*ImageInput    `json:"images"`
	Categories  []*CategoryInput `json:"categories"`
}

type User struct {
	ID        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Address   []*Address `json:"address"`
	Moods     *Moods     `json:"moods"`
	Habits    []*Habits  `json:"habits"`
}

func (User) IsActor() {}
func (User) IsNode()  {}

type UserInput struct {
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Age       int             `json:"age"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Address   []*AddressInput `json:"address"`
	Moods     *MoodsInput     `json:"moods"`
	Habits    []*HabitsInput  `json:"habits"`
}
