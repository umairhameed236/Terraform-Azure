package frontdoorapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
    "github.com/Azure/go-autorest/autorest/date"
)

        // NetworkExperimentProfilesClientAPI contains the set of methods on the NetworkExperimentProfilesClient type.
        type NetworkExperimentProfilesClientAPI interface {
            CreateOrUpdate(ctx context.Context, profileName string, resourceGroupName string, parameters frontdoor.Profile) (result frontdoor.NetworkExperimentProfilesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.NetworkExperimentProfilesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.Profile, err error)
            List(ctx context.Context) (result frontdoor.ProfileListPage, err error)
                ListComplete(ctx context.Context) (result frontdoor.ProfileListIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result frontdoor.ProfileListPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result frontdoor.ProfileListIterator, err error)
            Update(ctx context.Context, resourceGroupName string, profileName string, parameters frontdoor.ProfileUpdateModel) (result frontdoor.NetworkExperimentProfilesUpdateFuture, err error)
        }

        var _ NetworkExperimentProfilesClientAPI = (*frontdoor.NetworkExperimentProfilesClient)(nil)
        // PreconfiguredEndpointsClientAPI contains the set of methods on the PreconfiguredEndpointsClient type.
        type PreconfiguredEndpointsClientAPI interface {
            List(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.PreconfiguredEndpointListPage, err error)
                ListComplete(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.PreconfiguredEndpointListIterator, err error)
        }

        var _ PreconfiguredEndpointsClientAPI = (*frontdoor.PreconfiguredEndpointsClient)(nil)
        // ExperimentsClientAPI contains the set of methods on the ExperimentsClient type.
        type ExperimentsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters frontdoor.Experiment) (result frontdoor.ExperimentsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (result frontdoor.ExperimentsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, profileName string, experimentName string) (result frontdoor.Experiment, err error)
            ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.ExperimentListPage, err error)
                ListByProfileComplete(ctx context.Context, resourceGroupName string, profileName string) (result frontdoor.ExperimentListIterator, err error)
            Update(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters frontdoor.ExperimentUpdateModel) (result frontdoor.ExperimentsUpdateFuture, err error)
        }

        var _ ExperimentsClientAPI = (*frontdoor.ExperimentsClient)(nil)
        // ReportsClientAPI contains the set of methods on the ReportsClient type.
        type ReportsClientAPI interface {
            GetLatencyScorecards(ctx context.Context, resourceGroupName string, profileName string, experimentName string, aggregationInterval frontdoor.LatencyScorecardAggregationInterval, endDateTimeUTC string, country string) (result frontdoor.LatencyScorecard, err error)
            GetTimeseries(ctx context.Context, resourceGroupName string, profileName string, experimentName string, startDateTimeUTC date.Time, endDateTimeUTC date.Time, aggregationInterval frontdoor.TimeseriesAggregationInterval, timeseriesType frontdoor.TimeseriesType, endpoint string, country string) (result frontdoor.Timeseries, err error)
        }

        var _ ReportsClientAPI = (*frontdoor.ReportsClient)(nil)
        // NameAvailabilityClientAPI contains the set of methods on the NameAvailabilityClient type.
        type NameAvailabilityClientAPI interface {
            Check(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
        }

        var _ NameAvailabilityClientAPI = (*frontdoor.NameAvailabilityClient)(nil)
        // NameAvailabilityWithSubscriptionClientAPI contains the set of methods on the NameAvailabilityWithSubscriptionClient type.
        type NameAvailabilityWithSubscriptionClientAPI interface {
            Check(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
        }

        var _ NameAvailabilityWithSubscriptionClientAPI = (*frontdoor.NameAvailabilityWithSubscriptionClient)(nil)
        // FrontDoorsClientAPI contains the set of methods on the FrontDoorsClient type.
        type FrontDoorsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters frontdoor.FrontDoor) (result frontdoor.FrontDoorsCreateOrUpdateFutureType, err error)
            Delete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoorsDeleteFutureType, err error)
            Get(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoor, err error)
            List(ctx context.Context) (result frontdoor.ListResultPage, err error)
                ListComplete(ctx context.Context) (result frontdoor.ListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultIterator, err error)
            ValidateCustomDomain(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties frontdoor.ValidateCustomDomainInput) (result frontdoor.ValidateCustomDomainOutput, err error)
        }

        var _ FrontDoorsClientAPI = (*frontdoor.FrontDoorsClient)(nil)
        // FrontendEndpointsClientAPI contains the set of methods on the FrontendEndpointsClient type.
        type FrontendEndpointsClientAPI interface {
            DisableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpointsDisableHTTPSFuture, err error)
            EnableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration frontdoor.CustomHTTPSConfiguration) (result frontdoor.FrontendEndpointsEnableHTTPSFuture, err error)
            Get(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpoint, err error)
            ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultPage, err error)
                ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultIterator, err error)
        }

        var _ FrontendEndpointsClientAPI = (*frontdoor.FrontendEndpointsClient)(nil)
        // EndpointsClientAPI contains the set of methods on the EndpointsClient type.
        type EndpointsClientAPI interface {
            PurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths frontdoor.PurgeParameters) (result frontdoor.EndpointsPurgeContentFuture, err error)
        }

        var _ EndpointsClientAPI = (*frontdoor.EndpointsClient)(nil)
        // RulesEnginesClientAPI contains the set of methods on the RulesEnginesClient type.
        type RulesEnginesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string, rulesEngineParameters frontdoor.RulesEngine) (result frontdoor.RulesEnginesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (result frontdoor.RulesEnginesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, frontDoorName string, rulesEngineName string) (result frontdoor.RulesEngine, err error)
            ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.RulesEngineListResultPage, err error)
                ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.RulesEngineListResultIterator, err error)
        }

        var _ RulesEnginesClientAPI = (*frontdoor.RulesEnginesClient)(nil)
        // PoliciesClientAPI contains the set of methods on the PoliciesClient type.
        type PoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters frontdoor.WebApplicationFirewallPolicy) (result frontdoor.PoliciesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.PoliciesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.WebApplicationFirewallPolicy, err error)
            List(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListPage, err error)
                ListComplete(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListIterator, err error)
        }

        var _ PoliciesClientAPI = (*frontdoor.PoliciesClient)(nil)
        // ManagedRuleSetsClientAPI contains the set of methods on the ManagedRuleSetsClient type.
        type ManagedRuleSetsClientAPI interface {
            List(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListPage, err error)
                ListComplete(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListIterator, err error)
        }

        var _ ManagedRuleSetsClientAPI = (*frontdoor.ManagedRuleSetsClient)(nil)
