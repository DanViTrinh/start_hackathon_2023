package structs


type Recipe struct {
	Id int `json:"id"`
	Name string `json:"name"`
	MealTime string `json:"mealTime"`
	Information string `json:"information"`
	Image string `json:"image"`

}