package model

type Tag struct {
	//*Model
	Name       string `json:"name"`
	State      uint8  `json:"state"`
	ModifiedBy string `json:"modifiedBy"`
	VreatedBy  string `json:"createdBy"`
}

/*func (t Tag) TableName() string {
	return "blog_tag"
}*/
