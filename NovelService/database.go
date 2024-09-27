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

func GetStartedNovel(user string) ([]button, string) {
	var query = os.Getenv("SELECT_STARTED_NOVEL")
	query += user
	fmt.Println(query)

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
		r = append(r, button{name, fmt.Sprintf("%v/NOVEL/read/%v", user, codeName)})
	}
	if len(r) == 0 {
		s = "У вас нет начатой новеллы."
	} else {
		s = "Вот новеллы, которые Вы уже начали читать:"
	}
	return r, s
}

func GetNewNovel(user string) ([]button, string) {
	var query, _ = os.LookupEnv("SELECT_NEW_NOVEL")
	fmt.Println(query)
	query += user
	fmt.Println(query)

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
		r = append(r, button{name, fmt.Sprintf("%v/NOVEL/read/%v", user, codeName)})
	}
	if len(r) == 0 {
		s = "Нет такой новеллы, которую Вы не прочитали."
	} else {
		s = "Вот новеллы, которые Вы еще не читали:"
	}
	return r, s
}

func Read(user string, novel string) ([]button, string) {
	query, _ := os.LookupEnv("SELECT_STATUS_USER_X_NOVEL")
	result := runQuery(fmt.Sprintf(query, user, novel))
	var flg = ""
	for result.Next() {
		_ = result.Scan(&flg)
	}
	fmt.Println(flg)
	switch flg {
	case "":
		{
			//TODO ЗАПИХНУТЬ ВСЕ SQL-ЗАПРОСЫ В ПЕРЕМЕННЫЕ
			query, _ = os.LookupEnv("INSERT_USER_X_NOVEL")
			result = runQuery(fmt.Sprintf(query, user, novel))
			query, _ = os.LookupEnv("SELECT_INFO_CURRENT_CHAPTER")
			result = runQuery(fmt.Sprintf(query, novel, novel, user))
			var number, eng, rus, telegraph string
			for result.Next() {
				_ = result.Scan(&number, &eng, &rus, &telegraph)
			}
			fmt.Println("TGH: " + telegraph)
			return []button{
					{"Следующая глава", fmt.Sprintf("%v/NOVEL/read/%v", user, novel)},
					{"Главное меню", user + "/NOVEL/menu/"},
				},
				fmt.Sprintf("%v\nГлава #%v\n%v\n%v\nTelegraph: %v", novel, number, eng, rus, telegraph)
		}
	case "started":
		{
			query, _ = os.LookupEnv("SELECT_NUMBER_CURRENT_CHAPTER")
			result = runQuery(fmt.Sprintf(query, user, novel))
			var curChapter float64
			for result.Next() {
				_ = result.Scan(&curChapter)
			}
			query, _ = os.LookupEnv("SELECT_NUMBER_LAST_CHAPTER")
			result = runQuery(fmt.Sprintf(query, novel))
			var maxChapter float64
			for result.Next() {
				_ = result.Scan(&maxChapter)
			}
			if curChapter == maxChapter {
				query, _ = os.LookupEnv("SELECT_STATUS_NOVEL")
				result = runQuery(fmt.Sprintf(query, novel))
				var novelStatus string
				for result.Next() {
					_ = result.Scan(&novelStatus)
				}
				if novelStatus == "ongoing" {
					query, _ = os.LookupEnv("UPDATE_USER_X_NOVEL_WAITING")
					fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
					result = runQuery(fmt.Sprintf(query, user, novel))
					return []button{
							{"Главное меню", user + "/NOVEL/menu/"},
						},
						"Манга продолжается, но новые главы пока не вышли. Приходите позже!"
				} else {
					query, _ = os.LookupEnv("UPDATE_USER_X_NOVEL_FINISHED")
					result = runQuery(fmt.Sprintf(query, user, novel))
					return []button{
							{"Главное меню", user + "/NOVEL/menu/"},
						},
						"Вы закончили читать мангу!"
				}
			} else {
				// TODO нужно сделать не просто +1, а +1 от кнопки
				query, _ = os.LookupEnv("UPDATE_CURRENT_CHAPTER")
				result = runQuery(fmt.Sprintf(query, user, novel))
				query, _ = os.LookupEnv("SELECT_INFO_CURRENT_CHAPTER")
				result = runQuery(fmt.Sprintf(query, novel, novel, user))
				var number, eng, rus, telegraph string
				fmt.Println(result)
				for result.Next() {
					err := result.Scan(&number, &eng, &rus, &telegraph)
					if err != nil {
						fmt.Println(err)
					}
				}
				fmt.Println("TGH: !" + telegraph + "!")
				return []button{
						{"Следующая глава", fmt.Sprintf("%v/NOVEL/read/%v", user, novel)},
						{"Главное меню", user + "/NOVEL/menu/"},
					},
					fmt.Sprintf("%v\nГлава #%v\n%v\n%v\nTelegraph: %v", novel, number, eng, rus, telegraph)
			}
		}
	case "finished":
		{
			return []button{{"Главное меню", user + "/NOVEL/menu/"}},
				"Вы закончили читать мангу!"
		}
	case "waiting":
		{
			return []button{{"Главное меню", user + "/NOVEL/menu/"}},
				"Манга продолжается, но новые главы пока не вышли. Приходите позже!"
		}
	}
	return []button{{"1", "1"}}, "1"
}
