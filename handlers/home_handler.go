package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the root endpoint and serves API documentation
type HomeHandler struct{}

// NewHomeHandler creates a new home handler
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// ServeHome handles GET / requests and returns API documentation
func (h *HomeHandler) ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>User CRUD API</h1>
		<p>Available endpoints:</p>
		<ul>
			<li><strong>GET /</strong> - This help page</li>
			<li><strong>POST /users</strong> - Create a new user</li>
			<li><strong>GET /users</strong> - Get all users</li>
			<li><strong>GET /users/{id}</strong> - Get a specific user</li>
			<li><strong>PUT /users/{id}</strong> - Update a user</li>
			<li><strong>DELETE /users/{id}</strong> - Delete a user</li>
		</ul>
		<p>Example POST /users body:</p>
		<pre>{"name": "John Doe", "email": "john@example.com"}</pre>
	`)
}
