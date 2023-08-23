package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

// Integration Tests for ControllerHunter
func TestIntegration_ControllerHunter_ConfigurePrey(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
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
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			pr := &prey.Tuna{}

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

func TestIntegration_ControllerHunter_ConfigureHunter(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
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
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			ht := &hunter.WhiteShark{}

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

func TestIntegration_ControllerHunter_Hunt(t *testing.T) {
	type input struct { r *http.Request; rr *httptest.ResponseRecorder }
	type output struct { code int; body string }
	type testCase struct {
		name string
		input input
		output output
		// set-up
		setUpCfgTuna func (cfg *prey.ConfigTuna)
		setUpCfgCatchSimulatorDefault func (cfg *simulator.ConfigCatchSimulatorDefault)
		setUpCfgWhiteShark func (cfg *hunter.ConfigWhiteShark)
	}

	cases := []testCase{
		// succeed to hunt
		{
			name: "succeed to hunt - hunter is faster and in time",
			input: input{
				r: &http.Request{},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusOK,
				body: `{"message":"Hunter could catch the prey","data":null,"error":false}`,
			},
			setUpCfgTuna: func (cfg *prey.ConfigTuna) {
				(*cfg).Speed = 5
				(*cfg).Position = &positioner.Position{X: 100, Y: 0, Z: 0}
			},
			setUpCfgWhiteShark: func (cfg *hunter.ConfigWhiteShark) {
				(*cfg).Speed = 10
				(*cfg).Position = &positioner.Position{X: 0, Y: 0, Z: 0}
			},
			setUpCfgCatchSimulatorDefault: func (cfg *simulator.ConfigCatchSimulatorDefault) {
				(*cfg).MaxTimeToCatch = 100
			},
		},

		// fail to hunt
		{
			name: "fail to hunt - hunter is faster, but not in time",
			input: input{
				r: &http.Request{},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusInternalServerError,
				body: `{"message":"Hunter could not catch prey","data":null,"error":true}`,
			},
			setUpCfgTuna: func (cfg *prey.ConfigTuna) {
				(*cfg).Speed = 5
				(*cfg).Position = &positioner.Position{X: 100, Y: 0, Z: 0}
			},
			setUpCfgWhiteShark: func (cfg *hunter.ConfigWhiteShark) {
				(*cfg).Speed = 10
				(*cfg).Position = &positioner.Position{X: 0, Y: 0, Z: 0}
			},
			setUpCfgCatchSimulatorDefault: func (cfg *simulator.ConfigCatchSimulatorDefault) {
				(*cfg).MaxTimeToCatch = 1
			},
		},
		{
			name: "fail to hunt - hunter is slow",
			input: input{
				r: &http.Request{},
				rr: httptest.NewRecorder(),
			},
			output: output{
				code: http.StatusInternalServerError,
				body: `{"message":"Hunter could not catch prey","data":null,"error":true}`,
			},
			setUpCfgTuna: func (cfg *prey.ConfigTuna) {
				(*cfg).Speed = 10
				(*cfg).Position = &positioner.Position{X: 100, Y: 0, Z: 0}
			},
			setUpCfgWhiteShark: func (cfg *hunter.ConfigWhiteShark) {
				(*cfg).Speed = 5
				(*cfg).Position = &positioner.Position{X: 0, Y: 0, Z: 0}
			},
			setUpCfgCatchSimulatorDefault: func (cfg *simulator.ConfigCatchSimulatorDefault) {
				(*cfg).MaxTimeToCatch = 100
			},
		},
	}			

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			// -> prey
			cfgPr := &prey.ConfigTuna{}
			c.setUpCfgTuna(cfgPr)
			pr := prey.NewTuna(*cfgPr)

			// -> simulator
			ps := positioner.NewPositionerDefault()

			cfgSm := &simulator.ConfigCatchSimulatorDefault{
				Positioner: ps,
			}
			c.setUpCfgCatchSimulatorDefault(cfgSm)
			sm := simulator.NewCatchSimulatorDefault(cfgSm)

			// -> hunter
			cfgHt := &hunter.ConfigWhiteShark{
				Simulator: sm,
			}
			c.setUpCfgWhiteShark(cfgHt)
			ht := hunter.NewWhiteShark(*cfgHt)
			
			impl := NewControllerHunter(ht, pr)
			hd := impl.Hunt()

			// act
			hd(c.input.rr, c.input.r)

			// assert
			require.Equal(t, c.output.code, c.input.rr.Code)
			require.JSONEq(t, c.output.body, c.input.rr.Body.String())
		})
	}
}