// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Specifies the EBS volume upgrade information. The broker identifier must
// be set to the keyword ALL. This means the changes apply to all the brokers
// in the cluster.
type BrokerEBSVolumeInfo struct {
	KafkaBrokerNodeID *string `json:"kafkaBrokerNodeID,omitempty"`
	// Contains information about provisioned throughput for EBS storage volumes
	// attached to kafka broker nodes.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
	VolumeSizeGB          *int64                 `json:"volumeSizeGB,omitempty"`
}

// The broker logs configuration for this MSK cluster.
type BrokerLogs struct {
	// Details of the CloudWatch Logs destination for broker logs.
	CloudWatchLogs *CloudWatchLogs `json:"cloudWatchLogs,omitempty"`
	// Firehose details for BrokerLogs.
	Firehose *Firehose `json:"firehose,omitempty"`
	// The details of the Amazon S3 destination for broker logs.
	S3 *S3 `json:"s3,omitempty"`
}

// Describes the setup to be used for Apache Kafka broker nodes in the cluster.
type BrokerNodeGroupInfo struct {
	// The distribution of broker nodes across Availability Zones. By default, broker
	// nodes are distributed among the Availability Zones of your Region. Currently,
	// the only supported value is DEFAULT. You can either specify this value explicitly
	// or leave it out.
	BrokerAZDistribution *string   `json:"brokerAZDistribution,omitempty"`
	ClientSubnets        []*string `json:"clientSubnets,omitempty"`
	// Information about the broker access configuration.
	ConnectivityInfo *ConnectivityInfo `json:"connectivityInfo,omitempty"`
	InstanceType     *string           `json:"instanceType,omitempty"`
	SecurityGroups   []*string         `json:"securityGroups,omitempty"`
	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo *StorageInfo `json:"storageInfo,omitempty"`
}

// BrokerNodeInfo
type BrokerNodeInfo struct {
	AttachedENIID      *string `json:"attachedENIID,omitempty"`
	ClientSubnet       *string `json:"clientSubnet,omitempty"`
	ClientVPCIPAddress *string `json:"clientVPCIPAddress,omitempty"`
	// Information about the current software installed on the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo `json:"currentBrokerSoftwareInfo,omitempty"`
	Endpoints                 []*string           `json:"endpoints,omitempty"`
}

// Information about the current software installed on the cluster.
type BrokerSoftwareInfo struct {
	ConfigurationARN      *string `json:"configurationARN,omitempty"`
	ConfigurationRevision *int64  `json:"configurationRevision,omitempty"`
	KafkaVersion          *string `json:"kafkaVersion,omitempty"`
}

// Includes all client authentication information.
type ClientAuthentication struct {
	Sasl *Sasl `json:"sasl,omitempty"`
	// Details for client authentication using TLS.
	TLS *TLS `json:"tls,omitempty"`
	// Contains information about unauthenticated traffic to the cluster.
	Unauthenticated *Unauthenticated `json:"unauthenticated,omitempty"`
}

// Details of the CloudWatch Logs destination for broker logs.
type CloudWatchLogs struct {
	Enabled  *bool   `json:"enabled,omitempty"`
	LogGroup *string `json:"logGroup,omitempty"`
}

// Returns information about a cluster.
type ClusterInfo struct {
	ActiveOperationARN *string `json:"activeOperationARN,omitempty"`
	// Describes the setup to be used for Apache Kafka broker nodes in the cluster.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `json:"brokerNodeGroupInfo,omitempty"`
	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication `json:"clientAuthentication,omitempty"`
	ClusterARN           *string               `json:"clusterARN,omitempty"`
	ClusterName          *string               `json:"clusterName,omitempty"`
	CreationTime         *metav1.Time          `json:"creationTime,omitempty"`
	// Information about the current software installed on the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo `json:"currentBrokerSoftwareInfo,omitempty"`
	CurrentVersion            *string             `json:"currentVersion,omitempty"`
	// Includes encryption-related information, such as the AWS KMS key used for
	// encrypting data at rest and whether you want MSK to encrypt your data in
	// transit.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Specifies which metrics are gathered for the MSK cluster. This property has
	// the following possible values: DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER,
	// and PER_TOPIC_PER_PARTITION. For a list of the metrics associated with each
	// of these levels of monitoring, see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html).
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty"`
	// You can configure your MSK cluster to send broker logs to different destination
	// types. This is a container for the configuration details related to broker
	// logs.
	LoggingInfo         *LoggingInfo `json:"loggingInfo,omitempty"`
	NumberOfBrokerNodes *int64       `json:"numberOfBrokerNodes,omitempty"`
	// JMX and Node monitoring for the MSK cluster.
	OpenMonitoring *OpenMonitoring `json:"openMonitoring,omitempty"`
	// The state of an Apache Kafka cluster.
	State *string `json:"state,omitempty"`
	// Contains information about the state of the Amazon MSK cluster.
	StateInfo *StateInfo `json:"stateInfo,omitempty"`
	// Controls storage mode for various supported storage tiers.
	StorageMode               *string            `json:"storageMode,omitempty"`
	Tags                      map[string]*string `json:"tags,omitempty"`
	ZookeeperConnectString    *string            `json:"zookeeperConnectString,omitempty"`
	ZookeeperConnectStringTLS *string            `json:"zookeeperConnectStringTLS,omitempty"`
}

