package utils

import (
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pecet3/quizex/pkg/logger"
)

const PREFIX = "/v1"

func GetIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ip := xff
		if idx := net.ParseIP(ip); idx != nil {
			return ip
		}
	}
	xri := r.Header.Get("X-Real-Ip")
	if xri != "" {
		ip := xri
		if idx := net.ParseIP(ip); idx != nil {
			return ip
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}
	return userIP.String()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	logger.InfoWithCaller("Loaded .env")
}
