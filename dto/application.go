package dto

type ApplicationDTO struct {
	Name         string   `json:"name" bson:"name"`
	Icon         string   `json:"icon" bson:"icon"`
	URL          string   `json:"url" bson:"url"`
	Description  string   `json:"description" bson:"description"`
	RedirectURLs []string `json:"redirect_urls" bson:"redirect_urls"`
}
