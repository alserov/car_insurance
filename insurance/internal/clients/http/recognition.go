package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/alserov/car_insurance/insurance/internal/clients"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"io"
	"net/http"
)

func NewRecognitionClient(addr string) clients.RecognitionClient {
	return &recognition{}
}

type recognition struct {
	cl *http.Client

	addr string
}

func (r recognition) CheckIfCarIsOK(ctx context.Context, image []byte) error {
	_, err := r.makeRequest(http.MethodGet, bytes.NewReader(image))
	return err
}

func (r recognition) CalcDamageMultiplier(ctx context.Context, image []byte) (float32, error) {
	body, err := r.makeRequest(http.MethodGet, bytes.NewReader(image))
	if err != nil {
		return 0, utils.NewError(err.Error(), utils.Internal)
	}

	var mult float32

	if err = json.NewDecoder(body).Decode(&mult); err != nil {
		return 0, utils.NewError(err.Error(), utils.Internal)
	}

	return mult, nil
}

func (r recognition) makeRequest(method string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, r.addr, body)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	res, err := r.cl.Do(req)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusBadRequest {
			return nil, utils.NewError(err.Error(), utils.BadRequest)
		}
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	return res.Body, nil
}
