package rest

import (
	"context"
	"encoding/json"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/errors"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"io/ioutil"
	"net/http"
)
var (
	errorInvalidRequest = errors.New("invalid_request", "Error validating request", errors.StatusBadRequest) //nolint
)

func parseUserFromRequestID(r *http.Request) (string, error) { //nolint
	// cdbID
	usrID := web.Params(r)["id"]
	if usrID == "" {
		return "", errorInvalidRequest.WithMessage("empty user")
	}

	return usrID, nil
}

func parseUserFromRequestBODY(ctx context.Context, r *http.Request) (*basic.User, error) { //nolint
	// Unmarshal
	var request basic.User
	if err := readBody(ctx, r, &request); err != nil {
		return nil, errorInvalidRequest.WithError(err)
	}
	user := &basic.User{
		ID: request.ID,
		FirstName: request.FirstName,
		LastName: request.LastName,
		Dni:  request.Dni,
		BirthDate: request.BirthDate,
		Email:   request.Email,
		Nacionality: request.Nacionality,
		Address: request.Address,
	}


	// user validation
	if err := isValid(*user); err != nil {
		return nil, errorInvalidRequest.WithError(err)
	}


	return user, nil
}

func readBody(ctx context.Context, r *http.Request, v interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logger.Error(ctx, "Can't read body", logger.Tag{"error": err})
		return errors.New("invalid_request", "Can't read body", errors.StatusBadRequest)
	}

	// Unmarshal
	err = json.Unmarshal(b, v)
	if err != nil {
		logger.Error(ctx, "Can't unmarshal body", logger.Tag{"error": err})
		return errors.New("invalid_body", "Can't unmarshal body", errors.StatusBadRequest)
	}

	return nil
}


func isValid(request basic.User) error {
	if request.ID == ""  {
		return errors.New("invalid", "User request must has a param", errors.StatusBadRequest)
	}
	return nil
}




