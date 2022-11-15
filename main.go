package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// `id` int NOT NULL AUTO_INCREMENT,
// `title` varchar(100) NOT NULL,
// `content` text,
// `image` json DEFAULT NULL,
// `has_finished` tinyint(1) DEFAULT '0',
// `status` int DEFAULT '1',
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

type NoteUpdate struct {
	Title *string `json:"title" gorm:"column:title;"`
}

func (Note) Tablename() string {
	return "notes"
}

func main() {
	dsn := "dbadmin:dbadmin@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// INSERT
	// newNote := Note{Title: "Title Here", Content: "Content Here"}
	// if err := db.Create(&newNote); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(newNote)

	// SELECT
	// var notes []Note
	// db.Where("status =?", 1).Find(&notes)
	// fmt.Println(notes)

	// var note Note
	// if err := db.Where("id = 3").First(&note); err != nil {
	// 	log.Println(err) // log print # fmt print | show time vs note show time -> log print more controlable
	// }
	// fmt.Println(note)

	// DELETE -> we need Tablename | but actually we have to use Update instead of Delete for more safe
	// db.Table(Note{}.Tablename()).Where("id = 9").Delete(nil)

	// SOFT DELETE 1 = UPDATE 1
	// db.Table(Note{}.Tablename()).Where("id = 12").Updates(map[string]interface{}{
	// 	"title":   "New Title",
	// 	"content": "New Content",
	// })

	//  SOFT DELETE 2 = UPDATE 2: use another struct for update in API
	var note Note
	if err := db.Where("id = 11").First(&note); err != nil {
		log.Println(err) // log print # fmt print | show time vs note show time -> log print more controlable
	}

	newTitle := "Updated Title"
	db.Table(Note{}.Tablename()).Where("id = 11").Updates(&NoteUpdate{Title: &newTitle})
	fmt.Println(note)
}
