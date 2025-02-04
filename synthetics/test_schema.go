package synthetics

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO(dfurman): mark required attributes, when they are specified in the OpenAPI definitions
// TODO(dfurman): provide descriptions, when they are specified in the OpenAPI definitions

func makeTestSchema(mode schemaMode) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Required: requiredOnReadSingle(mode),
			Computed: computedOnCreateAndReadList(mode),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"status": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			// enumeration: "TEST_STATUS_UNSPECIFIED", TEST_STATUS_ACTIVE", "TEST_STATUS_PAUSED",
			// "TEST_STATUS_DELETED"
		},
		"settings": makeTestSettingsSchema(mode),
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cdate": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"edate": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_by":      makeUserInfoSchema(),
		"last_updated_by": makeUserInfoSchema(),
	}
}

//nolint: funlen
func makeTestSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"hostname": makeTestHostnameSchema(mode),
		"ip":       makeTestIPSchema(mode),
		"agent":    makeTestAgentSchema(mode),
		"flow":     makeTestFlowSchema(mode),
		"site":     makeTestSiteSchema(mode),
		"tag":      makeTestTagSchema(mode),
		"dns":      makeTestDNSSchema(mode),
		"url":      makeTestURLSchema(mode),
		"agent_ids": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"period": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"limit": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"tasks": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"health_settings":     makeTestHealthSettingsSchema(mode),
		"monitoring_settings": makeTestMonitoringSettingsSchema(mode),
		"ping":                makeTestPingSchema(mode),
		"trace":               makeTestTraceSchema(mode),
		"port": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"protocol": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"family": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			// enumeration: "IP_FAMILY_UNSPECIFIED", "IP_FAMILY_V4", "IP_FAMILY_V6", "IP_FAMILY_DUAL"
		},
		"servers": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"use_local_ip": {
			Type:     schema.TypeBool,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"reciprocal": {
			Type:     schema.TypeBool,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"rollup_level": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestHostnameSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestIPSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"targets": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestAgentSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestFlowSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"target_refresh_interval_millis": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"max_tasks": {
			Type:     schema.TypeInt,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"type": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestSiteSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestTagSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestDNSSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestURLSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"target": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestHealthSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"latency_critical": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"latency_warning": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"packet_loss_critical": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"packet_loss_warning": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"jitter_critical": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"jitter_warning": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"http_latency_critical": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"http_latency_warning": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"http_valid_codes": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"dns_valid_codes": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	})
}

func makeTestMonitoringSettingsSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"activation_grace_period": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"activation_time_unit": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"activation_time_window": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"activation_times": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"notification_channels": {
			Type:     schema.TypeList,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	})
}

func makeTestPingSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeTestTraceSchema(mode schemaMode) *schema.Schema {
	return makeNestedObjectSchema(mode, map[string]*schema.Schema{
		"period": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"count": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"protocol": {
			Type:     schema.TypeString,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"port": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"expiry": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
		"limit": {
			Type:     schema.TypeFloat,
			Optional: optionalOnCreate(mode),
			Computed: computedOnRead(mode),
		},
	})
}

func makeUserInfoSchema() *schema.Schema {
	return makeReadOnlyNestedObjectSchema(map[string]*schema.Schema{
		idKey: {
			Type:     schema.TypeString,
			Computed: true,
		},
		"email": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"full_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	})
}
