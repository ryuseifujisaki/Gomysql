package domain

type User struct {
    ID   int    `json:"id" gorm:"primary_key"`
    Name string `json:"name" gorm:""`
    Age int `json:"age" gorm:""`
}

fmt.Println(veo)


