package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type FrontdoorOriginGroupId struct {
	SubscriptionId  string
	ResourceGroup   string
	ProfileName     string
	OriginGroupName string
}

func NewFrontdoorOriginGroupID(subscriptionId, resourceGroup, profileName, originGroupName string) FrontdoorOriginGroupId {
	return FrontdoorOriginGroupId{
		SubscriptionId:  subscriptionId,
		ResourceGroup:   resourceGroup,
		ProfileName:     profileName,
		OriginGroupName: originGroupName,
	}
}

func (id FrontdoorOriginGroupId) String() string {
	segments := []string{
		fmt.Sprintf("Origin Group Name %q", id.OriginGroupName),
		fmt.Sprintf("Profile Name %q", id.ProfileName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Frontdoor Origin Group", segmentsStr)
}

func (id FrontdoorOriginGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Cdn/profiles/%s/originGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ProfileName, id.OriginGroupName)
}

// FrontdoorOriginGroupID parses a FrontdoorOriginGroup ID into an FrontdoorOriginGroupId struct
func FrontdoorOriginGroupID(input string) (*FrontdoorOriginGroupId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := FrontdoorOriginGroupId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.ProfileName, err = id.PopSegment("profiles"); err != nil {
		return nil, err
	}
	if resourceId.OriginGroupName, err = id.PopSegment("originGroups"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// FrontdoorOriginGroupIDInsensitively parses an FrontdoorOriginGroup ID into an FrontdoorOriginGroupId struct, insensitively
// This should only be used to parse an ID for rewriting, the FrontdoorOriginGroupID
// method should be used instead for validation etc.
//
// Whilst this may seem strange, this enables Terraform have consistent casing
// which works around issues in Core, whilst handling broken API responses.
func FrontdoorOriginGroupIDInsensitively(input string) (*FrontdoorOriginGroupId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := FrontdoorOriginGroupId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	// find the correct casing for the 'profiles' segment
	profilesKey := "profiles"
	for key := range id.Path {
		if strings.EqualFold(key, profilesKey) {
			profilesKey = key
			break
		}
	}
	if resourceId.ProfileName, err = id.PopSegment(profilesKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'originGroups' segment
	originGroupsKey := "originGroups"
	for key := range id.Path {
		if strings.EqualFold(key, originGroupsKey) {
			originGroupsKey = key
			break
		}
	}
	if resourceId.OriginGroupName, err = id.PopSegment(originGroupsKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
