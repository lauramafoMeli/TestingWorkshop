package handlers

import (
	"net/http"
	"testdoubles/hunter"
	"testdoubles/pkg/web/request"
	"testdoubles/pkg/web/response"
	"testdoubles/positioner"
	"testdoubles/prey"
)

// NewControllerHunter return handler for hunter
func NewControllerHunter(ht hunter.Hunter, pr prey.Prey) *ControllerHunter {
	return &ControllerHunter{ht: ht, pr: pr}
}

// ControllerHunter return handler for hunter
type ControllerHunter struct {
	ht hunter.Hunter
	pr prey.Prey
}

// ConfigurePrey configure prey
type RequestBodyPrey struct {
	Speed	 float64				`json:"speed"`
	Position *positioner.Position	`json:"position"`
}
type ResponseBodyPrey struct {
	Message	 string `json:"message"`
	Data	 any	`json:"data"`
	Error	 bool	`json:"error"`
}
func (ct *ControllerHunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyPrey
		err := request.JSON(r, &reqBody)
		if err != nil {
			code := http.StatusBadRequest
			body := ResponseBodyPrey{
				Message: "Invalid request body",
				Data:    nil,
				Error:   true,
			}

			response.JSON(w, code, body)
			return
		}

		// process
		ct.pr.Configure(reqBody.Speed, reqBody.Position)

		// response
		code := http.StatusOK
		body := ResponseBodyPrey{
			Message: "Prey configured",
			Data:    nil,
			Error:   false,
		}

		response.JSON(w, code, body)
	}
}

// ConfigureHunter configure hunter
type RequestBodyHunter struct {
	Speed	 float64				`json:"speed"`
	Position *positioner.Position	`json:"position"`
}
type ResponseBodyHunter struct {
	Message	 string `json:"message"`
	Data	 any	`json:"data"`
	Error	 bool	`json:"error"`
}
func (ct *ControllerHunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyHunter
		err := request.JSON(r, &reqBody)
		if err != nil {
			code := http.StatusBadRequest
			body := ResponseBodyHunter{
				Message: "Invalid request body",
				Data:    nil,
				Error:   true,
			}

			response.JSON(w, code, body)
			return
		}

		// process
		ct.ht.Configure(reqBody.Speed, reqBody.Position)

		// response
		code := http.StatusOK
		body := ResponseBodyHunter{
			Message: "Hunter configured",
			Data:    nil,
			Error:   false,
		}

		response.JSON(w, code, body)
	}
}

// Hunt simulate hunt
type ResponseBodyHunt struct {
	Message	 string `json:"message"`
	Data	 any	`json:"data"`
	Error	 bool	`json:"error"`
}
func (ct *ControllerHunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		err := ct.ht.Hunt(ct.pr)
		if err != nil {
			code := http.StatusInternalServerError
			body := ResponseBodyHunt{
				Message: "Hunter could not catch prey",
				Data:    nil,
				Error:   true,
			}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyHunt{
			Message: "Hunter could catch the prey",
			Data:    nil,
			Error:   false,
		}

		response.JSON(w, code, body)
	}
}