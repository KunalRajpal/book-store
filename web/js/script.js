document.addEventListener('DOMContentLoaded', fetchBooks);
document.getElementById('book-form').addEventListener('submit', submitBookForm);

function fetchBooks() {
    fetch('/book')
        .then(response => response.json())
        .then(books => displayBooks(books))
        .catch(error => console.error('Error:', error));
}

function displayBooks(books) {
    const booksDiv = document.getElementById('books');
    booksDiv.innerHTML = ''; // Clear current books
    books.forEach(book => {
        const bookElement = document.createElement('div');
        bookElement.innerHTML = `
            <h3>${book.Name}</h3>
            <p>Author: ${book.Author}</p>
            <p>Publication: ${book.Publication}</p>
            <button onclick="deleteBook(${book.ID})">Delete</button>
            <button onclick="populateForm(${book.ID})">Edit</button>
        `;
        booksDiv.appendChild(bookElement);
    });
}

function submitBookForm(event) {
    event.preventDefault();
    const bookId = document.getElementById('bookId').value;
    const bookData = {
        Name: document.getElementById('name').value,
        Author: document.getElementById('author').value,
        Publication: document.getElementById('publication').value
    };

    const method = bookId ? 'PUT' : 'POST';
    const endpoint = bookId ? `/book/${bookId}` : '/book';

    fetch(endpoint, {
        method: method,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookData)
    })
    .then(response => response.json())
    .then(data => {
        fetchBooks();
        resetForm();
        showMessage('Book saved successfully');
    })
    .catch(error => showMessage('Error saving book', true));
}

function deleteBook(bookId) {
    fetch(`/book/${bookId}`, { method: 'DELETE' })
        .then(response => response.json())
        .then(data => {
            fetchBooks();
            showMessage('Book deleted successfully');
        })
        .catch(error => showMessage('Error deleting book', true));
}

function populateForm(bookId) {
    fetch(`/book/${bookId}`)
        .then(response => response.json())
        .then(book => {
            document.getElementById('bookId').value = book.ID;
            document.getElementById('name').value = book.Name;
            document.getElementById('author').value = book.Author;
            document.getElementById('publication').value = book.Publication;
            document.getElementById('form-title').textContent = 'Update Book';
            document.getElementById('cancel-update').style.display = 'block';
        })
        .catch(error => showMessage('Error fetching book details', true));
}

document.getElementById('cancel-update').addEventListener('click', function() {
    resetForm();
});

function resetForm() {
    document.getElementById('book-form').reset();
    document.getElementById('bookId').value = '';
    document.getElementById('form-title').textContent = 'Add New Book';
    document.getElementById('cancel-update').style.display = 'none';
}

function showMessage(message, isError = false) {
    const messageDiv = document.getElementById('message');
    messageDiv.textContent = message;
    messageDiv.style.color = isError ? 'red' : 'green';
    setTimeout(() => { messageDiv.textContent = ''; }, 3000);
}