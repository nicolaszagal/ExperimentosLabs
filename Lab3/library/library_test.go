package library

import "testing"

func TestAddBook(t *testing.T) {
	library := NewLibrary()
	library.AddBook("Caraval")

	if len(library.books) != 1 {
		t.Errorf("Se esperaba 1 libro, pero se encontraron %d", len(library.books))
	}

	if library.books["Caraval"] != 1 {
		t.Errorf("Se esperaba que 'Caraval' tuviera 1 unidad, pero se encontro '%d'", library.books["Caraval"])
	}
}

func TestIsBookAvailable(t *testing.T) {
	library := NewLibrary()
	library.AddBook("Legendary")

	if !library.IsBookAvailable("Legendary") {
		t.Errorf("Se esperaba que 'Legendary' estuviera disponible, pero no lo estaba")
	}

	if library.IsBookAvailable("Finale") {
		t.Errorf("Se esperaba que 'Finale' no estuviera disponible, pero si lo estaba")
	}
}

func TestRemoveBook(t *testing.T) {
	library := NewLibrary()
	library.AddBook("Finale")
	library.AddBook("Finale")

	if !library.RemoveBook("Finale") {
		t.Errorf("Se esperaba que 'Finale' fuera eliminado, pero no lo fue")
	}

	if library.IsBookAvailable("Finale") {
		t.Errorf("Se esperaba que 'Finale' no estuviera disponible, pero si lo estaba")
	}

	library.RemoveBook("Finale")
	if library.IsBookAvailable("Finale") {
		t.Errorf("Se esperaba que 'Finale' no estuviera disponible, pero si lo estaba")
	}
}

func TestLibraryIntegration(t *testing.T) {
	library := NewLibrary()

	// Agregar libros
	library.AddBook("Ciudad de Hueso")
	library.AddBook("Ciudad de Hueso")
	library.AddBook("Ciudad de Ceniza")

	// Verificar que los libros se agregaron correctamente
	if len(library.books) != 2 {
		t.Errorf("Se esperaban 2 libros, pero se encontraron %d", len(library.books))
	}

	if library.books["Ciudad de Hueso"] == 2 {
		t.Errorf("Se esperaba que 'Ciudad de Hueso' tuviera 2 unidades, pero tiene %d", library.books["El Alquimista"])
	}

	if library.books["Ciudad de Ceniza"] != 1 {
		t.Errorf("Se esperaba que 'Ciudad de Ceniza' tuviera 1 unidad, pero tiene %d", library.books["1984"])
	}

	// Verificar disponibilidad
	if !library.IsBookAvailable("Ciudad de Hueso") {
		t.Errorf("Se esperaba que 'Ciudad de Hueso' estuviera disponible, pero no lo estaba")
	}

	if library.IsBookAvailable("El Código Da Vinci") {
		t.Errorf("Se esperaba que 'El Código Da Vinci' no estuviera disponible, pero sí lo estaba")
	}

	// Eliminar libros
	if !library.RemoveBook("Ciudad de Hueso") {
		t.Errorf("Se esperaba que 'Ciudad de Hueso' fuera eliminado, pero no lo fue")
	}

	if library.books["Ciudad de Hueso"] == 1 {
		t.Errorf("Se esperaba que quedara 1 unidad de 'Ciudad de Hueso', pero tiene %d", library.books["Ciudad de Hueso"])
	}

	library.RemoveBook("Ciudad de Hueso")
	if library.IsBookAvailable("Ciudad de Hueso") {
		t.Errorf("Se esperaba que 'Ciudad de Hueso' no estuviera disponible, pero sí lo estaba")
	}

	// Verificar eliminación completa
	library.RemoveBook("Ciudad de Ceniza")
	if library.IsBookAvailable("Ciudad de Ceniza") {
		t.Errorf("Se esperaba que 'Ciudad de Ceniza' no estuviera disponible, pero sí lo estaba")
	}

	if len(library.books) != 0 {
		t.Errorf("Se esperaba que no quedaran libros, pero se encontraron %d", len(library.books))
	}
}
