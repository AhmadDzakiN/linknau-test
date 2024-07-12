package questions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type jwtCustomClaims struct {
	Role string `json:"role"`
}

type response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type createTokenRequest struct {
	Role string `json:"role"`
}

func RunNewServer() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/create", createTokenHandler)

	fmt.Println("Starting server at 8080....")
	err := http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		log.Fatal("Failed to listen and serve: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := extractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jwtClaims, err := validateJWT(tokenStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Role authorization check
	claims, _ := jwtClaims.(jwtCustomClaims)
	if claims.Role != "admin" {
		http.Error(w, "forbidden access", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := response{
		Message: "Hello World!",
		Token:   tokenStr,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func createTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"role": req.Role, // this field will be used for role authorization
		"exp":  now.Add(time.Minute * 43200).Unix(),
		"iat":  now.Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := response{
		Message: "Create token success!",
		Token:   jwt,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

// Extract Authorization header to get JWT
func extractToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		return "", errors.New("invalid authorization header")
	}

	fields := strings.Fields(authorization)
	if len(fields) < 2 || fields[0] != "Bearer" {
		return "", errors.New("empty authorization header")
	}

	return fields[1], nil
}

// JWT validation from Authorization header value
func validateJWT(token string) (data interface{}, err error) {
	secretKey := "secret"
	extractedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there is an error in token parsing")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		err = fmt.Errorf("invalidate token: %w", err)
		return
	}

	if extractedToken == nil {
		err = errors.New("invalid token")
		return
	}

	claims, ok := extractedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("token error")
		return
	}

	data = jwtCustomClaims{
		Role: claims["role"].(string),
	}

	return
}

// User role authorization
func roleAuthorization(r *http.Request, role string) error {
	token := r.Context().Value("token")
	claims, ok := token.(jwtCustomClaims)
	if !ok || token == nil {
		return errors.New("invalid token")
	}

	if claims.Role == "" || claims.Role != role {
		return errors.New("unauthorized, forbidden access")
	}

	return nil
}
