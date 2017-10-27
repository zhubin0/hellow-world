// Package cdn implements the Azure ARM Cdn service API version 2016-10-02.
//
// Use these APIs to manage Azure CDN resources through the Azure Resource
// Manager. You must make sure that requests made to these resources are
// secure.
package cdn

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Cdn
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Cdn.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckNameAvailability check the availability of a resource name. This is
// needed for resources where name is globally unique, such as a CDN endpoint.
//
// checkNameAvailabilityInput is input to check.
func (client ManagementClient) CheckNameAvailability(checkNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "checkNameAvailabilityInput.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.ManagementClient", "CheckNameAvailability")
	}

	req, err := client.CheckNameAvailabilityPreparer(checkNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ManagementClient) CheckNameAvailabilityPreparer(checkNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Cdn/checkNameAvailability"),
		autorest.WithJSON(checkNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ManagementClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOperations lists all of the available CDN REST API operations.
func (client ManagementClient) ListOperations() (result OperationListResult, err error) {
	req, err := client.ListOperationsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", resp, "Failure responding to request")
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client ManagementClient) ListOperationsPreparer() (*http.Request, error) {
	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Cdn/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListOperationsResponder(resp *http.Response) (result OperationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOperationsNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListOperationsNextResults(lastResults OperationListResult) (result OperationListResult, err error) {
	req, err := lastResults.OperationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", resp, "Failure sending next results request")
	}

	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListOperations", resp, "Failure responding to next results request")
	}

	return
}

// ListResourceUsage check the quota and actual usage of the CDN profiles under
// the given subscription.
func (client ManagementClient) ListResourceUsage() (result ResourceUsageListResult, err error) {
	req, err := client.ListResourceUsagePreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", resp, "Failure sending request")
		return
	}

	result, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", resp, "Failure responding to request")
	}

	return
}

// ListResourceUsagePreparer prepares the ListResourceUsage request.
func (client ManagementClient) ListResourceUsagePreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkResourceUsage", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListResourceUsageSender sends the ListResourceUsage request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListResourceUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResourceUsageResponder handles the response to the ListResourceUsage request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListResourceUsageResponder(resp *http.Response) (result ResourceUsageListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListResourceUsageNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListResourceUsageNextResults(lastResults ResourceUsageListResult) (result ResourceUsageListResult, err error) {
	req, err := lastResults.ResourceUsageListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListResourceUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", resp, "Failure sending next results request")
	}

	result, err = client.ListResourceUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ManagementClient", "ListResourceUsage", resp, "Failure responding to next results request")
	}

	return
}
