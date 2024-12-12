package auth

import (
	"net/http"
	"time"
)

func (a *Auth) SetCookie(w http.ResponseWriter, name, value string, expires time.Time) {
	cookie := http.Cookie{
		Name:     name,    // Nazwa ciasteczka
		Value:    value,   // Wartość ciasteczka
		Expires:  expires, // Data wygaśnięcia (za 24 godziny)
		HttpOnly: true,    // Ciasteczko tylko dla HTTP (niedostępne dla JS)
		Secure:   true,    // Ustaw na true, jeśli używasz HTTPS
		Path:     "/",     // Ścieżka ciasteczka
	}

	http.SetCookie(w, &cookie)

}