// Returns information about a cluster operation.
type ClusterOperationInfo struct {
	ClientRequestID *string      `json:"clientRequestID,omitempty"`
	ClusterARN      *string      `json:"clusterARN,omitempty"`
	CreationTime    *metav1.Time `json:"creationTime,omitempty"`
	EndTime         *metav1.Time `json:"endTime,omitempty"`
	OperationARN    *string      `json:"operationARN,omitempty"`
	OperationState  *string      `json:"operationState,omitempty"`
	OperationType   *string      `json:"operationType,omitempty"`
}

// Step taken during a cluster operation.
type ClusterOperationStep struct {
	StepName *string `json:"stepName,omitempty"`
}

// State information about the operation step.
type ClusterOperationStepInfo struct {
	StepStatus *string `json:"stepStatus,omitempty"`
}

// Returns information about a cluster of either the provisioned or the serverless
// type.
type Cluster_SDK struct {
	ActiveOperationARN *string      `json:"activeOperationARN,omitempty"`
	ClusterARN         *string      `json:"clusterARN,omitempty"`
	ClusterName        *string      `json:"clusterName,omitempty"`
	CreationTime       *metav1.Time `json:"creationTime,omitempty"`
	CurrentVersion     *string      `json:"currentVersion,omitempty"`
	// The state of an Apache Kafka cluster.
	State *string `json:"state,omitempty"`
	// Contains information about the state of the Amazon MSK cluster.
	StateInfo *StateInfo         `json:"stateInfo,omitempty"`
	Tags      map[string]*string `json:"tags,omitempty"`
}

// Contains source Apache Kafka versions and compatible target Apache Kafka
// versions.
type CompatibleKafkaVersion struct {
	SourceVersion  *string   `json:"sourceVersion,omitempty"`
	TargetVersions []*string `json:"targetVersions,omitempty"`
}

// Represents an MSK Configuration.
type Configuration struct {
	ARN           *string      `json:"arn,omitempty"`
	CreationTime  *metav1.Time `json:"creationTime,omitempty"`
	Description   *string      `json:"description,omitempty"`
	KafkaVersions []*string    `json:"kafkaVersions,omitempty"`
	Name          *string      `json:"name,omitempty"`
}

// Specifies the configuration to use for the brokers.
type ConfigurationInfo struct {
	ARN      *string `json:"arn,omitempty"`
	Revision *int64  `json:"revision,omitempty"`
}

// Describes a configuration revision.
type ConfigurationRevision struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Revision     *int64       `json:"revision,omitempty"`
}

// Information about the broker access configuration.
type ConnectivityInfo struct {
	// Broker public access control.
	PublicAccess *PublicAccess `json:"publicAccess,omitempty"`
}

// Contains information about the EBS storage volumes attached to Apache Kafka
// broker nodes.
type EBSStorageInfo struct {
	// Contains information about provisioned throughput for EBS storage volumes
	// attached to kafka broker nodes.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
	VolumeSize            *int64                 `json:"volumeSize,omitempty"`
}

// The data-volume encryption details.
type EncryptionAtRest struct {
	DataVolumeKMSKeyID *string `json:"dataVolumeKMSKeyID,omitempty"`
}

// The settings for encrypting data in transit.
type EncryptionInTransit struct {
	// Client-broker encryption in transit setting.
	ClientBroker *string `json:"clientBroker,omitempty"`
	InCluster    *bool   `json:"inCluster,omitempty"`
}

// Includes encryption-related information, such as the AWS KMS key used for
// encrypting data at rest and whether you want MSK to encrypt your data in
// transit.
type EncryptionInfo struct {
	// The data-volume encryption details.
	EncryptionAtRest *EncryptionAtRest `json:"encryptionAtRest,omitempty"`
	// The settings for encrypting data in transit.
	EncryptionInTransit *EncryptionInTransit `json:"encryptionInTransit,omitempty"`
}

// Returns information about an error state of the cluster.
type ErrorInfo struct {
	ErrorCode   *string `json:"errorCode,omitempty"`
	ErrorString *string `json:"errorString,omitempty"`
}

