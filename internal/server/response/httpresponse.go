package response

import (
	"net/http"

	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/utils"
)

func ResponseError(w http.ResponseWriter, err error) {
	apiError, ok := err.(*apierrors.ApiError)
	if !ok {
		apiError = apierrors.InternalServerError(err.Error())
	}
	utils.Encode(w, apiError.Status, apiError)
}
