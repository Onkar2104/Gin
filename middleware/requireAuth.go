// package middleware

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/Onkar2104/go/initializers"
// 	"github.com/Onkar2104/go/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )

// func RequireAuth(c *gin.Context) {

// 	tokenString, err := c.Cookie("Authorization")

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(os.Getenv("SECRET")), nil
// 	})

// 	if err != nil || !token.Valid {
// 		log.Println("jwt parsing error:", err)
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		var user models.User
// 		initializers.DB.First(&user, claims["sub"])

// 		if user.ID == 0 {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		c.Set("user", user)

// 		c.Next()

// 	} else {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}
// }

package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Onkar2104/go/initializers"
	"github.com/Onkar2104/go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get the Authorization cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Println("missing authorization cookie")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		log.Println("jwt parsing error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("invalid jwt claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check if token is expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		log.Println("token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Convert `sub` claim to integer
	sub, ok := claims["sub"].(float64)
	if !ok {
		log.Println("invalid sub claim")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID := int(sub) // Convert `sub` to `int`

	// Find the user by ID
	var user models.User
	result := initializers.DB.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		log.Println("user not found:", result.Error)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach user to request
	c.Set("user", user)
	c.Next()
}
