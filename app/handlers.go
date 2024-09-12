package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func getcurrentTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	timeZones := queryParams.Get("tz")

	timeMap := make(map[string]string)
	if timeZones == "" {
		timeMap["UTC"] = time.Now().UTC().Format("2006-01-02 15:04:05 -0700 MST")
		json.NewEncoder(w).Encode(timeMap)
		return
	}

	timeZoneArray := strings.Split(timeZones, ",")

	for _, value := range timeZoneArray {
		loc, err := time.LoadLocation(value)
		if err != nil {
			http.Error(w, "Invalid timezone: "+value, http.StatusNotFound)
			return
		} else {
			timeMap[value] = time.Now().In(loc).Format("2006-01-02 15:04:05 -0700 MST")
		}
	}

	json.NewEncoder(w).Encode(timeMap)
}
