package controllers

import (
	"encoding/json"

	semp "github.com/lumberbarons/solace-operator/sempv2-config"
	"golang.org/x/net/context"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
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
	log := ctrllog.FromContext(ctx)

	swaggerError, _ := err.(semp.GenericSwaggerError)

	solaceError := SolaceError{}
	json.Unmarshal(swaggerError.Body(), &solaceError)

	if solaceError.SolaceErrorMeta.SolaceErrorError.Description != "" {
		params = append(params, "details", solaceError.SolaceErrorMeta.SolaceErrorError.Description)
	}

	log.Error(err, message, params...)

	return err
}
