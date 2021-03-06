// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/capabilities/v1beta1.Product":       schema_pkg_apis_capabilities_v1beta1_Product(ref),
		"./pkg/apis/capabilities/v1beta1.ProductSpec":   schema_pkg_apis_capabilities_v1beta1_ProductSpec(ref),
		"./pkg/apis/capabilities/v1beta1.ProductStatus": schema_pkg_apis_capabilities_v1beta1_ProductStatus(ref),
	}
}

func schema_pkg_apis_capabilities_v1beta1_Product(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Product is the Schema for the products API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/capabilities/v1beta1.ProductSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/capabilities/v1beta1.ProductStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/capabilities/v1beta1.ProductSpec", "./pkg/apis/capabilities/v1beta1.ProductStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_capabilities_v1beta1_ProductSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ProductSpec defines the desired state of Product",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is human readable name for the product",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"systemName": {
						SchemaProps: spec.SchemaProps{
							Description: "SystemName identifies uniquely the product within the account provider Default value will be sanitized Name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a human readable text of the product",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"deployment": {
						SchemaProps: spec.SchemaProps{
							Description: "Deployment defined 3scale product deployment mode",
							Ref:         ref("./pkg/apis/capabilities/v1beta1.ProductDeploymentSpec"),
						},
					},
					"mappingRules": {
						SchemaProps: spec.SchemaProps{
							Description: "Mapping Rules Array: MappingRule Spec",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/capabilities/v1beta1.MappingRuleSpec"),
									},
								},
							},
						},
					},
					"backendUsages": {
						SchemaProps: spec.SchemaProps{
							Description: "Backend usage will be a map of Map: system_name -> BackendUsageSpec Having system_name as the index, the structure ensures one backend is not used multiple times.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/capabilities/v1beta1.BackendUsageSpec"),
									},
								},
							},
						},
					},
					"metrics": {
						SchemaProps: spec.SchemaProps{
							Description: "Metrics Map: system_name -> MetricSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/capabilities/v1beta1.MetricSpec"),
									},
								},
							},
						},
					},
					"methods": {
						SchemaProps: spec.SchemaProps{
							Description: "Methods Map: system_name -> MethodSpec system_name attr is unique for all metrics AND methods In other words, if metric's system_name is A, there is no metric or method with system_name A.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/capabilities/v1beta1.MethodSpec"),
									},
								},
							},
						},
					},
					"applicationPlans": {
						SchemaProps: spec.SchemaProps{
							Description: "Application Plans Map: system_name -> Application Plan Spec",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/capabilities/v1beta1.ApplicationPlanSpec"),
									},
								},
							},
						},
					},
					"providerAccountRef": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderAccountRef references account provider credentials",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/capabilities/v1beta1.ApplicationPlanSpec", "./pkg/apis/capabilities/v1beta1.BackendUsageSpec", "./pkg/apis/capabilities/v1beta1.MappingRuleSpec", "./pkg/apis/capabilities/v1beta1.MethodSpec", "./pkg/apis/capabilities/v1beta1.MetricSpec", "./pkg/apis/capabilities/v1beta1.ProductDeploymentSpec", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}

func schema_pkg_apis_capabilities_v1beta1_ProductStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ProductStatus defines the observed state of Product",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"productId": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration reflects the generation of the most recently observed Product Spec.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Current state of the 3scale product. Conditions represent the latest available observations of an object's state",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/3scale/3scale-operator/pkg/common.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/3scale/3scale-operator/pkg/common.Condition"},
	}
}
