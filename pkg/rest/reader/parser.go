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
	usrID := web.Params(r)["user"]
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

	/*
	// user validation
	if err := rest.ValidateFields(); err != nil {
		return nil, errorInvalidRequest.WithError(err)
	}
	 */

	return user, nil
}

/*
func parseSearchRequestFromRequest(ctx context.Context, r *http.Request, request *Request) error {
	err := readBody(ctx, r, request)
	if err != nil {
		return err
	}

	err = isValid(request.SearchRequest)
	if err != nil {
		logger.Error(ctx, "Search request is empty", logger.Tag{"error": err})
		return err
	}



	return nil
}
*/

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

/*
func isValid(request basic.SearchRequest) error {
	if request.FundID == nil && request.UserID == nil && request.Status == nil {
		return errors.New("invalid_search_request", "Search request must has a param", errors.StatusBadRequest)
	}
	return nil
}

 */