// Firehose details for BrokerLogs.
type Firehose struct {
	DeliveryStream *string `json:"deliveryStream,omitempty"`
	Enabled        *bool   `json:"enabled,omitempty"`
}

type IAM struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// Indicates whether you want to enable or disable the JMX Exporter.
type JmxExporter struct {
	EnabledInBroker *bool `json:"enabledInBroker,omitempty"`
}

// Indicates whether you want to enable or disable the JMX Exporter.
type JmxExporterInfo struct {
	EnabledInBroker *bool `json:"enabledInBroker,omitempty"`
}

// Information about a Apache Kafka version.
type KafkaVersion struct {
	Version *string `json:"version,omitempty"`
}

// You can configure your MSK cluster to send broker logs to different destination
// types. This is a container for the configuration details related to broker
// logs.
type LoggingInfo struct {
	// The broker logs configuration for this MSK cluster.
	BrokerLogs *BrokerLogs `json:"brokerLogs,omitempty"`
}

// Information about cluster attributes that can be updated via update APIs.
type MutableClusterInfo struct {
	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication `json:"clientAuthentication,omitempty"`
	// Specifies the configuration to use for the brokers.
	ConfigurationInfo *ConfigurationInfo `json:"configurationInfo,omitempty"`
	// Information about the broker access configuration.
	ConnectivityInfo *ConnectivityInfo `json:"connectivityInfo,omitempty"`
	// Includes encryption-related information, such as the AWS KMS key used for
	// encrypting data at rest and whether you want MSK to encrypt your data in
	// transit.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Specifies which metrics are gathered for the MSK cluster. This property has
	// the following possible values: DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER,
	// and PER_TOPIC_PER_PARTITION. For a list of the metrics associated with each
	// of these levels of monitoring, see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html).
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty"`
	InstanceType       *string `json:"instanceType,omitempty"`
	KafkaVersion       *string `json:"kafkaVersion,omitempty"`
	// You can configure your MSK cluster to send broker logs to different destination
	// types. This is a container for the configuration details related to broker
	// logs.
	LoggingInfo         *LoggingInfo `json:"loggingInfo,omitempty"`
	NumberOfBrokerNodes *int64       `json:"numberOfBrokerNodes,omitempty"`
	// JMX and Node monitoring for the MSK cluster.
	OpenMonitoring *OpenMonitoring `json:"openMonitoring,omitempty"`
	// Controls storage mode for various supported storage tiers.
	StorageMode *string `json:"storageMode,omitempty"`
}

// Indicates whether you want to enable or disable the Node Exporter.
type NodeExporter struct {
	EnabledInBroker *bool `json:"enabledInBroker,omitempty"`
}

// Indicates whether you want to enable or disable the Node Exporter.
type NodeExporterInfo struct {
	EnabledInBroker *bool `json:"enabledInBroker,omitempty"`
}

// The node information object.
type NodeInfo struct {
	AddedToClusterTime *string `json:"addedToClusterTime,omitempty"`
	InstanceType       *string `json:"instanceType,omitempty"`
	NodeARN            *string `json:"nodeARN,omitempty"`
}

// JMX and Node monitoring for the MSK cluster.
type OpenMonitoring struct {
	// Prometheus settings for open monitoring.
	Prometheus *Prometheus `json:"prometheus,omitempty"`
}

// JMX and Node monitoring for the MSK cluster.
type OpenMonitoringInfo struct {
	// Prometheus settings.
	Prometheus *PrometheusInfo `json:"prometheus,omitempty"`
}

// Prometheus settings for open monitoring.
type Prometheus struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter *JmxExporter `json:"jmxExporter,omitempty"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter *NodeExporter `json:"nodeExporter,omitempty"`
}

// Prometheus settings.
type PrometheusInfo struct {
	// Indicates whether you want to enable or disable the JMX Exporter.
	JmxExporter *JmxExporterInfo `json:"jmxExporter,omitempty"`
	// Indicates whether you want to enable or disable the Node Exporter.
	NodeExporter *NodeExporterInfo `json:"nodeExporter,omitempty"`
}

// Describes the provisioned cluster.
type Provisioned struct {
	// Describes the setup to be used for Apache Kafka broker nodes in the cluster.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `json:"brokerNodeGroupInfo,omitempty"`
	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication `json:"clientAuthentication,omitempty"`
	// Information about the current software installed on the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo `json:"currentBrokerSoftwareInfo,omitempty"`
	// Includes encryption-related information, such as the AWS KMS key used for
	// encrypting data at rest and whether you want MSK to encrypt your data in
	// transit.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Specifies which metrics are gathered for the MSK cluster. This property has
	// the following possible values: DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER,
	// and PER_TOPIC_PER_PARTITION. For a list of the metrics associated with each
	// of these levels of monitoring, see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html).
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty"`
	// You can configure your MSK cluster to send broker logs to different destination
	// types. This is a container for the configuration details related to broker
	// logs.
	LoggingInfo         *LoggingInfo `json:"loggingInfo,omitempty"`
	NumberOfBrokerNodes *int64       `json:"numberOfBrokerNodes,omitempty"`
	// JMX and Node monitoring for the MSK cluster.
	OpenMonitoring *OpenMonitoringInfo `json:"openMonitoring,omitempty"`
	// Controls storage mode for various supported storage tiers.
	StorageMode               *string `json:"storageMode,omitempty"`
	ZookeeperConnectString    *string `json:"zookeeperConnectString,omitempty"`
	ZookeeperConnectStringTLS *string `json:"zookeeperConnectStringTLS,omitempty"`
}

