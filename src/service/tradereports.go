// Package service implemented APIs logic.
package service

import (
	"encoding/json"
	"log"
	"net/http"

	"../domain/trade"
)

// lastTradeOfAllInstruments contains last trade of all instruments data.
type lastTradeOfAllInstruments struct {
	InstumentName string `json:"instument_name"`
	trade.Trade
}

// GetLastTradeOfAllInstruments return a list of last trade of all instruments.
func getLastTradeOfAllInstruments() ([]lastTradeOfAllInstruments, error) {
	lastTradeOfAllInstrumentsList := []lastTradeOfAllInstruments{}
	rows, err := trade.GetLastTradeOfAllInstrumentsRows()
	if err != nil {
		return []lastTradeOfAllInstruments{}, err
	}
	defer rows.Close()
	lastTradeOfInstrument := lastTradeOfAllInstruments{}
	for rows.Next() {
		err := rows.Scan(&lastTradeOfInstrument.InstumentName, &lastTradeOfInstrument.TradeID, &lastTradeOfInstrument.InstumentID, &lastTradeOfInstrument.DateEN, &lastTradeOfInstrument.Open, &lastTradeOfInstrument.High, &lastTradeOfInstrument.Low, &lastTradeOfInstrument.Close)
		if err != nil {
			return []lastTradeOfAllInstruments{}, err
		}
		lastTradeOfAllInstrumentsList = append(lastTradeOfAllInstrumentsList, lastTradeOfInstrument)
	}
	return lastTradeOfAllInstrumentsList, nil
}

// GetLastTradeOfAllInstrumentsHandler write a list of last trade of all instruments on ResponseWriter.
func GetLastTradeOfAllInstrumentsHandler(w http.ResponseWriter, r *http.Request) {
	lastTradeOfAllInstruments, err := getLastTradeOfAllInstruments()
	if err == nil {
		responseBytes, errjson := json.Marshal(lastTradeOfAllInstruments)
		if errjson == nil {
			writeResultOnResponse(w, responseBytes)
		} else {
			writeErrorOnResponse(w, errjson)
		}
	} else {
		writeErrorOnResponse(w, err)
	}
}

func writeResultOnResponse(w http.ResponseWriter, responseBytes []byte) {
	w.Header().Add("Content-Type", "application/json")
	status := http.StatusOK
	w.WriteHeader(status)
	w.Write(responseBytes)
}

func writeErrorOnResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Write([]byte("Error."))
}
