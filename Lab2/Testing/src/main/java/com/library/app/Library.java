package com.library.app;

import java.util.ArrayList;
import java.util.List;

public class Library {
    private List<String> books;

    public Library() {
        this.books = new ArrayList<>();
    }

    // Metodo para agregar un libro a la librería
    public void addBook(String bookTitle) {
        books.add(bookTitle);
    }

    // Metodo para verificar si un libro está disponible
    public boolean isBookAvailable(String bookTitle) {
        return books.contains(bookTitle);
    }
}