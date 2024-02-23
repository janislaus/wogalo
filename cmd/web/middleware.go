package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request using Gorilla Sessions
func SessionLoad(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "session-name")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Example usage of the session: Checking if a "user_id" exists in the session
			if userID, ok := session.Values["remote_ip"]; ok {
				// User ID exists, you can use it for your logic
				fmt.Println("User ID from session:", userID)
			} else {
				// User ID doesn't exist, maybe set a default or perform some action
				fmt.Println("No user ID in session")
			}

			// Make sure to save the session data before writing to the response or returning from the handler
			defer sessions.Save(r, w)

			next.ServeHTTP(w, r)
		})
	}
}
