package model

// Topic is a discussion topic.
type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NameCode    string `json:"name_code"`
	Description string `json:"description"`
}

const (
	minLen = 1
	maxLen = 100
)

func NewCategory(name, nameCode, description string) (*Category, error) {
	category := &Category{
		Name:        name,
		NameCode:    nameCode,
		Description: description,
	}
	// Validate min len
	if err := category.Validate(); err != nil {
		return category, err
	}

	return category, nil
}

// ValidTopicTitle checks if Category data is valid.
func (p *Category) Validate() error {
	return nil
}
