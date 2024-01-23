# Book Store Management System Documentation

## Introduction

Welcome to the documentation for the Book Store Management System, a robust application implemented in Go with a front end using JavaScript, HTML, and CSS. This system empowers users to perform CRUD (Create, Read, Update, Delete) operations on a collection of books stored in a PostgreSQL database

## Project Structure

- **cmd/main/main.go**: The primary entry point where the HTTP server is initialized using the *Gorilla Mux router*. Static files are served from the 'web' directory, and the application listens on a specified port (default is 9010).

- **pkg/config/app.go**: Manages the connection to the PostgreSQL database using the *GORM* library. It dynamically retrieves the database connection string from the `DATABASE_URL` environment variable.

- **pkg/controllers/book-controller.go**: Houses handlers for various *HTTP* routes related to book operations. It interfaces with the data layer through the models package, ensuring separation of concerns.

- **pkg/models/book.go**: Defines the Book model, encapsulating its structure and methods for CRUD operations. It also handles the initialization of the database connection and schema automigration.

- **pkg/routes/book-routes.go**: Registers HTTP routes related to book operations using the Gorilla Mux router. Each route is associated with a corresponding controller function, promoting clarity and modularity.

- **pkg/utils/utils.go**: Provides utility functions, including `ParseBody` for seamless parsing of JSON request bodies.

## API Endpoints

The system exposes a RESTful API with the following endpoints:

- **POST /book**: Create a new book.
- **GET /book**: Retrieve a list of all books.
- **GET /book/{bookId}**: Retrieve details of a specific book by ID.
- **PUT /book/{bookId}**: Update the details of a specific book by ID.
- **DELETE /book/{bookId}**: Delete a specific book by ID.

## Front End

The user interface is implemented using JavaScript, HTML, and CSS. Static files are served from the 'web' directory, providing an easily customizable front end. Developers can tailor the user experience by modifying the HTML and CSS files in this directory.

## Configuration

- The application utilizes the `PORT` environment variable to determine the port on which the server will listen, allowing flexibility in deployment environments.
- The database connection string is obtained from the `DATABASE_URL` environment variable. In local development, a fallback connection string is provided for ease of setup.

## Dependencies

The project leverages the following external libraries, adhering to industry standards:

- **Gorilla Mux**: A powerful HTTP router and URL matcher.
- **GORM**: An Object Relational Mapper (ORM) for Go, facilitating seamless database interactions.

## Conclusion

The Book Store Management System offers a solid foundation for efficiently managing books through a well-designed API. Whether you are a recruiter evaluating technical prowess or a developer looking to extend functionality, this project is poised for adaptability. Feel free to explore and customize the system to meet your specific requirements.