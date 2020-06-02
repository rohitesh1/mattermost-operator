// +build !ignore_autogenerated

// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallation":       schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec":   schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDB":       schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDB(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBSpec":   schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDBSpec(ref),
		"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBStatus": schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDBStatus(ref),
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallation is the Schema for the clusterinstallations API",
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
							Description: "Specification of the desired behavior of the Mattermost cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recent observed status of the Mattermost cluster. Read-only. Not included when requesting from the apiserver, only from the Mattermost Operator API itself. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationSpec", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ClusterInstallationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_ClusterInstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterInstallationSpec defines the desired state of ClusterInstallation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image defines the ClusterInstallation Docker image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version defines the ClusterInstallation Docker image version.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Size defines the size of the ClusterInstallation. This is typically specified in number of users. This will set replica and resource requests/limits appropriately for the provided number of users. Accepted values are: 100users, 1000users, 5000users, 10000users, 250000users. Defaults to 5000users. Setting 'Replicas', 'Resources', 'Minio.Replicas', 'Minio.Resource', 'Database.Replicas', or 'Database.Resources' will override the values set by Size.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas defines the number of replicas to use for the Mattermost app servers. Setting this will override the number of replicas set by 'Size'.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the resource requests and limits for the Mattermost app server pods.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"ingressName": {
						SchemaProps: spec.SchemaProps{
							Description: "IngressName defines the name to be used when creating the ingress rules",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mattermostLicenseSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret that contains the mattermost license",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, affinity will define the pod's scheduling constraints",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"minio": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Minio"),
						},
					},
					"database": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Database"),
						},
					},
					"blueGreen": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.BlueGreen"),
						},
					},
					"canary": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Canary"),
						},
					},
					"elasticSearch": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ElasticSearch"),
						},
					},
					"useServiceLoadBalancer": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"serviceAnnotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"useIngressTLS": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"resourceLabels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"ingressAnnotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"mattermostEnv": {
						SchemaProps: spec.SchemaProps{
							Description: "Optional environment variables to set in the Mattermost application pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"livenessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the probe to check if the application is up and running.",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Defines the probe to check if the application is ready to accept traffic.",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"certificates": {
						SchemaProps: spec.SchemaProps{
							Description: "CACertificates tells the operator which secret contains the certificate PEM file(s). Each file is mounted separatelly as a volume inside Mattermost container filesystem. The secret should contain a map where the key defines the filename and the value defines the PEM data. For example: <certificate.pem>/<certificate data>",
							Ref:         ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.CACertificates"),
						},
					},
				},
				Required: []string{"ingressName"},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.BlueGreen", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.CACertificates", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Canary", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Database", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.ElasticSearch", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.Minio", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Probe", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDB(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MattermostRestoreDB is the Schema for the mattermostrestoredbs API",
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
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBSpec", "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1.MattermostRestoreDBStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDBSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MattermostRestoreDBSpec defines the desired state of MattermostRestoreDB",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mattermostClusterName": {
						SchemaProps: spec.SchemaProps{
							Description: "MattermostClusterName defines the ClusterInstallation name.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"restoreSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "RestoreSecret defines the secret that holds the credentials to MySQL Operator be able to download the DB backup file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"initBucketURL": {
						SchemaProps: spec.SchemaProps{
							Description: "InitBucketURL defines where the DB backup file is located.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mattermostDBUser": {
						SchemaProps: spec.SchemaProps{
							Description: "MattermostDBUser defines the user to access the database. Need to set if the user is different from `mmuser`.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mattermostDBPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "MattermostDBPassword defines the user password to access the database. Need to set if the user is different from the one created by the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mattermostDBName": {
						SchemaProps: spec.SchemaProps{
							Description: "MattermostDBName defines the database name. Need to set if different from `mattermost`.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_mattermost_v1alpha1_MattermostRestoreDBStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MattermostRestoreDBStatus defines the observed state of MattermostRestoreDB",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents the state of the Mattermost restore Database.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"originalDBReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "The original number of database replicas. will be used to restore after applying the db restore process.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
	}
}
