package domain

type User struct {
    ID   int    `json:"id" gorm:"primary_key"`
    Name string `json:"name" gorm:""`
      Agis int `json:"age" gorm:""`
}


