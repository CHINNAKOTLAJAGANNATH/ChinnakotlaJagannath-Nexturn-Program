package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for valid Basic Auth credentials.
func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Missing or invalid Authorization header")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Decode the Base64-encoded credentials
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Failed to decode Authorization header:", err)
			c.JSON(401, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		// Split the username and password
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format:", string(payload))
			c.JSON(401, gin.H{"error": "Invalid Credentials"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		// Log the decoded username and password (for debugging)
		fmt.Println("Decoded credentials - Username:", username, "Password:", password)

		// Validate credentials against the database
		var storedPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, the username doesn't exist
				fmt.Println("User not found:", username)
				c.JSON(401, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			// Database error
			fmt.Println("Error querying database:", err)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Compare the provided password with the stored password (in plaintext)
		if storedPassword != password {
			fmt.Println("Password mismatch for user:", username)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// If authentication is successful
		fmt.Println("Authentication successful for user:", username)
		c.Next()
	}
}
