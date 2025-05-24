# Shopper Online

**Shopper Online** is a full-stack e-commerce web application designed to provide users with a seamless online shopping experience. The application allows users to browse products, manage their shopping cart, and place orders securely.

## Features

### Frontend

* **User-Friendly Interface**: Built with Vue.js and Nuxt.js to ensure a dynamic and responsive user experience.
* **Responsive Design**: Utilizes TailwindCSS for a modern and mobile-friendly layout.
* **Product Browsing**: Users can browse through a catalog of products and view detailed information on each product page.

### Backend

* **RESTful API**: Developed using Go (Golang) to handle all server-side operations.
* **Database Management**: Employs MongoDB for efficient and scalable data storage.
* **Secure Transactions**: Implements secure methods for user authentication and order processing.

## Technologies Used

* **Frontend**:

  * [Vue.js](https://vuejs.org/)
  * [Nuxt.js](https://nuxtjs.org/)
  * [TailwindCSS](https://tailwindcss.com/)

* **Backend**:

  * [Go (Golang)](https://golang.org/)
  * [MongoDB](https://www.mongodb.com/)

## Getting Started

### Prerequisites

* **Node.js**: Ensure you have Node.js installed on your machine. You can download it from [here](https://nodejs.org/).
* **Go**: Install Go from the official website [here](https://golang.org/dl/).
* **MongoDB**: Install MongoDB and ensure it's running. Instructions can be found [here](https://docs.mongodb.com/manual/installation/).

### Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/PhoengZ/Shopper-online.git
   cd Shopper-online
   ```

2. **Install Backend Dependencies**:

   ```bash
   cd backend
   go mod download
   ```

3. **Install Frontend Dependencies**:

   ```bash
   cd ../frontend
   npm install
   ```

### Running the Application

1. **Start the Backend Server**:

   ```bash
   cd backend
   go run main.go
   ```

2. **Start the Frontend Server**:

   ```bash
   cd ../frontend
   npm run dev
   ```

3. **Access the Application**:

   Open your browser and navigate to `http://localhost:3000` to view the application.

## Project Structure

```
Shopper-online/
├── backend/           # Go backend
│   ├── controllers/   # Route controllers
│   ├── models/        # Data models
│   ├── routes/        # API routes
│   └── main.go        # Main application file
├── frontend/          # Nuxt.js frontend
│   ├── components/    # Vue.js components
│   ├── pages/         # Nuxt.js pages
│   └── nuxt.config.js # Nuxt.js configuration
└── README.md          # Project documentation
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the [MIT License](LICENSE).

---

For more details, visit the [Shopper-online GitHub repository](https://github.com/PhoengZ/Shopper-online/tree/main).
