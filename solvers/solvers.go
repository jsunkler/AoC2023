package solvers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var dayMap = make(map[int](func(io.ReadCloser) (string, string, error)))

func init() {
	dayMap[1] = solveDay1
	dayMap[2] = solveDay2
}

func Solve(day int) (string, string, error) {
	if _, ok := dayMap[day]; !ok {
		return "", "", fmt.Errorf("day %d not yet implemented.", day)
	}

	rd, err := loadData(day)
	if err != nil {
		return "", "", fmt.Errorf("solving day %d failed. error: %w", day, err)
	}
	s1, s2, err := dayMap[day](rd)
	if err != nil {
		return "", "", err
	}

	return s1, s2, nil
}

func loadData(day int) (io.ReadCloser, error) {
	currentDayFileName := fmt.Sprintf("input.%d.txt", day)

	if _, err := os.Stat(currentDayFileName); errors.Is(err, os.ErrNotExist) {
		err2 := loadDataOnline(day, currentDayFileName)
		if err2 != nil {
			return nil, err2
		}
	}

	file, err := os.Open(currentDayFileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func loadDataOnline(day int, fileName string) error {

	fmt.Println("Input data not cached yet. Downloading it...")

	cookieFile, err := os.Open("cookie.txt")
	if err != nil {
		return fmt.Errorf("error opening session-cookie-file. error=%w", err)
	}
	defer cookieFile.Close()

	bytes, err := io.ReadAll(cookieFile)
	if err != nil {
		return fmt.Errorf("error reading session-cookie-file. error=%w", err)
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
		return fmt.Errorf("error constructing request. error=%w", err)
	}

	req.AddCookie(&cookie)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading data. error=%w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("got wrong StatusCode on download. statusCode=%d", response.StatusCode)
	}

	inputFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("could not write cached input file. error=%w", err)
	}
	defer inputFile.Close()

	if _, err := io.Copy(inputFile, response.Body); err != nil {
		return fmt.Errorf("could not write cached input file. error=%w", err)
	}

	return nil
}
