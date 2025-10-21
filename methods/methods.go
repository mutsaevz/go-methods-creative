package methods

import "fmt"

type Library struct {
	NameLibrary  []string // Название библиотеки | №1
	Address      []string // Адрес | №2
	FoundedYear  int      // Год основания | №3
	BooksCount   int      // Общее количество книг | №4
	Librarian    string   // Имя главного библиотекаря | №5
	Sections     []string // Разделы (например, "Философия", "История", "Наука") | №6
	ReadersCount int      // Количество зарегистрированных читателей | №7
	OpenHours    string   // Время работы (например, "9:00–18:00") | №8
	PhoneNumber  string   // Контактный номер | №9
	IsOpen       bool     // Флаг, открыта ли библиотека сейчас | №10
}

type LibraryFull struct {
	libraryFull []Library
}

func (l *Library) FindLibrary(name string) string { // | №1
	for i := 0; i < len(l.NameLibrary); i++ {
		if l.NameLibrary[i] == name {
			return fmt.Sprintf("Такая библиотека нахоится по адресу", l.Address)
		}
	}
	return "Библиотека с таким названием не найдена"
}

func (l *Library) FindAddress(address string) string {
	for i := 0; i < len(l.Address); i++ {
		if l.Address[i] == address {
			return fmt.Sprintf("По этому адресу находится библитека", l.NameLibrary)
		}
	}
	return "По такому адресу не найдена ни одна библиотека"
}
