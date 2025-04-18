import com.library.app.Library;
import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.*;

public class TestLibrary {
    private Library library;

    @Before
    public void setUp() {
        library = new Library();
    }

    @Test
    public void testAddBook(){
        library.addBook("Caraval");
        assertTrue(library.isBookAvailable("Caraval"));
    }

    @Test
    public void testIsBookAvailable_BookNotAdded() {
        assertFalse(library.isBookAvailable("Legendary"));
    }

    @Test
    public void testAddDuplicateBook() {
        library.addBook("Finale");
        library.addBook("Finale");
        assertTrue(library.isBookAvailable("Finale"));
    }
}
