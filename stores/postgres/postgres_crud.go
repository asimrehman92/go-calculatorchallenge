package stores

import (
	"fmt"
	"time"
)

func CreateClient(uname, eq string) {

	client := Clients{
		Name:      uname,
		Equation:  eq,
		Timestamp: time.Now(),
	}

	db.Create(&client)
}

func UpdateClient(uname, eq string) {
	fmt.Println("I am in update func...")

	client := Clients{
		Name:      uname,
		Equation:  eq,
		Timestamp: time.Now(),
	}

	// db.Where("name <> ?", uname).Find(&client).Update("equation", eq)
	db.Where(&Clients{Name: uname}).Find(&client).Update("equation", eq)
}

// func main() {
// 	fmt.Println("Program started...")
// 	DataMigration()
// 	// UpdateClient("tabish", "1 + 1 + 3 = 5")
// 	CreateClient("Sidra", "5 * 4 = 20")
// }
