package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ControllerHunter
func TestControllerHunter_ConfigurePrey(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
		// set-up
		setUpPreyStub func(p *prey.PreyStub)
	}

	cases := []testCase{
		// succeed to configure prey
		{
			name: "succeed to configure prey",
			input: input{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(
						`{"speed": 10, "position": {"x": 10, "y": 10}}`,
					)),
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"Prey configured","data":null,"error":false}`,
			},
			setUpPreyStub: func(p *prey.PreyStub) {
				p.ConfigureFunc = func(speed float64, position *positioner.Position) {}
			},
		},

		// fail to configure prey
		{
			name: "fail to configure prey",
			input: input{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(
						`{wrong json}`,
					)),
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusBadRequest,
				body: `{"message":"Invalid request body","data":null,"error":true}`,
			},
			setUpPreyStub: func(p *prey.PreyStub) {
				p.ConfigureFunc = func(speed float64, position *positioner.Position) {
					panic("fail to configure prey")
				}
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			pr := prey.NewPreyStub()
			c.setUpPreyStub(pr)

			impl := NewControllerHunter(nil, pr)
			hd := impl.ConfigurePrey()

			// act
			hd(c.input.rr, c.input.r)

			// assert
			require.Equal(t, c.output.code, c.input.rr.Code)
			require.JSONEq(t, c.output.body, c.input.rr.Body.String())
		})
	}
}

func TestControllerHunter_ConfigureHunter(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
		// set-up
		setUpHunterMock func(h *hunter.HunterMock)
	}

	cases := []testCase{
		// succeed to configure hunter
		{
			name: "succeed to configure hunter",
			input: input{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(
						`{"speed": 10, "position": {"x": 10, "y": 10}}`,
					)),
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"Hunter configured","data":null,"error":false}`,
			},
			setUpHunterMock: func(h *hunter.HunterMock) {
				h.ConfigureFunc = func(speed float64, position *positioner.Position) {}
			},
		},

		// fail to configure hunter
		{
			name: "fail to configure hunter",
			input: input{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(
						`{wrong json}`,
					)),
				},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusBadRequest,
				body: `{"message":"Invalid request body","data":null,"error":true}`,
			},
			setUpHunterMock: func(h *hunter.HunterMock) {
				h.ConfigureFunc = func(speed float64, position *positioner.Position) {
					panic("fail to configure hunter")
				}
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			ht := hunter.NewHunterMock()
			c.setUpHunterMock(ht)

			impl := NewControllerHunter(ht, nil)
			hd := impl.ConfigureHunter()

			// act
			hd(c.input.rr, c.input.r)

			// assert
			require.Equal(t, c.output.code, c.input.rr.Code)
			require.JSONEq(t, c.output.body, c.input.rr.Body.String())
		})
	}
}

func TestControllerHunter_Hunt(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
		// set-up
		setUpHunterMock func(h *hunter.HunterMock)
	}

	cases := []testCase{
		// succeed to hunt
		{
			name: "succeed to hunt",
			input: input{
				r: &http.Request{},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"Hunter could catch the prey","data":null,"error":false}`,
			},
			setUpHunterMock: func(h *hunter.HunterMock) {
				h.HuntFunc = func(pr prey.Prey) (err error) {return}
			},
		},

		// fail to hunt
		{
			name: "fail to hunt",
			input: input{
				r: &http.Request{},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusInternalServerError,
				body: `{"message":"Hunter could not catch prey","data":null,"error":true}`,
			},
			setUpHunterMock: func(h *hunter.HunterMock) {
				h.HuntFunc = func(pr prey.Prey) (err error) {
					err = hunter.ErrCanNotHunt
					return
				}
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			ht := hunter.NewHunterMock()
			c.setUpHunterMock(ht)

			impl := NewControllerHunter(ht, nil)
			hd := impl.Hunt()

			// act
			hd(c.input.rr, c.input.r)

			// assert
			require.Equal(t, c.output.code, c.input.rr.Code)
			require.JSONEq(t, c.output.body, c.input.rr.Body.String())
		})
	}
}