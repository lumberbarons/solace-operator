package controllers

import (
	"golang.org/x/net/context"
)

type SolaceError struct {
	SolaceErrorMeta SolaceErrorMeta `json:"meta"`
	ResponseCode    string          `json:"responseCode"`
}

type SolaceErrorMeta struct {
	SolaceErrorError SolaceErrorError `json:"error"`
}

type SolaceErrorError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func solaceError(ctx context.Context, err error, message string, params ...interface{}) error {
	/* log := ctrllog.FromContext(ctx)

	swaggerError, _ := err.(semp.GenericSwaggerError)

	solaceError := SolaceError{}
	json.Unmarshal(swaggerError.Body(), &solaceError)

	if solaceError.SolaceErrorMeta.SolaceErrorError.Description != "" {
		params = append(params, "details", solaceError.SolaceErrorMeta.SolaceErrorError.Description)
	}

	log.Error(err, message, params...) */

	return err
}
