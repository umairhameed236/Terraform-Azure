package networking

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type PrivateEndpointConnectionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnection
}

// PrivateEndpointConnectionsGet ...
func (c NetworkingClient) PrivateEndpointConnectionsGet(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionsGetOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "networking.NetworkingClient", "PrivateEndpointConnectionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "networking.NetworkingClient", "PrivateEndpointConnectionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "networking.NetworkingClient", "PrivateEndpointConnectionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionsGet prepares the PrivateEndpointConnectionsGet request.
func (c NetworkingClient) preparerForPrivateEndpointConnectionsGet(ctx context.Context, id PrivateEndpointConnectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateEndpointConnectionsGet handles the response to the PrivateEndpointConnectionsGet request. The method always
// closes the http.Response Body.
func (c NetworkingClient) responderForPrivateEndpointConnectionsGet(resp *http.Response) (result PrivateEndpointConnectionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
