package library

// Library representa una libreria que almacena libros
type Library struct {
	books map[string]int
}

func NewLibrary() *Library {
	return &Library{books: make(map[string]int)}
}

// AddBook agrega un libro a la libreria
func (l *Library) AddBook(bookTitle string) {
	if l.IsBookAvailable(bookTitle) {
		l.books[bookTitle]++
	}
	l.books[bookTitle] = 1
}

// IsBookAvailable verifica si un libro esta disponible en la libreria
func (l *Library) IsBookAvailable(bookTitle string) bool {
	return l.books[bookTitle] > 0
}

// RemoveBook elimina un libro si esta disponible en la libreria
func (l *Library) RemoveBook(bookTitle string) bool {
	if l.IsBookAvailable(bookTitle) {
		l.books[bookTitle]-- //Elimina una unidad del libro
		if l.books[bookTitle] == 0 {
			delete(l.books, bookTitle) //Elimina el libro si no hay más unidades
		}
		return true
	}
	return false
}
