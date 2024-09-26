package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func runQuery(query string) *sql.Rows {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	//defer result.Close()
	return result
}

func GetStartedManga(user string) ([]button, string) {
	var query = os.Getenv("SELECT_STARTED_MANGA")
	query += user
	fmt.Println(query)
	//fmt.Println(query)
	//var query = "SELECT man.eng_name || '/' || man.rus_name as name, man.code_name FROM \"Manga\" man\nINNER JOIN \"Users_x_Manga\" uxm\nON uxm.manga_id = man.id\nAND uxm.status in ('started', 'waiting')\nAND uxm.user_id = " + user

	result := runQuery(query)
	defer result.Close()

	var r []button //TODO make
	var s string
	for result.Next() {
		var name, codeName string
		err := result.Scan(&name, &codeName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r = append(r, button{name, fmt.Sprintf("%v/MANGA/read/%v", user, codeName)})
	}
	if len(r) == 0 {
		s = "У вас нет начатой манги."
	} else {
		s = "Вот манга, которую Вы уже начали читать:"
	}
	return r, s
}

func GetNewManga(user string) ([]button, string) {
	var query, _ = os.LookupEnv("SELECT_NEW_MANGA")
	fmt.Println(query)
	query += user
	fmt.Println(query)
	//var query = "SELECT man.eng_name || '/' || man.rus_name as name, man.code_name\nFROM \"Manga\" man\nLEFT JOIN \"Users_x_Manga\" uxm\nON uxm.manga_id = man.id\nAND uxm.status not in ('started', 'waiting')\nAND uxm.user_id = " + user

	result := runQuery(query)
	var r []button //TODO make
	var s string
	for result.Next() {
		var name, codeName string
		err := result.Scan(&name, &codeName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r = append(r, button{name, fmt.Sprintf("%v/MANGA/read/%v", user, codeName)})
	}
	if len(r) == 0 {
		s = "Нет такой манги, которую Вы не прочитали."
	} else {
		s = "Вот манга, которую Вы еще не читали:"
	}
	return r, s
}

func Read(user string, manga string) ([]button, string) {
	query, _ := os.LookupEnv("SELECT_STATUS_USER_X_MANGA")
	result := runQuery(fmt.Sprintf(query, user, manga))
	//result := runQuery(fmt.Sprintf("SELECT uxm.status as status\nFROM \"Manga\" man\nINNER JOIN \"Users_x_Manga\" uxm\nON uxm.manga_id = man.id\nAND uxm.status = 'started'\nAND uxm.user_id = %v\nAND man.code_name = '%v' LIMIT 1", user, manga))
	var flg = ""
	for result.Next() {
		_ = result.Scan(&flg)
	}
	fmt.Println(flg)
	switch flg {
	case "":
		{
			//TODO ЗАПИХНУТЬ ВСЕ SQL-ЗАПРОСЫ В ПЕРЕМЕННЫЕ
			query, _ = os.LookupEnv("INSERT_USER_X_MANGA")
			result = runQuery(fmt.Sprintf(query, user, manga))
			//result = runQuery(fmt.Sprintf("insert into \"Users_x_Manga\" values (%v, (select id from \"Manga\" where code_name = '%v'), 'started', 1)", user, manga))
			query, _ = os.LookupEnv("SELECT_INFO_CURRENT_CHAPTER")
			result = runQuery(fmt.Sprintf(query, manga, manga, user))
			//result = runQuery(fmt.Sprintf("select number, eng_name, rus_name, telegraph from \"Chapters\"\nwhere manga_id = (select id from \"Manga\" where code_name = '%v')\nand serial_number = (select cur_chapt from \"Users_x_Manga\" where manga_id = (select id from \"Manga\" where code_name = '%v') and user_id = %v limit 1)", manga, manga, user))
			var number, eng, rus, telegraph string
			for result.Next() {
				_ = result.Scan(&number, &eng, &rus, &telegraph)
			}
			return []button{
					{"Следующая глава", fmt.Sprintf("%v/MANGA/read/%v", user, manga)},
					{"Главное меню", user + "/MANGA/menu/"},
				},
				fmt.Sprintf("%v\nГлава #%v\n%v\n%v\nTelegraph: %v", manga, number, eng, rus, telegraph)
		}
	case "started":
		{
			query, _ = os.LookupEnv("SELECT_NUMBER_CURRENT_CHAPTER")
			result = runQuery(fmt.Sprintf(query, user, manga))
			//result = runQuery(fmt.Sprintf("select cur_chapt from \"Users_x_Manga\" where user_id = %v and manga_id = (select id from \"Manga\" where code_name = '%v') limit 1", user, manga))
			var curChapter float64
			for result.Next() {
				_ = result.Scan(&curChapter)
			}
			query, _ = os.LookupEnv("SELECT_NUMBER_LAST_CHAPTER")
			result = runQuery(fmt.Sprintf(query, manga))
			//result = runQuery(fmt.Sprintf("select max(serial_number) from \"Chapters\" where manga_id = (select id from \"Manga\" where code_name = '%v')", manga))
			var maxChapter float64
			for result.Next() {
				_ = result.Scan(&maxChapter)
			}
			if curChapter == maxChapter {
				query, _ = os.LookupEnv("SELECT_STATUS_MANGA")
				result = runQuery(fmt.Sprintf(query, manga))
				//result = runQuery(fmt.Sprintf("select status from \"Manga\" where code_name = '%v'", manga))
				var mangaStatus string
				for result.Next() {
					_ = result.Scan(&mangaStatus)
				}
				if mangaStatus == "ongoing" {
					query, _ = os.LookupEnv("UPDATE_USER_X_MANGA_WAITING")
					result = runQuery(fmt.Sprintf(query, user, manga))
					//result = runQuery(fmt.Sprintf("update \"Users_x_Manga\" set status = 'waiting' where user_id = %v and manga_id = (select id from \"Manga\" where code_name = '%v')", user, manga))
					return []button{
							{"Главное меню", user + "/MANGA/menu/"},
						},
						"Манга продолжается, но новые главы пока не вышли. Приходите позже!"
				} else {
					query, _ = os.LookupEnv("UPDATE_USER_X_MANGA_FINISHED")
					result = runQuery(fmt.Sprintf(query, user, manga))
					//result = runQuery(fmt.Sprintf("update \"Users_x_Manga\" set status = 'finished' where user_id = %v and manga_id = (select id from \"Manga\" where code_name = '%v')", user, manga))
					return []button{
							{"Главное меню", user + "/MANGA/menu/"},
						},
						"Вы закончили читать мангу!"
				}
			} else {
				// TODO нужно сделать не просто +1, а +1 от кнопки
				query, _ = os.LookupEnv("UPDATE_CURRENT_CHAPTER")
				result = runQuery(fmt.Sprintf(query, user, manga))
				//result = runQuery(fmt.Sprintf("update \"Users_x_Manga\" set cur_chapt = cur_chapt + 1  where user_id = %v and manga_id = (select id from \"Manga\" where code_name = '%v')", user, manga))
				query, _ = os.LookupEnv("SELECT_INFO_CURRENT_CHAPTER")
				result = runQuery(fmt.Sprintf(query, manga, manga, user))
				//result = runQuery(fmt.Sprintf("select number, eng_name, rus_name, telegraph from \"Chapters\"\nwhere manga_id = (select id from \"Manga\" where code_name = '%v')\nand serial_number = (select cur_chapt from \"Users_x_Manga\" where manga_id = (select id from \"Manga\" where code_name = '%v') and user_id = %v limit 1)", manga, manga, user))
				var number, eng, rus, telegraph string
				for result.Next() {
					_ = result.Scan(&number, &eng, &rus, &telegraph)
				}
				return []button{
						{"Следующая глава", fmt.Sprintf("%v/MANGA/read/%v", user, manga)},
						{"Главное меню", user + "/MANGA/menu/"},
					},
					fmt.Sprintf("%v\nГлава #%v\n%v\n%v\nTelegraph: %v", manga, number, eng, rus, telegraph)
			}
		}
	case "finished":
		{
			return []button{{"Главное меню", user + "/MANGA/menu/"}},
				"Вы закончили читать мангу!"
		}
	case "waiting":
		{
			return []button{{"Главное меню", user + "/MANGA/menu/"}},
				"Манга продолжается, но новые главы пока не вышли. Приходите позже!"
		}
	}
	return []button{{"1", "1"}}, "1"
}
