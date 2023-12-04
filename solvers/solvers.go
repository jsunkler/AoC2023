package solvers

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var dayMap = make(map[int](func(io.ReadCloser) (string, string)))

func init() {
	dayMap[1] = solveDay1
}

func Solve(day int) (string, string) {
	rd := loadData(day)
	return dayMap[day](rd)
}

func loadData(day int) io.ReadCloser {
	file, err := os.Open("cookie.txt")
	if err != nil {
		slog.Error("Error opening session-cookie-file.", "err", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Error reading session-cookie-file", "err", err)
	}

	cookieVal := string(bytes)

	cookie := http.Cookie{
		Name:     "session",
		Value:    cookieVal,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day), nil)
	if err != nil {
		slog.Error("Error constructing request.", "err", err)
	}

	req.AddCookie(&cookie)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Error downloading data.", "err", err)
	}

	slog.Info("Download of inputdata succeeded.", "statusCode", response.StatusCode, "contentLength", response.ContentLength)

	return response.Body
}
