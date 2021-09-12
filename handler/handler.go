// package handler release function to processing http requests
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paysolut/converter"
)

// base server structure with global objects such as db and config
type ServerType struct {
}

type RomanInputDataType struct {
	Numeral string
}

// filtering request type, decoding request body and convert roman numeral to arabic numeral.
// in case errors during those stages response with error code and specific message (if it needs).
// build response with arabic numeral.
func (server *ServerType) Handler(w http.ResponseWriter, r *http.Request) {
	var (
		err error

		decoder       *json.Decoder
		romanInput    RomanInputDataType
		arabicNumeral int
	)

	fmt.Println()

	// we wait only POST request
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)

		_, err = w.Write([]byte("I'm ready to POST only"))
		if err != nil {
			fmt.Println("[error] processing wrong request type:", err)
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}

		return
	}

	// convert request body to input data structure
	decoder = json.NewDecoder(r.Body)
	err = decoder.Decode(&romanInput)
	if err != nil {
		fmt.Println("[error] decode request params:", err)
		http.Error(w, "error", http.StatusInternalServerError)

		return
	}

	fmt.Println("roman numeral:", romanInput.Numeral)

	// convert numeral from roman to arabic
	arabicNumeral = converter.Convert(romanInput.Numeral)
	if arabicNumeral == 0 {
		fmt.Println("[error] numeral conversion failed")
		http.Error(w, "error", http.StatusInternalServerError)

		return
	}

	fmt.Println("arabic numeral:", arabicNumeral)

	// prepare response
	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(fmt.Sprintf("%d", arabicNumeral)))
	if err != nil {
		fmt.Println("[error] build success response:", err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}
