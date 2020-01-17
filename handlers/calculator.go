package handlers

import (
	"encoding/json"
	"github.com/luisfmelo/calculator-api/utils"
	"net/http"
)

type Operation struct {
	Operation string
	Operand1  int
	Operand2  int
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	o := Operation{}
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result int

	switch o.Operation {
	case "sum":
		result = utils.Sum(o.Operand1, o.Operand2)
	case "subtract":
		http.Error(w, "Not supported", http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"result": result})
}
