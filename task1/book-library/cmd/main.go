package main

import (
	"book-library/pkg/lib"
	"fmt"
)

// Демонстрация сценария
func main() {
	// Создаем слайс книг
	data := []lib.Book{
		{Title: "Крутой маршрут", Author: "Евгения Гинзбург"},
		{Title: "На краю света", Author: "Сергей Безбородов"},
	}

	// Создаем библиотеку
	library := lib.CreateLibrary(lib.IDGenerator)

	// Загружаем книги в библиотеку
	for _, book := range data {
		library.AddBook(book)
	}

	// Ищем книгу по имени
	book, err := library.SearchByName("Крутой маршрут")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Найдена книга: %+v\n", book)
	}

	// Меняем генератор ID
	library.ReplaceIDGenerator(lib.NewIDGenerator)

	// Добавляем новую книгу
	newBook := lib.Book{Title: "Пикник на обочине", Author: "Аркадий и Борис Стругацкие"}
	library.AddBook(newBook)

	// Ищем добавленную книгу по имени
	book, err = library.SearchByName("Пикник на обочине")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Найдена книга: %+v\n", book)
	}

	// Заменяем хранилище
	// Предварительно очищаем
	library.ClearStorage()

	// Создаем новый слайс книг
	new_data := []lib.Book{
		{Title: "Педагогическая поэма", Author: "Антон Макаренко "},
		{Title: "Последнее желание", Author: "Анджей Сапковский"},
	}

	// Загружаем книги в библиотеку
	for _, book := range new_data {
		library.AddBook(book)
	}

	// Ищем добавленную книгу по имени
	book, err = library.SearchByName("Последнее желание")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Найдена книга: %+v\n", book)
	}

}
