package cdn

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cdn/parse"
	track1 "github.com/hashicorp/terraform-provider-azurerm/internal/services/cdn/sdk/2021-06-01"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/cdn/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceFrontdoorOriginGroup() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceFrontdoorOriginGroupCreate,
		Read:   resourceFrontdoorOriginGroupRead,
		Update: resourceFrontdoorOriginGroupUpdate,
		Delete: resourceFrontdoorOriginGroupDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.FrontdoorOriginGroupID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:     pluginsdk.TypeString,
				Required: true,
				ForceNew: true,
			},

			"frontdoor_profile_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.FrontdoorProfileID,
			},

			"deployment_status": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"health_probe": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,

				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{

						"interval_in_seconds": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							Default:      240,
							ValidateFunc: validation.IntBetween(5, 31536000),
						},

						"path": {
							Type:         pluginsdk.TypeString,
							Optional:     true,
							Default:      "/",
							ValidateFunc: validation.StringIsNotEmpty,
						},

						"protocol": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Default:  string(track1.ProbeProtocolHTTPS),
							ValidateFunc: validation.StringInSlice([]string{
								string(track1.ProbeProtocolHTTP),
								string(track1.ProbeProtocolHTTPS),
								string(track1.ProbeProtocolNotSet),
							}, false),
						},

						"request_type": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Default:  string(track1.HealthProbeRequestTypeGET),
							ValidateFunc: validation.StringInSlice([]string{
								string(track1.HealthProbeRequestTypeGET),
								string(track1.HealthProbeRequestTypeHEAD),
								string(track1.HealthProbeRequestTypeNotSet),
							}, false),
						},
					},
				},
			},

			"load_balancing": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,

				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{

						"additional_latency_in_milliseconds": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(0, 1000),
						},

						"sample_size": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							Default:      16,
							ValidateFunc: validation.IntBetween(0, 255),
						},

						"successful_samples_required": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							Default:      3,
							ValidateFunc: validation.IntBetween(0, 255),
						},
					},
				},
			},

			"frontdoor_profile_name": {
				Type:     pluginsdk.TypeString,
				Computed: true,
			},

			"response_based_origin_error_detection": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,

				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{

						"http_error_ranges": {
							Type:     pluginsdk.TypeList,
							Optional: true,

							Elem: &pluginsdk.Resource{
								Schema: map[string]*pluginsdk.Schema{

									"begin": {
										Type:         pluginsdk.TypeInt,
										Optional:     true,
										Default:      300,
										ValidateFunc: validation.IntBetween(100, 999),
									},

									"end": {
										Type:         pluginsdk.TypeInt,
										Optional:     true,
										Default:      599,
										ValidateFunc: validation.IntBetween(100, 999),
									},
								},
							},
						},

						"detected_error_types": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Default:  string(track1.ResponseBasedDetectedErrorTypesTCPAndHTTPErrors),
							ValidateFunc: validation.StringInSlice([]string{
								string(track1.ResponseBasedDetectedErrorTypesNone),
								string(track1.ResponseBasedDetectedErrorTypesTCPAndHTTPErrors),
								string(track1.ResponseBasedDetectedErrorTypesTCPErrorsOnly),
							}, false),
						},

						"failover_threshold_percentage": {
							Type:         pluginsdk.TypeInt,
							Optional:     true,
							Default:      10,
							ValidateFunc: validation.IntBetween(0, 100),
						},
					},
				},
			},

			"session_affinity": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"restore_traffic_or_new_endpoints_time": {
				Type:         pluginsdk.TypeInt,
				Optional:     true,
				Default:      10,
				ValidateFunc: validation.IntBetween(0, 50),
			},
		},
	}
}

func resourceFrontdoorOriginGroupCreate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Cdn.FrontDoorOriginGroupsClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	profileId, err := parse.FrontdoorProfileID(d.Get("frontdoor_profile_id").(string))
	if err != nil {
		return err
	}

	id := parse.NewFrontdoorOriginGroupID(profileId.SubscriptionId, profileId.ResourceGroup, profileId.ProfileName, d.Get("name").(string))

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.ProfileName, id.OriginGroupName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for existing %s: %+v", id, err)
			}
		}

		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_frontdoor_origin_group", id.ID())
		}
	}

	props := track1.AFDOriginGroup{
		AFDOriginGroupProperties: &track1.AFDOriginGroupProperties{
			HealthProbeSettings:                                   expandOriginGroupHealthProbeParameters(d.Get("health_probe").([]interface{})),
			LoadBalancingSettings:                                 expandOriginGroupLoadBalancingSettingsParameters(d.Get("load_balancing").([]interface{})),
			ResponseBasedAfdOriginErrorDetectionSettings:          expandOriginGroupResponseBasedOriginErrorDetectionParameters(d.Get("response_based_origin_error_detection").([]interface{})),
			SessionAffinityState:                                  ConvertBoolToEnabledState(d.Get("session_affinity").(bool)),
			TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: utils.Int32(int32(d.Get("restore_traffic_or_new_endpoints_time").(int))),
		},
	}

	future, err := client.Create(ctx, id.ResourceGroup, id.ProfileName, id.OriginGroupName, props)
	if err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the creation of %s: %+v", id, err)
	}

	d.SetId(id.ID())
	return resourceFrontdoorOriginGroupRead(d, meta)
}

func resourceFrontdoorOriginGroupRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Cdn.FrontDoorOriginGroupsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FrontdoorOriginGroupID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.ProfileName, id.OriginGroupName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", id, err)
	}

	d.Set("name", id.OriginGroupName)
	d.Set("frontdoor_profile_id", parse.NewFrontdoorProfileID(id.SubscriptionId, id.ResourceGroup, id.ProfileName).ID())

	if props := resp.AFDOriginGroupProperties; props != nil {
		d.Set("deployment_status", props.DeploymentStatus)

		if err := d.Set("health_probe", flattenOriginGroupHealthProbeParameters(props.HealthProbeSettings)); err != nil {
			return fmt.Errorf("setting `health_probe`: %+v", err)
		}

		if err := d.Set("load_balancing", flattenOriginGroupLoadBalancingSettingsParameters(props.LoadBalancingSettings)); err != nil {
			return fmt.Errorf("setting `load_balancing`: %+v", err)
		}

		if err := d.Set("response_based_origin_error_detection", flattenOriginGroupResponseBasedOriginErrorDetectionParameters(props.ResponseBasedAfdOriginErrorDetectionSettings)); err != nil {
			return fmt.Errorf("setting `response_based_origin_error_detection`: %+v", err)
		}

		// BUG: API does not return the profile name, pull it form the ID
		d.Set("frontdoor_profile_name", id.ProfileName)
		d.Set("session_affinity", ConvertEnabledStateToBool(&props.SessionAffinityState))
		d.Set("restore_traffic_or_new_endpoints_time", props.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes)
	}

	return nil
}

func resourceFrontdoorOriginGroupUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Cdn.FrontDoorOriginGroupsClient
	ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FrontdoorOriginGroupID(d.Id())
	if err != nil {
		return err
	}

	props := track1.AFDOriginGroupUpdateParameters{
		AFDOriginGroupUpdatePropertiesParameters: &track1.AFDOriginGroupUpdatePropertiesParameters{
			HealthProbeSettings:                                   expandOriginGroupHealthProbeParameters(d.Get("health_probe").([]interface{})),
			LoadBalancingSettings:                                 expandOriginGroupLoadBalancingSettingsParameters(d.Get("load_balancing").([]interface{})),
			ResponseBasedAfdOriginErrorDetectionSettings:          expandOriginGroupResponseBasedOriginErrorDetectionParameters(d.Get("response_based_origin_error_detection").([]interface{})),
			SessionAffinityState:                                  ConvertBoolToEnabledState(d.Get("session_affinity").(bool)),
			TrafficRestorationTimeToHealedOrNewEndpointsInMinutes: utils.Int32(int32(d.Get("restore_traffic_or_new_endpoints_time").(int))),
		},
	}

	future, err := client.Update(ctx, id.ResourceGroup, id.ProfileName, id.OriginGroupName, props)
	if err != nil {
		return fmt.Errorf("updating %s: %+v", *id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the update of %s: %+v", *id, err)
	}

	return resourceFrontdoorOriginGroupRead(d, meta)
}

func resourceFrontdoorOriginGroupDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Cdn.FrontDoorOriginGroupsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.FrontdoorOriginGroupID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.ProfileName, id.OriginGroupName)
	if err != nil {
		return fmt.Errorf("deleting %s: %+v", *id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the deletion of %s: %+v", *id, err)
	}
	return err
}

func expandOriginGroupHealthProbeParameters(input []interface{}) *track1.HealthProbeParameters {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	v := input[0].(map[string]interface{})

	probeProtocolValue := track1.ProbeProtocol(v["protocol"].(string))
	probeRequestTypeValue := track1.HealthProbeRequestType(v["request_type"].(string))
	return &track1.HealthProbeParameters{
		ProbeIntervalInSeconds: utils.Int32(int32(v["interval_in_seconds"].(int))),
		ProbePath:              utils.String(v["path"].(string)),
		ProbeProtocol:          probeProtocolValue,
		ProbeRequestType:       probeRequestTypeValue,
	}
}

