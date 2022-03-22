package frontdoor

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// ExperimentsClient is the frontDoor Client
type ExperimentsClient struct {
    BaseClient
}
// NewExperimentsClient creates an instance of the ExperimentsClient client.
func NewExperimentsClient(subscriptionID string) ExperimentsClient {
    return NewExperimentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExperimentsClientWithBaseURI creates an instance of the ExperimentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewExperimentsClientWithBaseURI(baseURI string, subscriptionID string) ExperimentsClient {
        return ExperimentsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate sends the create or update request.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
        // experimentName - the Experiment identifier associated with the Experiment
        // parameters - the Experiment resource
func (client ExperimentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment) (result ExperimentsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: profileName,
         Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: experimentName,
         Constraints: []validation.Constraint{	{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("frontdoor.ExperimentsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, profileName, experimentName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ExperimentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters Experiment) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "experimentName": autorest.Encode("path",experimentName),
        "profileName": autorest.Encode("path",profileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExperimentsClient) CreateOrUpdateSender(req *http.Request) (future ExperimentsCreateOrUpdateFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client ExperimentsClient) CreateOrUpdateResponder(resp *http.Response) (result Experiment, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete sends the delete request.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
        // experimentName - the Experiment identifier associated with the Experiment
func (client ExperimentsClient) Delete(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (result ExperimentsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: profileName,
         Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: experimentName,
         Constraints: []validation.Constraint{	{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("frontdoor.ExperimentsClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, profileName, experimentName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ExperimentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "experimentName": autorest.Encode("path",experimentName),
        "profileName": autorest.Encode("path",profileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExperimentsClient) DeleteSender(req *http.Request) (future ExperimentsDeleteFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client ExperimentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get sends the get request.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
        // experimentName - the Experiment identifier associated with the Experiment
func (client ExperimentsClient) Get(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (result Experiment, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: profileName,
         Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: experimentName,
         Constraints: []validation.Constraint{	{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("frontdoor.ExperimentsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, profileName, experimentName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ExperimentsClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "experimentName": autorest.Encode("path",experimentName),
        "profileName": autorest.Encode("path",profileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExperimentsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ExperimentsClient) GetResponder(resp *http.Response) (result Experiment, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByProfile sends the list by profile request.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
func (client ExperimentsClient) ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result ExperimentListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.ListByProfile")
        defer func() {
            sc := -1
        if result.el.Response.Response != nil {
        sc = result.el.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: profileName,
         Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("frontdoor.ExperimentsClient", "ListByProfile", err.Error())
        }

            result.fn = client.listByProfileNextResults
    req, err := client.ListByProfilePreparer(ctx, resourceGroupName, profileName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "ListByProfile", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByProfileSender(req)
        if err != nil {
        result.el.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "ListByProfile", resp, "Failure sending request")
        return
        }

        result.el, err = client.ListByProfileResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "ListByProfile", resp, "Failure responding to request")
        return
        }
            if result.el.hasNextLink() && result.el.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListByProfilePreparer prepares the ListByProfile request.
    func (client ExperimentsClient) ListByProfilePreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "profileName": autorest.Encode("path",profileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByProfileSender sends the ListByProfile request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExperimentsClient) ListByProfileSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListByProfileResponder handles the response to the ListByProfile request. The method always
    // closes the http.Response Body.
    func (client ExperimentsClient) ListByProfileResponder(resp *http.Response) (result ExperimentList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByProfileNextResults retrieves the next set of results, if any.
            func (client ExperimentsClient) listByProfileNextResults(ctx context.Context, lastResults ExperimentList) (result ExperimentList, err error) {
            req, err := lastResults.experimentListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "listByProfileNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByProfileSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "listByProfileNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByProfileResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "listByProfileNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByProfileComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ExperimentsClient) ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result ExperimentListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.ListByProfile")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByProfile(ctx, resourceGroupName, profileName)
                            return
            }

// Update updates an Experiment
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
        // experimentName - the Experiment identifier associated with the Experiment
        // parameters - the Experiment Update Model
func (client ExperimentsClient) Update(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel) (result ExperimentsUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExperimentsClient.Update")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: profileName,
         Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
        { TargetValue: experimentName,
         Constraints: []validation.Constraint{	{Target: "experimentName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("frontdoor.ExperimentsClient", "Update", err.Error())
        }

        req, err := client.UpdatePreparer(ctx, resourceGroupName, profileName, experimentName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Update", nil , "Failure preparing request")
    return
    }

        result, err = client.UpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "frontdoor.ExperimentsClient", "Update", result.Response(), "Failure sending request")
        return
        }

    return
}

    // UpdatePreparer prepares the Update request.
    func (client ExperimentsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters ExperimentUpdateModel) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "experimentName": autorest.Encode("path",experimentName),
        "profileName": autorest.Encode("path",profileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPatch(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExperimentsClient) UpdateSender(req *http.Request) (future ExperimentsUpdateFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // UpdateResponder handles the response to the Update request. The method always
    // closes the http.Response Body.
    func (client ExperimentsClient) UpdateResponder(resp *http.Response) (result Experiment, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