// Creates a provisioned cluster.
type ProvisionedRequest struct {
	// Describes the setup to be used for Apache Kafka broker nodes in the cluster.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `json:"brokerNodeGroupInfo,omitempty"`
	// Includes all client authentication information.
	ClientAuthentication *ClientAuthentication `json:"clientAuthentication,omitempty"`
	// Specifies the configuration to use for the brokers.
	ConfigurationInfo *ConfigurationInfo `json:"configurationInfo,omitempty"`
	// Includes encryption-related information, such as the AWS KMS key used for
	// encrypting data at rest and whether you want MSK to encrypt your data in
	// transit.
	EncryptionInfo *EncryptionInfo `json:"encryptionInfo,omitempty"`
	// Specifies which metrics are gathered for the MSK cluster. This property has
	// the following possible values: DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER,
	// and PER_TOPIC_PER_PARTITION. For a list of the metrics associated with each
	// of these levels of monitoring, see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html).
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty"`
	KafkaVersion       *string `json:"kafkaVersion,omitempty"`
	// You can configure your MSK cluster to send broker logs to different destination
	// types. This is a container for the configuration details related to broker
	// logs.
	LoggingInfo         *LoggingInfo `json:"loggingInfo,omitempty"`
	NumberOfBrokerNodes *int64       `json:"numberOfBrokerNodes,omitempty"`
	// JMX and Node monitoring for the MSK cluster.
	OpenMonitoring *OpenMonitoringInfo `json:"openMonitoring,omitempty"`
	// Controls storage mode for various supported storage tiers.
	StorageMode *string `json:"storageMode,omitempty"`
}

// Contains information about provisioned throughput for EBS storage volumes
// attached to kafka broker nodes.
type ProvisionedThroughput struct {
	Enabled          *bool  `json:"enabled,omitempty"`
	VolumeThroughput *int64 `json:"volumeThroughput,omitempty"`
}

// Broker public access control.
type PublicAccess struct {
	Type *string `json:"type_,omitempty"`
}

// The details of the Amazon S3 destination for broker logs.
type S3 struct {
	Bucket  *string `json:"bucket,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	Prefix  *string `json:"prefix,omitempty"`
}

type Sasl struct {
	IAM   *IAM   `json:"iam,omitempty"`
	Scram *Scram `json:"scram,omitempty"`
}

type Scram struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// Describes the serverless cluster SASL information.
type ServerlessSasl struct {
	IAM *IAM `json:"iam,omitempty"`
}

// Contains information about the state of the Amazon MSK cluster.
type StateInfo struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Contains information about storage volumes attached to MSK broker nodes.
type StorageInfo struct {
	// Contains information about the EBS storage volumes attached to Apache Kafka
	// broker nodes.
	EBSStorageInfo *EBSStorageInfo `json:"ebsStorageInfo,omitempty"`
}

// Details for client authentication using TLS.
type TLS struct {
	CertificateAuthorityARNList []*string `json:"certificateAuthorityARNList,omitempty"`
	Enabled                     *bool     `json:"enabled,omitempty"`
}

// Contains information about unauthenticated traffic to the cluster.
type Unauthenticated struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type UnprocessedScramSecret struct {
	ErrorCode    *string `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	SecretARN    *string `json:"secretARN,omitempty"`
}

// The configuration of the Amazon VPCs for the cluster.
type VPCConfig struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	SubnetIDs        []*string `json:"subnetIDs,omitempty"`
}

// Zookeeper node information.
type ZookeeperNodeInfo struct {
	AttachedENIID      *string   `json:"attachedENIID,omitempty"`
	ClientVPCIPAddress *string   `json:"clientVPCIPAddress,omitempty"`
	Endpoints          []*string `json:"endpoints,omitempty"`
	ZookeeperVersion   *string   `json:"zookeeperVersion,omitempty"`
}