func expandOriginGroupLoadBalancingSettingsParameters(input []interface{}) *track1.LoadBalancingSettingsParameters {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	v := input[0].(map[string]interface{})

	return &track1.LoadBalancingSettingsParameters{
		AdditionalLatencyInMilliseconds: utils.Int32(int32(v["additional_latency_in_milliseconds"].(int))),
		SampleSize:                      utils.Int32(int32(v["sample_size"].(int))),
		SuccessfulSamplesRequired:       utils.Int32(int32(v["successful_samples_required"].(int))),
	}
}

func expandOriginGroupResponseBasedOriginErrorDetectionParameters(input []interface{}) *track1.ResponseBasedOriginErrorDetectionParameters {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	v := input[0].(map[string]interface{})

	responseBasedDetectedErrorTypesValue := track1.ResponseBasedDetectedErrorTypes(v["detected_error_types"].(string))
	return &track1.ResponseBasedOriginErrorDetectionParameters{
		HTTPErrorRanges:                          expandOriginGroupHttpErrorRangeParametersArray(v["http_error_ranges"].([]interface{})),
		ResponseBasedDetectedErrorTypes:          responseBasedDetectedErrorTypesValue,
		ResponseBasedFailoverThresholdPercentage: utils.Int32(int32(v["failover_threshold_percentage"].(int))),
	}
}

func expandOriginGroupHttpErrorRangeParametersArray(input []interface{}) *[]track1.HTTPErrorRangeParameters {
	results := make([]track1.HTTPErrorRangeParameters, 0)
	for _, item := range input {
		v := item.(map[string]interface{})

		results = append(results, track1.HTTPErrorRangeParameters{
			Begin: utils.Int32(int32(v["begin"].(int))),
			End:   utils.Int32(int32(v["end"].(int))),
		})
	}
	return &results
}

func flattenOriginGroupLoadBalancingSettingsParameters(input *track1.LoadBalancingSettingsParameters) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	result := make(map[string]interface{})

	if input.AdditionalLatencyInMilliseconds != nil {
		result["additional_latency_in_milliseconds"] = *input.AdditionalLatencyInMilliseconds
	}

	if input.SampleSize != nil {
		result["sample_size"] = *input.SampleSize
	}

	if input.SuccessfulSamplesRequired != nil {
		result["successful_samples_required"] = *input.SuccessfulSamplesRequired
	}
	return append(results, result)
}

func flattenOriginGroupResponseBasedOriginErrorDetectionParameters(input *track1.ResponseBasedOriginErrorDetectionParameters) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	result := make(map[string]interface{})
	result["http_error_ranges"] = flattenOriginGroupHttpErrorRangeParametersArray(input.HTTPErrorRanges)

	if input.ResponseBasedDetectedErrorTypes != "" {
		result["detected_error_types"] = input.ResponseBasedDetectedErrorTypes
	}

	if input.ResponseBasedFailoverThresholdPercentage != nil {
		result["failover_threshold_percentage"] = *input.ResponseBasedFailoverThresholdPercentage
	}
	return append(results, result)
}

func flattenOriginGroupHttpErrorRangeParametersArray(inputs *[]track1.HTTPErrorRangeParameters) []interface{} {
	results := make([]interface{}, 0)
	if inputs == nil {
		return results
	}

	for _, input := range *inputs {
		result := make(map[string]interface{})

		if input.Begin != nil {
			result["begin"] = *input.Begin
		}

		if input.End != nil {
			result["end"] = *input.End
		}
		results = append(results, result)
	}

	return results
}

func flattenOriginGroupHealthProbeParameters(input *track1.HealthProbeParameters) []interface{} {
	results := make([]interface{}, 0)
	if input == nil {
		return results
	}

	result := make(map[string]interface{})

	if input.ProbeIntervalInSeconds != nil {
		result["interval_in_seconds"] = *input.ProbeIntervalInSeconds
	}

	if input.ProbePath != nil {
		result["path"] = *input.ProbePath
	}

	if input.ProbeProtocol != "" {
		result["protocol"] = input.ProbeProtocol
	}

	if input.ProbeRequestType != "" {
		result["request_type"] = input.ProbeRequestType
	}

	return append(results, result)
}
