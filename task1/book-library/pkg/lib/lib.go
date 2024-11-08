package lib

import (
	"fmt"
	"math/rand"
)

// Структура книги
type Book struct {
	ID     string
	Title  string
	Author string
}

// Интерфейс для поиска книги по ID
type Search interface {
	SearchByID(id string) (Book, error)
}

// Интерфейс для поиска книг по имени
type Library interface {
	SearchByName(name string) (Book, error)
	AddBook(book Book) string
	ReplaceIDGenerator(generator func() string)
}

// Библиотека
type LibraryImpl struct {
	books       map[string]Book
	idGenerator func() string
	nameToID    map[string]string // Мапа для связи имени с ID книги
}

// Конструктор библиотеки
func CreateLibrary(idGenerator func() string) *LibraryImpl {
	return &LibraryImpl{
		books:       make(map[string]Book),
		idGenerator: idGenerator,
		nameToID:    make(map[string]string),
	}
}

// Очистка хранилища книг
func (lib *LibraryImpl) ClearStorage() {
	lib.books = make(map[string]Book)
	lib.nameToID = make(map[string]string)
}

// Метод добавления книги в библиотеку
func (lib *LibraryImpl) AddBook(book Book) string {
	id := lib.idGenerator()
	book.ID = id
	lib.books[id], lib.nameToID[book.Title] = book, id
	return id
}

// Метод поиска книги по имени
func (lib *LibraryImpl) SearchByName(name string) (Book, error) {
	if id, found := lib.nameToID[name]; found {
		return lib.SearchByID(id)
	}
	return Book{}, fmt.Errorf("Книга с названием '%s' не найдена", name)
}

// Метод поиска книги по ID
func (lib *LibraryImpl) SearchByID(id string) (Book, error) {
	if book, found := lib.books[id]; found {
		return book, nil
	}
	return Book{}, fmt.Errorf("Книга с ID '%s' не найдена", id)
}

// Метод для замены генератора ID
func (lib *LibraryImpl) ReplaceIDGenerator(generator func() string) {
	lib.idGenerator = generator
}

// Функция генерации ID
func IDGenerator() string {
	return fmt.Sprint(rand.Int())
}

// Новая функция генерации ID
func NewIDGenerator() string {
	return fmt.Sprint(rand.Int() + 1)
}
