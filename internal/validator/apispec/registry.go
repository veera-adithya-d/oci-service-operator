package apispec

import (
	"reflect"

	artifactsv1beta1 "github.com/oracle/oci-service-operator/api/artifacts/v1beta1"
	certificatesv1beta1 "github.com/oracle/oci-service-operator/api/certificates/v1beta1"
	certificatesmanagementv1beta1 "github.com/oracle/oci-service-operator/api/certificatesmanagement/v1beta1"
	containerenginev1beta1 "github.com/oracle/oci-service-operator/api/containerengine/v1beta1"
	corev1beta1 "github.com/oracle/oci-service-operator/api/core/v1beta1"
	databasev1beta1 "github.com/oracle/oci-service-operator/api/database/v1beta1"
	dnsv1beta1 "github.com/oracle/oci-service-operator/api/dns/v1beta1"
	eventsv1beta1 "github.com/oracle/oci-service-operator/api/events/v1beta1"
	functionsv1beta1 "github.com/oracle/oci-service-operator/api/functions/v1beta1"
	identityv1beta1 "github.com/oracle/oci-service-operator/api/identity/v1beta1"
	keymanagementv1beta1 "github.com/oracle/oci-service-operator/api/keymanagement/v1beta1"
	limitsv1beta1 "github.com/oracle/oci-service-operator/api/limits/v1beta1"
	loadbalancerv1beta1 "github.com/oracle/oci-service-operator/api/loadbalancer/v1beta1"
	loggingv1beta1 "github.com/oracle/oci-service-operator/api/logging/v1beta1"
	monitoringv1beta1 "github.com/oracle/oci-service-operator/api/monitoring/v1beta1"
	mysqlv1beta1 "github.com/oracle/oci-service-operator/api/mysql/v1beta1"
	networkloadbalancerv1beta1 "github.com/oracle/oci-service-operator/api/networkloadbalancer/v1beta1"
	nosqlv1beta1 "github.com/oracle/oci-service-operator/api/nosql/v1beta1"
	objectstoragev1beta1 "github.com/oracle/oci-service-operator/api/objectstorage/v1beta1"
	onsv1beta1 "github.com/oracle/oci-service-operator/api/ons/v1beta1"
	psqlv1beta1 "github.com/oracle/oci-service-operator/api/psql/v1beta1"
	queuev1beta1 "github.com/oracle/oci-service-operator/api/queue/v1beta1"
	secretsv1beta1 "github.com/oracle/oci-service-operator/api/secrets/v1beta1"
	streamingv1beta1 "github.com/oracle/oci-service-operator/api/streaming/v1beta1"
	vaultv1beta1 "github.com/oracle/oci-service-operator/api/vault/v1beta1"
	workrequestsv1beta1 "github.com/oracle/oci-service-operator/api/workrequests/v1beta1"
)

type Target struct {
	Name       string
	SpecType   reflect.Type
	StatusType reflect.Type
	SDKStructs []string
}

var targets = []Target{
	{
		Name:       "AutonomousDatabases",
		SpecType:   reflect.TypeOf(databasev1beta1.AutonomousDatabasesSpec{}),
		StatusType: reflect.TypeOf(databasev1beta1.AutonomousDatabasesStatus{}),
		SDKStructs: []string{
			"database.CreateAutonomousDatabaseDetails",
			"database.UpdateAutonomousDatabaseDetails",
		},
	},
	{
		Name:       "MySqlDbSystem",
		SpecType:   reflect.TypeOf(mysqlv1beta1.MySqlDbSystemSpec{}),
		StatusType: reflect.TypeOf(mysqlv1beta1.MySqlDbSystemStatus{}),
		SDKStructs: []string{
			"mysql.CreateDbSystemDetails",
			"mysql.UpdateDbSystemDetails",
		},
	},
	{
		Name:       "Stream",
		SpecType:   reflect.TypeOf(streamingv1beta1.StreamSpec{}),
		StatusType: reflect.TypeOf(streamingv1beta1.StreamStatus{}),
		SDKStructs: []string{
			"streaming.CreateStreamDetails",
			"streaming.UpdateStreamDetails",
			"streaming.Stream",
			"streaming.StreamSummary",
		},
	},
	{
		Name:       "Channel",
		SpecType:   reflect.TypeOf(queuev1beta1.ChannelSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.ChannelStatus{}),
		SDKStructs: []string{
			"queue.ChannelCollection",
		},
	},
	{
		Name:       "Queue",
		SpecType:   reflect.TypeOf(queuev1beta1.QueueSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.QueueStatus{}),
		SDKStructs: []string{
			"queue.CreateQueueDetails",
			"queue.UpdateQueueDetails",
			"queue.Queue",
			"queue.QueueCollection",
			"queue.QueueSummary",
		},
	},
	{
		Name:       "QueueMessage",
		SpecType:   reflect.TypeOf(queuev1beta1.MessageSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.MessageStatus{}),
		SDKStructs: []string{
			"queue.UpdateMessageDetails",
		},
	},
	{
		Name:       "QueueWorkRequest",
		SpecType:   reflect.TypeOf(queuev1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"queue.WorkRequest",
			"queue.WorkRequestSummary",
		},
	},
	{
		Name:       "QueueWorkRequestError",
		SpecType:   reflect.TypeOf(queuev1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"queue.WorkRequestError",
			"queue.WorkRequestErrorCollection",
		},
	},
	{
		Name:       "Stats",
		SpecType:   reflect.TypeOf(queuev1beta1.StatsSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.StatsObservedState{}),
		SDKStructs: []string{
			"queue.Stats",
		},
	},
	{
		Name:       "WorkRequestLog",
		SpecType:   reflect.TypeOf(queuev1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(queuev1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"queue.WorkRequestLogEntry",
			"queue.WorkRequestLogEntryCollection",
		},
	},
	{
		Name:       "FunctionsApplication",
		SpecType:   reflect.TypeOf(functionsv1beta1.ApplicationSpec{}),
		StatusType: reflect.TypeOf(functionsv1beta1.ApplicationStatus{}),
		SDKStructs: []string{
			"functions.CreateApplicationDetails",
			"functions.UpdateApplicationDetails",
			"functions.Application",
			"functions.ApplicationSummary",
		},
	},
	{
		Name:       "FunctionsFunction",
		SpecType:   reflect.TypeOf(functionsv1beta1.FunctionSpec{}),
		StatusType: reflect.TypeOf(functionsv1beta1.FunctionStatus{}),
		SDKStructs: []string{
			"functions.CreateFunctionDetails",
			"functions.UpdateFunctionDetails",
			"functions.Function",
			"functions.FunctionSummary",
		},
	},
	{
		Name:       "FunctionsPbfListing",
		SpecType:   reflect.TypeOf(functionsv1beta1.PbfListingSpec{}),
		StatusType: reflect.TypeOf(functionsv1beta1.PbfListingStatus{}),
		SDKStructs: []string{
			"functions.PbfListing",
			"functions.PbfListingVersionSummary",
			"functions.PbfListingSummary",
		},
	},
	{
		Name:       "FunctionsPbfListingVersion",
		SpecType:   reflect.TypeOf(functionsv1beta1.PbfListingVersionSpec{}),
		StatusType: reflect.TypeOf(functionsv1beta1.PbfListingVersionStatus{}),
		SDKStructs: []string{
			"functions.PbfListingVersion",
			"functions.PbfListingVersionSummary",
		},
	},
	{
		Name:       "FunctionsTrigger",
		SpecType:   reflect.TypeOf(functionsv1beta1.TriggerSpec{}),
		StatusType: reflect.TypeOf(functionsv1beta1.TriggerStatus{}),
		SDKStructs: []string{
			"functions.Trigger",
			"functions.TriggerSummary",
		},
	},
	{
		Name:       "NoSQLIndex",
		SpecType:   reflect.TypeOf(nosqlv1beta1.IndexSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.IndexStatus{}),
		SDKStructs: []string{
			"nosql.CreateIndexDetails",
			"nosql.Index",
			"nosql.IndexCollection",
			"nosql.IndexSummary",
		},
	},
	{
		Name:       "NoSQLReplica",
		SpecType:   reflect.TypeOf(nosqlv1beta1.ReplicaSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.ReplicaStatus{}),
		SDKStructs: []string{
			"nosql.CreateReplicaDetails",
			"nosql.Replica",
		},
	},
	{
		Name:       "NoSQLRow",
		SpecType:   reflect.TypeOf(nosqlv1beta1.RowSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.RowStatus{}),
		SDKStructs: []string{
			"nosql.UpdateRowDetails",
			"nosql.Row",
		},
	},
	{
		Name:       "NoSQLTable",
		SpecType:   reflect.TypeOf(nosqlv1beta1.TableSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.TableStatus{}),
		SDKStructs: []string{
			"nosql.CreateTableDetails",
			"nosql.UpdateTableDetails",
			"nosql.Table",
			"nosql.TableCollection",
			"nosql.TableSummary",
		},
	},
	{
		Name:       "NoSQLTableUsage",
		SpecType:   reflect.TypeOf(nosqlv1beta1.TableUsageSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.TableUsageStatus{}),
		SDKStructs: []string{
			"nosql.TableUsageCollection",
			"nosql.TableUsageSummary",
		},
	},
	{
		Name:       "NoSQLWorkRequest",
		SpecType:   reflect.TypeOf(nosqlv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"nosql.WorkRequest",
			"nosql.WorkRequestCollection",
			"nosql.WorkRequestSummary",
		},
	},
	{
		Name:       "NoSQLWorkRequestError",
		SpecType:   reflect.TypeOf(nosqlv1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"nosql.WorkRequestError",
			"nosql.WorkRequestErrorCollection",
		},
	},
	{
		Name:       "NoSQLWorkRequestLog",
		SpecType:   reflect.TypeOf(nosqlv1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(nosqlv1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"nosql.WorkRequestLogEntry",
			"nosql.WorkRequestLogEntryCollection",
		},
	},
	{
		Name:       "ObjectStorageBucket",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.BucketSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.BucketStatus{}),
		SDKStructs: []string{
			"objectstorage.CreateBucketDetails",
			"objectstorage.UpdateBucketDetails",
			"objectstorage.Bucket",
			"objectstorage.BucketSummary",
		},
	},
	{
		Name:       "ObjectStorageMultipartUpload",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.MultipartUploadSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.MultipartUploadStatus{}),
		SDKStructs: []string{
			"objectstorage.CreateMultipartUploadDetails",
			"objectstorage.MultipartUpload",
		},
	},
	{
		Name:       "ObjectStorageMultipartUploadPart",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.MultipartUploadPartSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.MultipartUploadPartStatus{}),
		SDKStructs: []string{
			"objectstorage.MultipartUploadPartSummary",
		},
	},
	{
		Name:       "ObjectStorageNamespace",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.NamespaceSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.NamespaceStatus{}),
		SDKStructs: []string{
			"objectstorage.NamespaceMetadata",
		},
	},
	{
		Name:       "ObjectStorageNamespaceMetadata",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.NamespaceMetadataSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.NamespaceMetadataStatus{}),
		SDKStructs: []string{
			"objectstorage.UpdateNamespaceMetadataDetails",
			"objectstorage.NamespaceMetadata",
		},
	},
	{
		Name:       "ObjectStorageObject",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ObjectSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ObjectStatus{}),
		SDKStructs: []string{
			"objectstorage.ObjectVersionSummary",
			"objectstorage.ObjectSummary",
		},
	},
	{
		Name:       "ObjectStorageObjectLifecyclePolicy",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ObjectLifecyclePolicySpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ObjectLifecyclePolicyStatus{}),
		SDKStructs: []string{
			"objectstorage.ObjectLifecyclePolicy",
		},
	},
	{
		Name:       "ObjectStorageObjectStorageTier",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ObjectStorageTierSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ObjectStorageTierStatus{}),
		SDKStructs: []string{
			"objectstorage.UpdateObjectStorageTierDetails",
		},
	},
	{
		Name:       "ObjectStorageObjectVersion",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ObjectVersionSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ObjectVersionStatus{}),
		SDKStructs: []string{
			"objectstorage.ObjectVersionCollection",
			"objectstorage.ObjectVersionSummary",
		},
	},
	{
		Name:       "ObjectStoragePreauthenticatedRequest",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.PreauthenticatedRequestSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.PreauthenticatedRequestStatus{}),
		SDKStructs: []string{
			"objectstorage.CreatePreauthenticatedRequestDetails",
			"objectstorage.PreauthenticatedRequest",
			"objectstorage.PreauthenticatedRequestSummary",
		},
	},
	{
		Name:       "ObjectStorageReplicationPolicy",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ReplicationPolicySpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ReplicationPolicyStatus{}),
		SDKStructs: []string{
			"objectstorage.CreateReplicationPolicyDetails",
			"objectstorage.ReplicationPolicy",
			"objectstorage.ReplicationPolicySummary",
		},
	},
	{
		Name:       "ObjectStorageReplicationSource",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.ReplicationSourceSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.ReplicationSourceStatus{}),
		SDKStructs: []string{
			"objectstorage.ReplicationSource",
		},
	},
	{
		Name:       "ObjectStorageRetentionRule",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.RetentionRuleSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.RetentionRuleStatus{}),
		SDKStructs: []string{
			"objectstorage.CreateRetentionRuleDetails",
			"objectstorage.UpdateRetentionRuleDetails",
			"objectstorage.RetentionRuleDetails",
			"objectstorage.RetentionRule",
			"objectstorage.RetentionRuleCollection",
			"objectstorage.RetentionRuleSummary",
		},
	},
	{
		Name:       "ObjectStorageWorkRequest",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"objectstorage.WorkRequest",
			"objectstorage.WorkRequestSummary",
		},
	},
	{
		Name:       "ObjectStorageWorkRequestError",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"objectstorage.WorkRequestError",
		},
	},
	{
		Name:       "ObjectStorageWorkRequestLog",
		SpecType:   reflect.TypeOf(objectstoragev1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(objectstoragev1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"objectstorage.WorkRequestLogEntry",
		},
	},
	{
		Name:       "NotificationConfirmSubscription",
		SpecType:   reflect.TypeOf(onsv1beta1.ConfirmSubscriptionSpec{}),
		StatusType: reflect.TypeOf(onsv1beta1.ConfirmSubscriptionStatus{}),
		SDKStructs: []string{
			"ons.ConfirmationResult",
		},
	},
	{
		Name:       "NotificationTopic",
		SpecType:   reflect.TypeOf(onsv1beta1.TopicSpec{}),
		StatusType: reflect.TypeOf(onsv1beta1.TopicStatus{}),
		SDKStructs: []string{
			"ons.CreateTopicDetails",
			"ons.NotificationTopic",
			"ons.NotificationTopicSummary",
		},
	},
	{
		Name:       "NotificationUnsubscription",
		SpecType:   reflect.TypeOf(onsv1beta1.UnsubscriptionSpec{}),
		StatusType: reflect.TypeOf(onsv1beta1.UnsubscriptionStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "ONSSubscription",
		SpecType:   reflect.TypeOf(onsv1beta1.SubscriptionSpec{}),
		StatusType: reflect.TypeOf(onsv1beta1.SubscriptionStatus{}),
		SDKStructs: []string{
			"ons.CreateSubscriptionDetails",
			"ons.UpdateSubscriptionDetails",
			"ons.Subscription",
			"ons.SubscriptionSummary",
		},
	},
	{
		Name:       "LogGroup",
		SpecType:   reflect.TypeOf(loggingv1beta1.LogGroupSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.LogGroupStatus{}),
		SDKStructs: []string{
			"logging.CreateLogGroupDetails",
			"logging.UpdateLogGroupDetails",
			"logging.LogGroup",
			"logging.LogGroupSummary",
		},
	},
	{
		Name:       "LoggingLog",
		SpecType:   reflect.TypeOf(loggingv1beta1.LogSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.LogStatus{}),
		SDKStructs: []string{
			"logging.CreateLogDetails",
			"logging.UpdateLogDetails",
			"logging.Log",
			"logging.LogSummary",
		},
	},
	{
		Name:       "LoggingLogSavedSearch",
		SpecType:   reflect.TypeOf(loggingv1beta1.LogSavedSearchSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.LogSavedSearchStatus{}),
		SDKStructs: []string{
			"logging.CreateLogSavedSearchDetails",
			"logging.UpdateLogSavedSearchDetails",
			"logging.LogSavedSearch",
			"logging.LogSavedSearchSummary",
		},
	},
	{
		Name:       "LoggingService",
		SpecType:   reflect.TypeOf(loggingv1beta1.ServiceSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.ServiceStatus{}),
		SDKStructs: []string{
			"logging.ServiceSummary",
		},
	},
	{
		Name:       "LoggingUnifiedAgentConfiguration",
		SpecType:   reflect.TypeOf(loggingv1beta1.UnifiedAgentConfigurationSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.UnifiedAgentConfigurationStatus{}),
		SDKStructs: []string{
			"logging.CreateUnifiedAgentConfigurationDetails",
			"logging.UpdateUnifiedAgentConfigurationDetails",
			"logging.UnifiedAgentConfiguration",
			"logging.UnifiedAgentConfigurationCollection",
			"logging.UnifiedAgentConfigurationSummary",
		},
	},
	{
		Name:       "LoggingWorkRequest",
		SpecType:   reflect.TypeOf(loggingv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"logging.WorkRequest",
			"logging.WorkRequestSummary",
		},
	},
	{
		Name:       "LoggingWorkRequestError",
		SpecType:   reflect.TypeOf(loggingv1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"logging.WorkRequestError",
		},
	},
	{
		Name:       "LoggingWorkRequestLog",
		SpecType:   reflect.TypeOf(loggingv1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(loggingv1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"logging.WorkRequestLog",
		},
	},
	{
		Name:       "PSQLBackup",
		SpecType:   reflect.TypeOf(psqlv1beta1.BackupSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.BackupStatus{}),
		SDKStructs: []string{
			"psql.CreateBackupDetails",
			"psql.UpdateBackupDetails",
			"psql.Backup",
			"psql.BackupCollection",
			"psql.BackupSummary",
		},
	},
	{
		Name:       "PSQLConfiguration",
		SpecType:   reflect.TypeOf(psqlv1beta1.ConfigurationSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.ConfigurationStatus{}),
		SDKStructs: []string{
			"psql.CreateConfigurationDetails",
			"psql.UpdateConfigurationDetails",
			"psql.ConfigurationDetails",
			"psql.Configuration",
			"psql.ConfigurationCollection",
			"psql.ConfigurationSummary",
		},
	},
	{
		Name:       "PSQLConnectionDetail",
		SpecType:   reflect.TypeOf(psqlv1beta1.ConnectionDetailSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.ConnectionDetailStatus{}),
		SDKStructs: []string{
			"psql.ConnectionDetails",
		},
	},
	{
		Name:       "PSQLDbSystemDbInstance",
		SpecType:   reflect.TypeOf(psqlv1beta1.DbSystemDbInstanceSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.DbSystemDbInstanceStatus{}),
		SDKStructs: []string{
			"psql.UpdateDbSystemDbInstanceDetails",
		},
	},
	{
		Name:       "PSQLDefaultConfiguration",
		SpecType:   reflect.TypeOf(psqlv1beta1.DefaultConfigurationSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.DefaultConfigurationStatus{}),
		SDKStructs: []string{
			"psql.DefaultConfigurationDetails",
			"psql.DefaultConfiguration",
			"psql.DefaultConfigurationCollection",
			"psql.DefaultConfigurationSummary",
		},
	},
	{
		Name:       "PSQLPrimaryDbInstance",
		SpecType:   reflect.TypeOf(psqlv1beta1.PrimaryDbInstanceSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.PrimaryDbInstanceStatus{}),
		SDKStructs: []string{
			"psql.PrimaryDbInstanceDetails",
		},
	},
	{
		Name:       "PSQLShape",
		SpecType:   reflect.TypeOf(psqlv1beta1.ShapeSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.ShapeStatus{}),
		SDKStructs: []string{
			"psql.ShapeCollection",
			"psql.ShapeSummary",
		},
	},
	{
		Name:       "PSQLWorkRequest",
		SpecType:   reflect.TypeOf(psqlv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"psql.WorkRequest",
			"psql.WorkRequestSummary",
		},
	},
	{
		Name:       "PSQLWorkRequestError",
		SpecType:   reflect.TypeOf(psqlv1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"psql.WorkRequestError",
			"psql.WorkRequestErrorCollection",
		},
	},
	{
		Name:       "PSQLWorkRequestLog",
		SpecType:   reflect.TypeOf(psqlv1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"psql.WorkRequestLogEntry",
			"psql.WorkRequestLogEntryCollection",
		},
	},
	{
		Name:       "PostgreSQLDbSystem",
		SpecType:   reflect.TypeOf(psqlv1beta1.DbSystemSpec{}),
		StatusType: reflect.TypeOf(psqlv1beta1.DbSystemStatus{}),
		SDKStructs: []string{
			"psql.CreateDbSystemDetails",
			"psql.UpdateDbSystemDetails",
			"psql.DbSystemDetails",
			"psql.DbSystem",
			"psql.DbSystemCollection",
			"psql.DbSystemSummary",
		},
	},
	{
		Name:       "EventsRule",
		SpecType:   reflect.TypeOf(eventsv1beta1.RuleSpec{}),
		StatusType: reflect.TypeOf(eventsv1beta1.RuleStatus{}),
		SDKStructs: []string{
			"events.CreateRuleDetails",
			"events.UpdateRuleDetails",
			"events.Rule",
			"events.RuleSummary",
		},
	},
	{
		Name:       "MonitoringAlarm",
		SpecType:   reflect.TypeOf(monitoringv1beta1.AlarmSpec{}),
		StatusType: reflect.TypeOf(monitoringv1beta1.AlarmObservedState{}),
		SDKStructs: []string{
			"monitoring.CreateAlarmDetails",
			"monitoring.UpdateAlarmDetails",
			"monitoring.Alarm",
			"monitoring.AlarmSummary",
		},
	},
	{
		Name:       "MonitoringAlarmHistory",
		SpecType:   reflect.TypeOf(monitoringv1beta1.AlarmHistorySpec{}),
		StatusType: reflect.TypeOf(monitoringv1beta1.AlarmHistoryStatus{}),
		SDKStructs: []string{
			"monitoring.AlarmHistoryCollection",
			"monitoring.AlarmHistoryEntry",
		},
	},
	{
		Name:       "MonitoringAlarmStatus",
		SpecType:   reflect.TypeOf(monitoringv1beta1.AlarmStatusSpec{}),
		StatusType: reflect.TypeOf(monitoringv1beta1.AlarmStatusObservedState{}),
		SDKStructs: []string{
			"monitoring.AlarmStatusSummary",
		},
	},
	{
		Name:       "MonitoringAlarmSuppression",
		SpecType:   reflect.TypeOf(monitoringv1beta1.AlarmSuppressionSpec{}),
		StatusType: reflect.TypeOf(monitoringv1beta1.AlarmSuppressionStatus{}),
		SDKStructs: []string{
			"monitoring.CreateAlarmSuppressionDetails",
			"monitoring.AlarmSuppression",
			"monitoring.AlarmSuppressionCollection",
			"monitoring.AlarmSuppressionSummary",
		},
	},
	{
		Name:       "MonitoringMetric",
		SpecType:   reflect.TypeOf(monitoringv1beta1.MetricSpec{}),
		StatusType: reflect.TypeOf(monitoringv1beta1.MetricStatus{}),
		SDKStructs: []string{
			"monitoring.Metric",
		},
	},
	{
		Name:       "DNSDomainRecord",
		SpecType:   reflect.TypeOf(dnsv1beta1.DomainRecordSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.DomainRecordStatus{}),
		SDKStructs: []string{
			"dns.Record",
		},
	},
	{
		Name:       "DNSRRSet",
		SpecType:   reflect.TypeOf(dnsv1beta1.RRSetSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.RRSetStatus{}),
		SDKStructs: []string{
			"dns.UpdateRrSetDetails",
			"dns.RrSet",
		},
	},
	{
		Name:       "DNSResolver",
		SpecType:   reflect.TypeOf(dnsv1beta1.ResolverSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ResolverStatus{}),
		SDKStructs: []string{
			"dns.UpdateResolverDetails",
			"dns.Resolver",
			"dns.ResolverSummary",
		},
	},
	{
		Name:       "DNSResolverEndpoint",
		SpecType:   reflect.TypeOf(dnsv1beta1.ResolverEndpointSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ResolverEndpointStatus{}),
		SDKStructs: []string{
			"dns.ResolverVnicEndpoint",
			"dns.ResolverVnicEndpointSummary",
		},
	},
	{
		Name:       "DNSSteeringPolicy",
		SpecType:   reflect.TypeOf(dnsv1beta1.SteeringPolicySpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.SteeringPolicyStatus{}),
		SDKStructs: []string{
			"dns.CreateSteeringPolicyDetails",
			"dns.UpdateSteeringPolicyDetails",
			"dns.SteeringPolicy",
			"dns.SteeringPolicySummary",
		},
	},
	{
		Name:       "DNSSteeringPolicyAttachment",
		SpecType:   reflect.TypeOf(dnsv1beta1.SteeringPolicyAttachmentSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.SteeringPolicyAttachmentStatus{}),
		SDKStructs: []string{
			"dns.CreateSteeringPolicyAttachmentDetails",
			"dns.UpdateSteeringPolicyAttachmentDetails",
			"dns.SteeringPolicyAttachment",
			"dns.SteeringPolicyAttachmentSummary",
		},
	},
	{
		Name:       "DNSTsigKey",
		SpecType:   reflect.TypeOf(dnsv1beta1.TsigKeySpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.TsigKeyStatus{}),
		SDKStructs: []string{
			"dns.CreateTsigKeyDetails",
			"dns.UpdateTsigKeyDetails",
			"dns.TsigKey",
			"dns.TsigKeySummary",
		},
	},
	{
		Name:       "DNSView",
		SpecType:   reflect.TypeOf(dnsv1beta1.ViewSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ViewStatus{}),
		SDKStructs: []string{
			"dns.CreateViewDetails",
			"dns.UpdateViewDetails",
			"dns.View",
			"dns.ViewSummary",
		},
	},
	{
		Name:       "DNSZone",
		SpecType:   reflect.TypeOf(dnsv1beta1.ZoneSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ZoneStatus{}),
		SDKStructs: []string{
			"dns.CreateZoneDetails",
			"dns.UpdateZoneDetails",
			"dns.Zone",
			"dns.ZoneSummary",
		},
	},
	{
		Name:       "DNSZoneContent",
		SpecType:   reflect.TypeOf(dnsv1beta1.ZoneContentSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ZoneContentStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "DNSZoneFromZoneFile",
		SpecType:   reflect.TypeOf(dnsv1beta1.ZoneFromZoneFileSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ZoneFromZoneFileStatus{}),
		SDKStructs: []string{
			"dns.Zone",
		},
	},
	{
		Name:       "DNSZoneRecord",
		SpecType:   reflect.TypeOf(dnsv1beta1.ZoneRecordSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ZoneRecordStatus{}),
		SDKStructs: []string{
			"dns.Record",
		},
	},
	{
		Name:       "DNSZoneTransferServer",
		SpecType:   reflect.TypeOf(dnsv1beta1.ZoneTransferServerSpec{}),
		StatusType: reflect.TypeOf(dnsv1beta1.ZoneTransferServerStatus{}),
		SDKStructs: []string{
			"dns.ZoneTransferServer",
		},
	},
	{
		Name:       "LoadBalancer",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.LoadBalancerSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.LoadBalancerStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateLoadBalancerDetails",
			"loadbalancer.UpdateLoadBalancerDetails",
			"loadbalancer.LoadBalancer",
		},
	},
	{
		Name:       "LoadBalancerBackend",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.BackendSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.BackendStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateBackendDetails",
			"loadbalancer.UpdateBackendDetails",
			"loadbalancer.BackendDetails",
			"loadbalancer.Backend",
		},
	},
	{
		Name:       "LoadBalancerBackendHealth",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.BackendHealthSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.BackendHealthStatus{}),
		SDKStructs: []string{
			"loadbalancer.BackendHealth",
		},
	},
	{
		Name:       "LoadBalancerBackendSet",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.BackendSetSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.BackendSetStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateBackendSetDetails",
			"loadbalancer.UpdateBackendSetDetails",
			"loadbalancer.BackendSetDetails",
			"loadbalancer.BackendSet",
		},
	},
	{
		Name:       "LoadBalancerBackendSetHealth",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.BackendSetHealthSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.BackendSetHealthStatus{}),
		SDKStructs: []string{
			"loadbalancer.BackendSetHealth",
		},
	},
	{
		Name:       "LoadBalancerCertificate",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.CertificateSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.CertificateStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateCertificateDetails",
			"loadbalancer.CertificateDetails",
			"loadbalancer.Certificate",
		},
	},
	{
		Name:       "LoadBalancerHealthChecker",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.HealthCheckerSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.HealthCheckerStatus{}),
		SDKStructs: []string{
			"loadbalancer.UpdateHealthCheckerDetails",
			"loadbalancer.HealthCheckerDetails",
			"loadbalancer.HealthChecker",
		},
	},
	{
		Name:       "LoadBalancerHostname",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.HostnameSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.HostnameStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateHostnameDetails",
			"loadbalancer.UpdateHostnameDetails",
			"loadbalancer.HostnameDetails",
			"loadbalancer.Hostname",
		},
	},
	{
		Name:       "LoadBalancerListener",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.ListenerSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.ListenerStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateListenerDetails",
			"loadbalancer.UpdateListenerDetails",
			"loadbalancer.ListenerDetails",
			"loadbalancer.Listener",
		},
	},
	{
		Name:       "LoadBalancerListenerRule",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.ListenerRuleSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.ListenerRuleStatus{}),
		SDKStructs: []string{
			"loadbalancer.ListenerRuleSummary",
		},
	},
	{
		Name:       "LoadBalancerLoadBalancerHealth",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.LoadBalancerHealthSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.LoadBalancerHealthStatus{}),
		SDKStructs: []string{
			"loadbalancer.LoadBalancerHealth",
			"loadbalancer.LoadBalancerHealthSummary",
		},
	},
	{
		Name:       "LoadBalancerLoadBalancerShape",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.LoadBalancerShapeSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.LoadBalancerShapeStatus{}),
		SDKStructs: []string{
			"loadbalancer.UpdateLoadBalancerShapeDetails",
			"loadbalancer.LoadBalancerShape",
		},
	},
	{
		Name:       "LoadBalancerNetworkSecurityGroup",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.NetworkSecurityGroupSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.NetworkSecurityGroupStatus{}),
		SDKStructs: []string{
			"loadbalancer.UpdateNetworkSecurityGroupsDetails",
		},
	},
	{
		Name:       "LoadBalancerPathRouteSet",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.PathRouteSetSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.PathRouteSetStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreatePathRouteSetDetails",
			"loadbalancer.UpdatePathRouteSetDetails",
			"loadbalancer.PathRouteSetDetails",
			"loadbalancer.PathRouteSet",
		},
	},
	{
		Name:       "LoadBalancerPolicy",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.PolicySpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.PolicyStatus{}),
		SDKStructs: []string{
			"loadbalancer.LoadBalancerPolicy",
		},
	},
	{
		Name:       "LoadBalancerProtocol",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.ProtocolSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.ProtocolStatus{}),
		SDKStructs: []string{
			"loadbalancer.LoadBalancerProtocol",
		},
	},
	{
		Name:       "LoadBalancerRoutingPolicy",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.RoutingPolicySpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.RoutingPolicyStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateRoutingPolicyDetails",
			"loadbalancer.UpdateRoutingPolicyDetails",
			"loadbalancer.RoutingPolicyDetails",
			"loadbalancer.RoutingPolicy",
		},
	},
	{
		Name:       "LoadBalancerRuleSet",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.RuleSetSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.RuleSetStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateRuleSetDetails",
			"loadbalancer.UpdateRuleSetDetails",
			"loadbalancer.RuleSetDetails",
			"loadbalancer.RuleSet",
		},
	},
	{
		Name:       "LoadBalancerSSLCipherSuite",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.SSLCipherSuiteSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.SSLCipherSuiteStatus{}),
		SDKStructs: []string{
			"loadbalancer.CreateSslCipherSuiteDetails",
			"loadbalancer.UpdateSslCipherSuiteDetails",
			"loadbalancer.SslCipherSuiteDetails",
			"loadbalancer.SslCipherSuite",
		},
	},
	{
		Name:       "LoadBalancerShape",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.ShapeSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.ShapeStatus{}),
		SDKStructs: []string{
			"loadbalancer.UpdateLoadBalancerShapeDetails",
			"loadbalancer.ShapeDetails",
			"loadbalancer.LoadBalancerShape",
		},
	},
	{
		Name:       "LoadBalancerWorkRequest",
		SpecType:   reflect.TypeOf(loadbalancerv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(loadbalancerv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"loadbalancer.WorkRequest",
		},
	},
	{
		Name:       "NetworkLoadBalancer",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancerSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancerStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.CreateNetworkLoadBalancerDetails",
			"networkloadbalancer.UpdateNetworkLoadBalancerDetails",
			"networkloadbalancer.NetworkLoadBalancer",
			"networkloadbalancer.NetworkLoadBalancerCollection",
			"networkloadbalancer.NetworkLoadBalancerSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerBackend",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.BackendSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.BackendStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.CreateBackendDetails",
			"networkloadbalancer.UpdateBackendDetails",
			"networkloadbalancer.BackendDetails",
			"networkloadbalancer.Backend",
			"networkloadbalancer.BackendCollection",
			"networkloadbalancer.BackendSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerBackendHealth",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.BackendHealthSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.BackendHealthStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.BackendHealth",
		},
	},
	{
		Name:       "NetworkLoadBalancerBackendSet",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.BackendSetSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.BackendSetStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.CreateBackendSetDetails",
			"networkloadbalancer.UpdateBackendSetDetails",
			"networkloadbalancer.BackendSetDetails",
			"networkloadbalancer.BackendSet",
			"networkloadbalancer.BackendSetCollection",
			"networkloadbalancer.BackendSetSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerBackendSetHealth",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.BackendSetHealthSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.BackendSetHealthStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.BackendSetHealth",
		},
	},
	{
		Name:       "NetworkLoadBalancerHealthChecker",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.HealthCheckerSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.HealthCheckerStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.UpdateHealthCheckerDetails",
			"networkloadbalancer.HealthCheckerDetails",
			"networkloadbalancer.HealthChecker",
		},
	},
	{
		Name:       "NetworkLoadBalancerListener",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.ListenerSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.ListenerStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.CreateListenerDetails",
			"networkloadbalancer.UpdateListenerDetails",
			"networkloadbalancer.ListenerDetails",
			"networkloadbalancer.Listener",
			"networkloadbalancer.ListenerCollection",
			"networkloadbalancer.ListenerSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerNetworkLoadBalancerHealth",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancerHealthSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancerHealthStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.NetworkLoadBalancerHealth",
			"networkloadbalancer.NetworkLoadBalancerHealthCollection",
			"networkloadbalancer.NetworkLoadBalancerHealthSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerNetworkLoadBalancersPolicy",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancersPolicySpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancersPolicyStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.NetworkLoadBalancersPolicyCollection",
		},
	},
	{
		Name:       "NetworkLoadBalancerNetworkLoadBalancersProtocol",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancersProtocolSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.NetworkLoadBalancersProtocolStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.NetworkLoadBalancersProtocolCollection",
		},
	},
	{
		Name:       "NetworkLoadBalancerNetworkSecurityGroup",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.NetworkSecurityGroupSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.NetworkSecurityGroupStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.UpdateNetworkSecurityGroupsDetails",
		},
	},
	{
		Name:       "NetworkLoadBalancerWorkRequest",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.WorkRequest",
			"networkloadbalancer.WorkRequestCollection",
			"networkloadbalancer.WorkRequestSummary",
		},
	},
	{
		Name:       "NetworkLoadBalancerWorkRequestError",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.WorkRequestError",
			"networkloadbalancer.WorkRequestErrorCollection",
		},
	},
	{
		Name:       "NetworkLoadBalancerWorkRequestLog",
		SpecType:   reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(networkloadbalancerv1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"networkloadbalancer.WorkRequestLogEntry",
			"networkloadbalancer.WorkRequestLogEntryCollection",
		},
	},
	{
		Name:       "ArtifactsContainerConfiguration",
		SpecType:   reflect.TypeOf(artifactsv1beta1.ContainerConfigurationSpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.ContainerConfigurationStatus{}),
		SDKStructs: []string{
			"artifacts.UpdateContainerConfigurationDetails",
			"artifacts.ContainerConfiguration",
		},
	},
	{
		Name:       "ArtifactsContainerImage",
		SpecType:   reflect.TypeOf(artifactsv1beta1.ContainerImageSpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.ContainerImageStatus{}),
		SDKStructs: []string{
			"artifacts.UpdateContainerImageDetails",
			"artifacts.ContainerImage",
			"artifacts.ContainerImageCollection",
			"artifacts.ContainerImageSummary",
		},
	},
	{
		Name:       "ArtifactsContainerImageSignature",
		SpecType:   reflect.TypeOf(artifactsv1beta1.ContainerImageSignatureSpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.ContainerImageSignatureStatus{}),
		SDKStructs: []string{
			"artifacts.CreateContainerImageSignatureDetails",
			"artifacts.UpdateContainerImageSignatureDetails",
			"artifacts.ContainerImageSignature",
			"artifacts.ContainerImageSignatureCollection",
			"artifacts.ContainerImageSignatureSummary",
		},
	},
	{
		Name:       "ArtifactsContainerRepository",
		SpecType:   reflect.TypeOf(artifactsv1beta1.ContainerRepositorySpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.ContainerRepositoryStatus{}),
		SDKStructs: []string{
			"artifacts.CreateContainerRepositoryDetails",
			"artifacts.UpdateContainerRepositoryDetails",
			"artifacts.ContainerRepository",
			"artifacts.ContainerRepositoryCollection",
			"artifacts.ContainerRepositorySummary",
		},
	},
	{
		Name:       "ArtifactsGenericArtifact",
		SpecType:   reflect.TypeOf(artifactsv1beta1.GenericArtifactSpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.GenericArtifactStatus{}),
		SDKStructs: []string{
			"artifacts.UpdateGenericArtifactDetails",
			"artifacts.GenericArtifact",
			"artifacts.GenericArtifactCollection",
			"artifacts.GenericArtifactSummary",
		},
	},
	{
		Name:       "ArtifactsGenericArtifactByPath",
		SpecType:   reflect.TypeOf(artifactsv1beta1.GenericArtifactByPathSpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.GenericArtifactByPathStatus{}),
		SDKStructs: []string{
			"artifacts.UpdateGenericArtifactByPathDetails",
		},
	},
	{
		Name:       "ArtifactsRepository",
		SpecType:   reflect.TypeOf(artifactsv1beta1.RepositorySpec{}),
		StatusType: reflect.TypeOf(artifactsv1beta1.RepositoryStatus{}),
		SDKStructs: []string{
			"artifacts.ContainerRepository",
			"artifacts.GenericRepository",
			"artifacts.RepositoryCollection",
		},
	},
	{
		Name:       "CertificatesCaBundle",
		SpecType:   reflect.TypeOf(certificatesv1beta1.CaBundleSpec{}),
		StatusType: reflect.TypeOf(certificatesv1beta1.CaBundleStatus{}),
		SDKStructs: []string{
			"certificates.CaBundle",
		},
	},
	{
		Name:       "CertificatesCertificateAuthorityBundle",
		SpecType:   reflect.TypeOf(certificatesv1beta1.CertificateAuthorityBundleSpec{}),
		StatusType: reflect.TypeOf(certificatesv1beta1.CertificateAuthorityBundleStatus{}),
		SDKStructs: []string{
			"certificates.CertificateAuthorityBundle",
			"certificates.CertificateAuthorityBundleVersionSummary",
		},
	},
	{
		Name:       "CertificatesCertificateAuthorityBundleVersion",
		SpecType:   reflect.TypeOf(certificatesv1beta1.CertificateAuthorityBundleVersionSpec{}),
		StatusType: reflect.TypeOf(certificatesv1beta1.CertificateAuthorityBundleVersionStatus{}),
		SDKStructs: []string{
			"certificates.CertificateAuthorityBundleVersionCollection",
			"certificates.CertificateAuthorityBundleVersionSummary",
		},
	},
	{
		Name:       "CertificatesCertificateBundle",
		SpecType:   reflect.TypeOf(certificatesv1beta1.CertificateBundleSpec{}),
		StatusType: reflect.TypeOf(certificatesv1beta1.CertificateBundleStatus{}),
		SDKStructs: []string{
			"certificates.CertificateBundlePublicOnly",
			"certificates.CertificateBundleVersionSummary",
		},
	},
	{
		Name:       "CertificatesCertificateBundleVersion",
		SpecType:   reflect.TypeOf(certificatesv1beta1.CertificateBundleVersionSpec{}),
		StatusType: reflect.TypeOf(certificatesv1beta1.CertificateBundleVersionStatus{}),
		SDKStructs: []string{
			"certificates.CertificateBundleVersionCollection",
			"certificates.CertificateBundleVersionSummary",
		},
	},
	{
		Name:       "CertificatesManagementAssociation",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.AssociationSpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.AssociationStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.Association",
			"certificatesmanagement.AssociationCollection",
			"certificatesmanagement.AssociationSummary",
		},
	},
	{
		Name:       "CertificatesManagementCaBundle",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.CaBundleSpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.CaBundleStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.CreateCaBundleDetails",
			"certificatesmanagement.UpdateCaBundleDetails",
			"certificatesmanagement.CaBundle",
			"certificatesmanagement.CaBundleCollection",
			"certificatesmanagement.CaBundleSummary",
		},
	},
	{
		Name:       "CertificatesManagementCertificate",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.CertificateSpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.CertificateStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.CreateCertificateDetails",
			"certificatesmanagement.UpdateCertificateDetails",
			"certificatesmanagement.Certificate",
			"certificatesmanagement.CertificateCollection",
			"certificatesmanagement.CertificateVersionSummary",
			"certificatesmanagement.CertificateSummary",
		},
	},
	{
		Name:       "CertificatesManagementCertificateAuthority",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.CertificateAuthoritySpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.CertificateAuthorityStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.CreateCertificateAuthorityDetails",
			"certificatesmanagement.UpdateCertificateAuthorityDetails",
			"certificatesmanagement.CertificateAuthority",
			"certificatesmanagement.CertificateAuthorityCollection",
			"certificatesmanagement.CertificateAuthorityVersionSummary",
			"certificatesmanagement.CertificateAuthoritySummary",
		},
	},
	{
		Name:       "CertificatesManagementCertificateAuthorityVersion",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.CertificateAuthorityVersionSpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.CertificateAuthorityVersionStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.CertificateAuthorityVersion",
			"certificatesmanagement.CertificateAuthorityVersionCollection",
			"certificatesmanagement.CertificateAuthorityVersionSummary",
		},
	},
	{
		Name:       "CertificatesManagementCertificateVersion",
		SpecType:   reflect.TypeOf(certificatesmanagementv1beta1.CertificateVersionSpec{}),
		StatusType: reflect.TypeOf(certificatesmanagementv1beta1.CertificateVersionStatus{}),
		SDKStructs: []string{
			"certificatesmanagement.CertificateVersion",
			"certificatesmanagement.CertificateVersionCollection",
			"certificatesmanagement.CertificateVersionSummary",
		},
	},
	{
		Name:       "ContainerEngineAddon",
		SpecType:   reflect.TypeOf(containerenginev1beta1.AddonSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.AddonStatus{}),
		SDKStructs: []string{
			"containerengine.UpdateAddonDetails",
			"containerengine.Addon",
			"containerengine.AddonSummary",
		},
	},
	{
		Name:       "ContainerEngineAddonOption",
		SpecType:   reflect.TypeOf(containerenginev1beta1.AddonOptionSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.AddonOptionStatus{}),
		SDKStructs: []string{
			"containerengine.AddonOptionSummary",
		},
	},
	{
		Name:       "ContainerEngineCluster",
		SpecType:   reflect.TypeOf(containerenginev1beta1.ClusterSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.ClusterStatus{}),
		SDKStructs: []string{
			"containerengine.CreateClusterDetails",
			"containerengine.UpdateClusterDetails",
			"containerengine.Cluster",
			"containerengine.ClusterSummary",
		},
	},
	{
		Name:       "ContainerEngineClusterEndpointConfig",
		SpecType:   reflect.TypeOf(containerenginev1beta1.ClusterEndpointConfigSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.ClusterEndpointConfigStatus{}),
		SDKStructs: []string{
			"containerengine.CreateClusterEndpointConfigDetails",
			"containerengine.UpdateClusterEndpointConfigDetails",
			"containerengine.ClusterEndpointConfig",
		},
	},
	{
		Name:       "ContainerEngineClusterMigrateToNativeVcnStatus",
		SpecType:   reflect.TypeOf(containerenginev1beta1.ClusterMigrateToNativeVcnStatusSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.ClusterMigrateToNativeVcnStatusObservedState{}),
		SDKStructs: []string{
			"containerengine.ClusterMigrateToNativeVcnStatus",
		},
	},
	{
		Name:       "ContainerEngineClusterOption",
		SpecType:   reflect.TypeOf(containerenginev1beta1.ClusterOptionSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.ClusterOptionStatus{}),
		SDKStructs: []string{
			"containerengine.ClusterOptions",
		},
	},
	{
		Name:       "ContainerEngineCredentialRotationStatus",
		SpecType:   reflect.TypeOf(containerenginev1beta1.CredentialRotationStatusSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.CredentialRotationStatusObservedState{}),
		SDKStructs: []string{
			"containerengine.CredentialRotationStatus",
		},
	},
	{
		Name:       "ContainerEngineKubeconfig",
		SpecType:   reflect.TypeOf(containerenginev1beta1.KubeconfigSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.KubeconfigStatus{}),
		SDKStructs: []string{
			"containerengine.CreateClusterKubeconfigContentDetails",
		},
	},
	{
		Name:       "ContainerEngineNode",
		SpecType:   reflect.TypeOf(containerenginev1beta1.NodeSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.NodeStatus{}),
		SDKStructs: []string{
			"containerengine.Node",
		},
	},
	{
		Name:       "ContainerEngineNodePool",
		SpecType:   reflect.TypeOf(containerenginev1beta1.NodePoolSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.NodePoolStatus{}),
		SDKStructs: []string{
			"containerengine.CreateNodePoolDetails",
			"containerengine.UpdateNodePoolDetails",
			"containerengine.NodePool",
			"containerengine.NodePoolSummary",
		},
	},
	{
		Name:       "ContainerEngineNodePoolOption",
		SpecType:   reflect.TypeOf(containerenginev1beta1.NodePoolOptionSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.NodePoolOptionStatus{}),
		SDKStructs: []string{
			"containerengine.NodePoolOptions",
		},
	},
	{
		Name:       "ContainerEnginePodShape",
		SpecType:   reflect.TypeOf(containerenginev1beta1.PodShapeSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.PodShapeStatus{}),
		SDKStructs: []string{
			"containerengine.PodShape",
			"containerengine.PodShapeSummary",
		},
	},
	{
		Name:       "ContainerEngineVirtualNode",
		SpecType:   reflect.TypeOf(containerenginev1beta1.VirtualNodeSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.VirtualNodeStatus{}),
		SDKStructs: []string{
			"containerengine.VirtualNode",
			"containerengine.VirtualNodeSummary",
		},
	},
	{
		Name:       "ContainerEngineVirtualNodePool",
		SpecType:   reflect.TypeOf(containerenginev1beta1.VirtualNodePoolSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.VirtualNodePoolStatus{}),
		SDKStructs: []string{
			"containerengine.CreateVirtualNodePoolDetails",
			"containerengine.UpdateVirtualNodePoolDetails",
			"containerengine.VirtualNodePool",
			"containerengine.VirtualNodePoolSummary",
		},
	},
	{
		Name:       "ContainerEngineWorkRequest",
		SpecType:   reflect.TypeOf(containerenginev1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"containerengine.WorkRequest",
			"containerengine.WorkRequestSummary",
		},
	},
	{
		Name:       "ContainerEngineWorkRequestError",
		SpecType:   reflect.TypeOf(containerenginev1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"containerengine.WorkRequestError",
		},
	},
	{
		Name:       "ContainerEngineWorkRequestLog",
		SpecType:   reflect.TypeOf(containerenginev1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"containerengine.WorkRequestLogEntry",
		},
	},
	{
		Name:       "ContainerEngineWorkloadMapping",
		SpecType:   reflect.TypeOf(containerenginev1beta1.WorkloadMappingSpec{}),
		StatusType: reflect.TypeOf(containerenginev1beta1.WorkloadMappingStatus{}),
		SDKStructs: []string{
			"containerengine.CreateWorkloadMappingDetails",
			"containerengine.UpdateWorkloadMappingDetails",
			"containerengine.WorkloadMapping",
			"containerengine.WorkloadMappingSummary",
		},
	},
	{
		Name:       "IdentityAllowedDomainLicenseType",
		SpecType:   reflect.TypeOf(identityv1beta1.AllowedDomainLicenseTypeSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.AllowedDomainLicenseTypeStatus{}),
		SDKStructs: []string{
			"identity.AllowedDomainLicenseTypeSummary",
		},
	},
	{
		Name:       "IdentityApiKey",
		SpecType:   reflect.TypeOf(identityv1beta1.ApiKeySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.ApiKeyStatus{}),
		SDKStructs: []string{
			"identity.CreateApiKeyDetails",
			"identity.ApiKey",
		},
	},
	{
		Name:       "IdentityAuthToken",
		SpecType:   reflect.TypeOf(identityv1beta1.AuthTokenSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.AuthTokenStatus{}),
		SDKStructs: []string{
			"identity.CreateAuthTokenDetails",
			"identity.UpdateAuthTokenDetails",
			"identity.AuthToken",
		},
	},
	{
		Name:       "IdentityAuthenticationPolicy",
		SpecType:   reflect.TypeOf(identityv1beta1.AuthenticationPolicySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.AuthenticationPolicyStatus{}),
		SDKStructs: []string{
			"identity.UpdateAuthenticationPolicyDetails",
			"identity.AuthenticationPolicy",
		},
	},
	{
		Name:       "IdentityAvailabilityDomain",
		SpecType:   reflect.TypeOf(identityv1beta1.AvailabilityDomainSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.AvailabilityDomainStatus{}),
		SDKStructs: []string{
			"identity.AvailabilityDomain",
		},
	},
	{
		Name:       "IdentityBulkActionResourceType",
		SpecType:   reflect.TypeOf(identityv1beta1.BulkActionResourceTypeSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.BulkActionResourceTypeStatus{}),
		SDKStructs: []string{
			"identity.BulkActionResourceType",
			"identity.BulkActionResourceTypeCollection",
		},
	},
	{
		Name:       "IdentityBulkEditTagsResourceType",
		SpecType:   reflect.TypeOf(identityv1beta1.BulkEditTagsResourceTypeSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.BulkEditTagsResourceTypeStatus{}),
		SDKStructs: []string{
			"identity.BulkEditTagsResourceType",
			"identity.BulkEditTagsResourceTypeCollection",
		},
	},
	{
		Name:       "IdentityCompartment",
		SpecType:   reflect.TypeOf(identityv1beta1.CompartmentSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.CompartmentStatus{}),
		SDKStructs: []string{
			"identity.CreateCompartmentDetails",
			"identity.UpdateCompartmentDetails",
			"identity.Compartment",
		},
	},
	{
		Name:       "IdentityCostTrackingTag",
		SpecType:   reflect.TypeOf(identityv1beta1.CostTrackingTagSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.CostTrackingTagStatus{}),
		SDKStructs: []string{
			"identity.Tag",
		},
	},
	{
		Name:       "IdentityCustomerSecretKey",
		SpecType:   reflect.TypeOf(identityv1beta1.CustomerSecretKeySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.CustomerSecretKeyStatus{}),
		SDKStructs: []string{
			"identity.CreateCustomerSecretKeyDetails",
			"identity.UpdateCustomerSecretKeyDetails",
			"identity.CustomerSecretKey",
			"identity.CustomerSecretKeySummary",
		},
	},
	{
		Name:       "IdentityDbCredential",
		SpecType:   reflect.TypeOf(identityv1beta1.DbCredentialSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.DbCredentialStatus{}),
		SDKStructs: []string{
			"identity.CreateDbCredentialDetails",
			"identity.DbCredential",
			"identity.DbCredentialSummary",
		},
	},
	{
		Name:       "IdentityDomain",
		SpecType:   reflect.TypeOf(identityv1beta1.DomainSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.DomainStatus{}),
		SDKStructs: []string{
			"identity.CreateDomainDetails",
			"identity.UpdateDomainDetails",
			"identity.Domain",
			"identity.DomainSummary",
		},
	},
	{
		Name:       "IdentityDynamicGroup",
		SpecType:   reflect.TypeOf(identityv1beta1.DynamicGroupSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.DynamicGroupStatus{}),
		SDKStructs: []string{
			"identity.CreateDynamicGroupDetails",
			"identity.UpdateDynamicGroupDetails",
			"identity.DynamicGroup",
		},
	},
	{
		Name:       "IdentityFaultDomain",
		SpecType:   reflect.TypeOf(identityv1beta1.FaultDomainSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.FaultDomainStatus{}),
		SDKStructs: []string{
			"identity.FaultDomain",
		},
	},
	{
		Name:       "IdentityGroup",
		SpecType:   reflect.TypeOf(identityv1beta1.GroupSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.GroupStatus{}),
		SDKStructs: []string{
			"identity.CreateGroupDetails",
			"identity.UpdateGroupDetails",
			"identity.Group",
		},
	},
	{
		Name:       "IdentityIamWorkRequest",
		SpecType:   reflect.TypeOf(identityv1beta1.IamWorkRequestSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IamWorkRequestStatus{}),
		SDKStructs: []string{
			"identity.IamWorkRequest",
			"identity.IamWorkRequestSummary",
		},
	},
	{
		Name:       "IdentityIamWorkRequestError",
		SpecType:   reflect.TypeOf(identityv1beta1.IamWorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IamWorkRequestErrorStatus{}),
		SDKStructs: []string{
			"identity.IamWorkRequestErrorSummary",
		},
	},
	{
		Name:       "IdentityIamWorkRequestLog",
		SpecType:   reflect.TypeOf(identityv1beta1.IamWorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IamWorkRequestLogStatus{}),
		SDKStructs: []string{
			"identity.IamWorkRequestLogSummary",
		},
	},
	{
		Name:       "IdentityIdentityProvider",
		SpecType:   reflect.TypeOf(identityv1beta1.IdentityProviderSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IdentityProviderStatus{}),
		SDKStructs: []string{
			"identity.Saml2IdentityProvider",
		},
	},
	{
		Name:       "IdentityIdentityProviderGroup",
		SpecType:   reflect.TypeOf(identityv1beta1.IdentityProviderGroupSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IdentityProviderGroupStatus{}),
		SDKStructs: []string{
			"identity.IdentityProviderGroupSummary",
		},
	},
	{
		Name:       "IdentityIdpGroupMapping",
		SpecType:   reflect.TypeOf(identityv1beta1.IdpGroupMappingSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.IdpGroupMappingStatus{}),
		SDKStructs: []string{
			"identity.CreateIdpGroupMappingDetails",
			"identity.UpdateIdpGroupMappingDetails",
			"identity.IdpGroupMapping",
		},
	},
	{
		Name:       "IdentityMfaTotpDevice",
		SpecType:   reflect.TypeOf(identityv1beta1.MfaTotpDeviceSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.MfaTotpDeviceStatus{}),
		SDKStructs: []string{
			"identity.MfaTotpDevice",
			"identity.MfaTotpDeviceSummary",
		},
	},
	{
		Name:       "IdentityNetworkSource",
		SpecType:   reflect.TypeOf(identityv1beta1.NetworkSourceSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.NetworkSourceStatus{}),
		SDKStructs: []string{
			"identity.CreateNetworkSourceDetails",
			"identity.UpdateNetworkSourceDetails",
			"identity.NetworkSources",
		},
	},
	{
		Name:       "IdentityOAuthClientCredential",
		SpecType:   reflect.TypeOf(identityv1beta1.OAuthClientCredentialSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.OAuthClientCredentialStatus{}),
		SDKStructs: []string{
			"identity.CreateOAuth2ClientCredentialDetails",
			"identity.UpdateOAuth2ClientCredentialDetails",
			"identity.OAuth2ClientCredential",
			"identity.OAuth2ClientCredentialSummary",
		},
	},
	{
		Name:       "IdentityOrResetUIPassword",
		SpecType:   reflect.TypeOf(identityv1beta1.OrResetUIPasswordSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.OrResetUIPasswordStatus{}),
		SDKStructs: []string{
			"identity.UiPassword",
		},
	},
	{
		Name:       "IdentityPolicy",
		SpecType:   reflect.TypeOf(identityv1beta1.PolicySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.PolicyStatus{}),
		SDKStructs: []string{
			"identity.CreatePolicyDetails",
			"identity.UpdatePolicyDetails",
			"identity.Policy",
		},
	},
	{
		Name:       "IdentityRegion",
		SpecType:   reflect.TypeOf(identityv1beta1.RegionSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.RegionStatus{}),
		SDKStructs: []string{
			"identity.Region",
		},
	},
	{
		Name:       "IdentityRegionSubscription",
		SpecType:   reflect.TypeOf(identityv1beta1.RegionSubscriptionSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.RegionSubscriptionStatus{}),
		SDKStructs: []string{
			"identity.CreateRegionSubscriptionDetails",
			"identity.RegionSubscription",
		},
	},
	{
		Name:       "IdentitySmtpCredential",
		SpecType:   reflect.TypeOf(identityv1beta1.SmtpCredentialSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.SmtpCredentialStatus{}),
		SDKStructs: []string{
			"identity.CreateSmtpCredentialDetails",
			"identity.UpdateSmtpCredentialDetails",
			"identity.SmtpCredential",
			"identity.SmtpCredentialSummary",
		},
	},
	{
		Name:       "IdentityStandardTagNamespace",
		SpecType:   reflect.TypeOf(identityv1beta1.StandardTagNamespaceSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.StandardTagNamespaceStatus{}),
		SDKStructs: []string{
			"identity.StandardTagNamespaceTemplate",
			"identity.StandardTagNamespaceTemplateSummary",
		},
	},
	{
		Name:       "IdentityStandardTagTemplate",
		SpecType:   reflect.TypeOf(identityv1beta1.StandardTagTemplateSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.StandardTagTemplateStatus{}),
		SDKStructs: []string{
			"identity.StandardTagDefinitionTemplate",
		},
	},
	{
		Name:       "IdentitySwiftPassword",
		SpecType:   reflect.TypeOf(identityv1beta1.SwiftPasswordSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.SwiftPasswordStatus{}),
		SDKStructs: []string{
			"identity.CreateSwiftPasswordDetails",
			"identity.UpdateSwiftPasswordDetails",
			"identity.SwiftPassword",
		},
	},
	{
		Name:       "IdentityTag",
		SpecType:   reflect.TypeOf(identityv1beta1.TagSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TagStatus{}),
		SDKStructs: []string{
			"identity.CreateTagDetails",
			"identity.UpdateTagDetails",
			"identity.Tag",
			"identity.TagSummary",
		},
	},
	{
		Name:       "IdentityTagDefault",
		SpecType:   reflect.TypeOf(identityv1beta1.TagDefaultSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TagDefaultStatus{}),
		SDKStructs: []string{
			"identity.CreateTagDefaultDetails",
			"identity.UpdateTagDefaultDetails",
			"identity.TagDefault",
			"identity.TagDefaultSummary",
		},
	},
	{
		Name:       "IdentityTagNamespace",
		SpecType:   reflect.TypeOf(identityv1beta1.TagNamespaceSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TagNamespaceStatus{}),
		SDKStructs: []string{
			"identity.CreateTagNamespaceDetails",
			"identity.UpdateTagNamespaceDetails",
			"identity.TagNamespace",
			"identity.TagNamespaceSummary",
		},
	},
	{
		Name:       "IdentityTaggingWorkRequest",
		SpecType:   reflect.TypeOf(identityv1beta1.TaggingWorkRequestSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TaggingWorkRequestStatus{}),
		SDKStructs: []string{
			"identity.TaggingWorkRequest",
			"identity.TaggingWorkRequestSummary",
		},
	},
	{
		Name:       "IdentityTaggingWorkRequestError",
		SpecType:   reflect.TypeOf(identityv1beta1.TaggingWorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TaggingWorkRequestErrorStatus{}),
		SDKStructs: []string{
			"identity.TaggingWorkRequestErrorSummary",
		},
	},
	{
		Name:       "IdentityTaggingWorkRequestLog",
		SpecType:   reflect.TypeOf(identityv1beta1.TaggingWorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TaggingWorkRequestLogStatus{}),
		SDKStructs: []string{
			"identity.TaggingWorkRequestLogSummary",
		},
	},
	{
		Name:       "IdentityTenancy",
		SpecType:   reflect.TypeOf(identityv1beta1.TenancySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.TenancyStatus{}),
		SDKStructs: []string{
			"identity.Tenancy",
		},
	},
	{
		Name:       "IdentityUser",
		SpecType:   reflect.TypeOf(identityv1beta1.UserSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.UserStatus{}),
		SDKStructs: []string{
			"identity.CreateUserDetails",
			"identity.UpdateUserDetails",
			"identity.User",
		},
	},
	{
		Name:       "IdentityUserCapability",
		SpecType:   reflect.TypeOf(identityv1beta1.UserCapabilitySpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.UserCapabilityStatus{}),
		SDKStructs: []string{
			"identity.UserCapabilities",
		},
	},
	{
		Name:       "IdentityUserGroupMembership",
		SpecType:   reflect.TypeOf(identityv1beta1.UserGroupMembershipSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.UserGroupMembershipStatus{}),
		SDKStructs: []string{
			"identity.UserGroupMembership",
		},
	},
	{
		Name:       "IdentityUserState",
		SpecType:   reflect.TypeOf(identityv1beta1.UserStateSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.UserStateStatus{}),
		SDKStructs: []string{
			"identity.User",
		},
	},
	{
		Name:       "IdentityUserUIPasswordInformation",
		SpecType:   reflect.TypeOf(identityv1beta1.UserUIPasswordInformationSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.UserUIPasswordInformationStatus{}),
		SDKStructs: []string{
			"identity.UiPasswordInformation",
		},
	},
	{
		Name:       "IdentityWorkRequest",
		SpecType:   reflect.TypeOf(identityv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(identityv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"identity.WorkRequest",
			"identity.WorkRequestSummary",
		},
	},
	{
		Name:       "KeyManagementEkmsPrivateEndpoint",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.EkmsPrivateEndpointSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.EkmsPrivateEndpointStatus{}),
		SDKStructs: []string{
			"keymanagement.CreateEkmsPrivateEndpointDetails",
			"keymanagement.UpdateEkmsPrivateEndpointDetails",
			"keymanagement.EkmsPrivateEndpoint",
			"keymanagement.EkmsPrivateEndpointSummary",
		},
	},
	{
		Name:       "KeyManagementHsmCluster",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.HsmClusterSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.HsmClusterStatus{}),
		SDKStructs: []string{
			"keymanagement.CreateHsmClusterDetails",
			"keymanagement.UpdateHsmClusterDetails",
			"keymanagement.HsmCluster",
			"keymanagement.HsmClusterCollection",
			"keymanagement.HsmClusterSummary",
		},
	},
	{
		Name:       "KeyManagementHsmPartition",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.HsmPartitionSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.HsmPartitionStatus{}),
		SDKStructs: []string{
			"keymanagement.HsmPartition",
			"keymanagement.HsmPartitionCollection",
			"keymanagement.HsmPartitionSummary",
		},
	},
	{
		Name:       "KeyManagementKey",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.KeySpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.KeyStatus{}),
		SDKStructs: []string{
			"keymanagement.CreateKeyDetails",
			"keymanagement.UpdateKeyDetails",
			"keymanagement.Key",
			"keymanagement.KeyVersionSummary",
			"keymanagement.KeySummary",
		},
	},
	{
		Name:       "KeyManagementKeyVersion",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.KeyVersionSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.KeyVersionStatus{}),
		SDKStructs: []string{
			"keymanagement.KeyVersion",
			"keymanagement.KeyVersionSummary",
		},
	},
	{
		Name:       "KeyManagementPreCoUserCredential",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.PreCoUserCredentialSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.PreCoUserCredentialStatus{}),
		SDKStructs: []string{
			"keymanagement.PreCoUserCredentials",
		},
	},
	{
		Name:       "KeyManagementReplicationStatus",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.ReplicationStatusSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.ReplicationStatusObservedState{}),
		SDKStructs: []string{
			"keymanagement.ReplicationStatusDetails",
		},
	},
	{
		Name:       "KeyManagementVault",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.VaultSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.VaultStatus{}),
		SDKStructs: []string{
			"keymanagement.CreateVaultDetails",
			"keymanagement.UpdateVaultDetails",
			"keymanagement.Vault",
			"keymanagement.VaultSummary",
		},
	},
	{
		Name:       "KeyManagementVaultReplica",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.VaultReplicaSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.VaultReplicaStatus{}),
		SDKStructs: []string{
			"keymanagement.CreateVaultReplicaDetails",
			"keymanagement.VaultReplicaDetails",
			"keymanagement.VaultReplicaSummary",
		},
	},
	{
		Name:       "KeyManagementVaultUsage",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.VaultUsageSpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.VaultUsageStatus{}),
		SDKStructs: []string{
			"keymanagement.VaultUsage",
		},
	},
	{
		Name:       "KeyManagementWrappingKey",
		SpecType:   reflect.TypeOf(keymanagementv1beta1.WrappingKeySpec{}),
		StatusType: reflect.TypeOf(keymanagementv1beta1.WrappingKeyStatus{}),
		SDKStructs: []string{
			"keymanagement.WrappingKey",
		},
	},
	{
		Name:       "LimitsLimitDefinition",
		SpecType:   reflect.TypeOf(limitsv1beta1.LimitDefinitionSpec{}),
		StatusType: reflect.TypeOf(limitsv1beta1.LimitDefinitionStatus{}),
		SDKStructs: []string{
			"limits.LimitDefinitionSummary",
		},
	},
	{
		Name:       "LimitsLimitValue",
		SpecType:   reflect.TypeOf(limitsv1beta1.LimitValueSpec{}),
		StatusType: reflect.TypeOf(limitsv1beta1.LimitValueStatus{}),
		SDKStructs: []string{
			"limits.LimitValueSummary",
		},
	},
	{
		Name:       "LimitsQuota",
		SpecType:   reflect.TypeOf(limitsv1beta1.QuotaSpec{}),
		StatusType: reflect.TypeOf(limitsv1beta1.QuotaStatus{}),
		SDKStructs: []string{
			"limits.CreateQuotaDetails",
			"limits.UpdateQuotaDetails",
			"limits.Quota",
			"limits.QuotaSummary",
		},
	},
	{
		Name:       "LimitsResourceAvailability",
		SpecType:   reflect.TypeOf(limitsv1beta1.ResourceAvailabilitySpec{}),
		StatusType: reflect.TypeOf(limitsv1beta1.ResourceAvailabilityStatus{}),
		SDKStructs: []string{
			"limits.ResourceAvailability",
		},
	},
	{
		Name:       "LimitsService",
		SpecType:   reflect.TypeOf(limitsv1beta1.ServiceSpec{}),
		StatusType: reflect.TypeOf(limitsv1beta1.ServiceStatus{}),
		SDKStructs: []string{
			"limits.ServiceSummary",
		},
	},
	{
		Name:       "SecretsSecretBundle",
		SpecType:   reflect.TypeOf(secretsv1beta1.SecretBundleSpec{}),
		StatusType: reflect.TypeOf(secretsv1beta1.SecretBundleStatus{}),
		SDKStructs: []string{
			"secrets.SecretBundle",
			"secrets.SecretBundleVersionSummary",
		},
	},
	{
		Name:       "SecretsSecretBundleByName",
		SpecType:   reflect.TypeOf(secretsv1beta1.SecretBundleByNameSpec{}),
		StatusType: reflect.TypeOf(secretsv1beta1.SecretBundleByNameStatus{}),
		SDKStructs: []string{
			"secrets.SecretBundle",
			"secrets.SecretBundleVersionSummary",
		},
	},
	{
		Name:       "SecretsSecretBundleVersion",
		SpecType:   reflect.TypeOf(secretsv1beta1.SecretBundleVersionSpec{}),
		StatusType: reflect.TypeOf(secretsv1beta1.SecretBundleVersionStatus{}),
		SDKStructs: []string{
			"secrets.SecretBundleVersionSummary",
		},
	},
	{
		Name:       "VaultSecret",
		SpecType:   reflect.TypeOf(vaultv1beta1.SecretSpec{}),
		StatusType: reflect.TypeOf(vaultv1beta1.SecretStatus{}),
		SDKStructs: []string{
			"vault.CreateSecretDetails",
			"vault.UpdateSecretDetails",
			"vault.Secret",
			"vault.SecretVersionSummary",
			"vault.SecretSummary",
		},
	},
	{
		Name:       "VaultSecretVersion",
		SpecType:   reflect.TypeOf(vaultv1beta1.SecretVersionSpec{}),
		StatusType: reflect.TypeOf(vaultv1beta1.SecretVersionStatus{}),
		SDKStructs: []string{
			"vault.SecretVersion",
			"vault.SecretVersionSummary",
		},
	},
	{
		Name:       "CoreAllDrgAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.AllDrgAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AllDrgAttachmentStatus{}),
		SDKStructs: []string{
			"core.DrgAttachmentInfo",
		},
	},
	{
		Name:       "CoreAllowedIkeIPSecParameter",
		SpecType:   reflect.TypeOf(corev1beta1.AllowedIkeIPSecParameterSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AllowedIkeIPSecParameterStatus{}),
		SDKStructs: []string{
			"core.AllowedIkeIpSecParameters",
		},
	},
	{
		Name:       "CoreAllowedPeerRegionsForRemotePeering",
		SpecType:   reflect.TypeOf(corev1beta1.AllowedPeerRegionsForRemotePeeringSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AllowedPeerRegionsForRemotePeeringStatus{}),
		SDKStructs: []string{
			"core.PeerRegionForRemotePeering",
		},
	},
	{
		Name:       "CoreAppCatalogListing",
		SpecType:   reflect.TypeOf(corev1beta1.AppCatalogListingSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AppCatalogListingStatus{}),
		SDKStructs: []string{
			"core.AppCatalogListing",
			"core.AppCatalogListingSummary",
		},
	},
	{
		Name:       "CoreAppCatalogListingAgreement",
		SpecType:   reflect.TypeOf(corev1beta1.AppCatalogListingAgreementSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AppCatalogListingAgreementStatus{}),
		SDKStructs: []string{
			"core.AppCatalogListingResourceVersionAgreements",
		},
	},
	{
		Name:       "CoreAppCatalogListingResourceVersion",
		SpecType:   reflect.TypeOf(corev1beta1.AppCatalogListingResourceVersionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AppCatalogListingResourceVersionStatus{}),
		SDKStructs: []string{
			"core.AppCatalogListingResourceVersion",
			"core.AppCatalogListingResourceVersionSummary",
		},
	},
	{
		Name:       "CoreAppCatalogSubscription",
		SpecType:   reflect.TypeOf(corev1beta1.AppCatalogSubscriptionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.AppCatalogSubscriptionStatus{}),
		SDKStructs: []string{
			"core.CreateAppCatalogSubscriptionDetails",
			"core.AppCatalogSubscription",
			"core.AppCatalogSubscriptionSummary",
		},
	},
	{
		Name:       "CoreBlockVolumeReplica",
		SpecType:   reflect.TypeOf(corev1beta1.BlockVolumeReplicaSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BlockVolumeReplicaStatus{}),
		SDKStructs: []string{
			"core.BlockVolumeReplicaDetails",
			"core.BlockVolumeReplica",
		},
	},
	{
		Name:       "CoreBootVolume",
		SpecType:   reflect.TypeOf(corev1beta1.BootVolumeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BootVolumeStatus{}),
		SDKStructs: []string{
			"core.CreateBootVolumeDetails",
			"core.UpdateBootVolumeDetails",
			"core.BootVolume",
		},
	},
	{
		Name:       "CoreBootVolumeAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.BootVolumeAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BootVolumeAttachmentStatus{}),
		SDKStructs: []string{
			"core.BootVolumeAttachment",
		},
	},
	{
		Name:       "CoreBootVolumeBackup",
		SpecType:   reflect.TypeOf(corev1beta1.BootVolumeBackupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BootVolumeBackupStatus{}),
		SDKStructs: []string{
			"core.CreateBootVolumeBackupDetails",
			"core.UpdateBootVolumeBackupDetails",
			"core.BootVolumeBackup",
		},
	},
	{
		Name:       "CoreBootVolumeKMSKey",
		SpecType:   reflect.TypeOf(corev1beta1.BootVolumeKmsKeySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BootVolumeKmsKeyStatus{}),
		SDKStructs: []string{
			"core.UpdateBootVolumeKmsKeyDetails",
			"core.BootVolumeKmsKey",
		},
	},
	{
		Name:       "CoreBootVolumeReplica",
		SpecType:   reflect.TypeOf(corev1beta1.BootVolumeReplicaSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.BootVolumeReplicaStatus{}),
		SDKStructs: []string{
			"core.BootVolumeReplicaDetails",
			"core.BootVolumeReplica",
		},
	},
	{
		Name:       "CoreByoipAllocatedRange",
		SpecType:   reflect.TypeOf(corev1beta1.ByoipAllocatedRangeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ByoipAllocatedRangeStatus{}),
		SDKStructs: []string{
			"core.ByoipAllocatedRangeCollection",
			"core.ByoipAllocatedRangeSummary",
		},
	},
	{
		Name:       "CoreByoipRange",
		SpecType:   reflect.TypeOf(corev1beta1.ByoipRangeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ByoipRangeStatus{}),
		SDKStructs: []string{
			"core.CreateByoipRangeDetails",
			"core.UpdateByoipRangeDetails",
			"core.ByoipRange",
			"core.ByoipRangeCollection",
			"core.ByoipRangeSummary",
		},
	},
	{
		Name:       "CoreCPE",
		SpecType:   reflect.TypeOf(corev1beta1.CpeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CpeStatus{}),
		SDKStructs: []string{
			"core.CreateCpeDetails",
			"core.UpdateCpeDetails",
			"core.Cpe",
		},
	},
	{
		Name:       "CoreCaptureFilter",
		SpecType:   reflect.TypeOf(corev1beta1.CaptureFilterSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CaptureFilterStatus{}),
		SDKStructs: []string{
			"core.CreateCaptureFilterDetails",
			"core.UpdateCaptureFilterDetails",
			"core.CaptureFilter",
		},
	},
	{
		Name:       "CoreClusterNetwork",
		SpecType:   reflect.TypeOf(corev1beta1.ClusterNetworkSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ClusterNetworkStatus{}),
		SDKStructs: []string{
			"core.CreateClusterNetworkDetails",
			"core.UpdateClusterNetworkDetails",
			"core.ClusterNetwork",
			"core.ClusterNetworkSummary",
		},
	},
	{
		Name:       "CoreClusterNetworkInstance",
		SpecType:   reflect.TypeOf(corev1beta1.ClusterNetworkInstanceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ClusterNetworkInstanceStatus{}),
		SDKStructs: []string{
			"core.InstanceSummary",
		},
	},
	{
		Name:       "CoreComputeCapacityReport",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityReportSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityReportStatus{}),
		SDKStructs: []string{
			"core.CreateComputeCapacityReportDetails",
			"core.ComputeCapacityReport",
		},
	},
	{
		Name:       "CoreComputeCapacityReservation",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityReservationSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityReservationStatus{}),
		SDKStructs: []string{
			"core.CreateComputeCapacityReservationDetails",
			"core.UpdateComputeCapacityReservationDetails",
			"core.ComputeCapacityReservation",
			"core.ComputeCapacityReservationSummary",
		},
	},
	{
		Name:       "CoreComputeCapacityReservationInstance",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityReservationInstanceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityReservationInstanceStatus{}),
		SDKStructs: []string{
			"core.CapacityReservationInstanceSummary",
		},
	},
	{
		Name:       "CoreComputeCapacityReservationInstanceShape",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityReservationInstanceShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityReservationInstanceShapeStatus{}),
		SDKStructs: []string{
			"core.ComputeCapacityReservationInstanceShapeSummary",
		},
	},
	{
		Name:       "CoreComputeCapacityTopology",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityTopologySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityTopologyStatus{}),
		SDKStructs: []string{
			"core.CreateComputeCapacityTopologyDetails",
			"core.UpdateComputeCapacityTopologyDetails",
			"core.ComputeCapacityTopology",
			"core.ComputeCapacityTopologyCollection",
			"core.ComputeCapacityTopologySummary",
		},
	},
	{
		Name:       "CoreComputeCapacityTopologyComputeBareMetalHost",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeBareMetalHostSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeBareMetalHostStatus{}),
		SDKStructs: []string{
			"core.ComputeBareMetalHostCollection",
		},
	},
	{
		Name:       "CoreComputeCapacityTopologyComputeHpcIsland",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeHpcIslandSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeHpcIslandStatus{}),
		SDKStructs: []string{
			"core.ComputeHpcIslandCollection",
		},
	},
	{
		Name:       "CoreComputeCapacityTopologyComputeNetworkBlock",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeNetworkBlockSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeCapacityTopologyComputeNetworkBlockStatus{}),
		SDKStructs: []string{
			"core.ComputeNetworkBlockCollection",
		},
	},
	{
		Name:       "CoreComputeCluster",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeClusterSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeClusterStatus{}),
		SDKStructs: []string{
			"core.CreateComputeClusterDetails",
			"core.UpdateComputeClusterDetails",
			"core.ComputeCluster",
			"core.ComputeClusterCollection",
			"core.ComputeClusterSummary",
		},
	},
	{
		Name:       "CoreComputeGlobalImageCapabilitySchema",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeGlobalImageCapabilitySchemaSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeGlobalImageCapabilitySchemaStatus{}),
		SDKStructs: []string{
			"core.ComputeGlobalImageCapabilitySchema",
			"core.ComputeGlobalImageCapabilitySchemaVersionSummary",
			"core.ComputeGlobalImageCapabilitySchemaSummary",
		},
	},
	{
		Name:       "CoreComputeGlobalImageCapabilitySchemaVersion",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeGlobalImageCapabilitySchemaVersionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeGlobalImageCapabilitySchemaVersionStatus{}),
		SDKStructs: []string{
			"core.ComputeGlobalImageCapabilitySchemaVersion",
			"core.ComputeGlobalImageCapabilitySchemaVersionSummary",
		},
	},
	{
		Name:       "CoreComputeImageCapabilitySchema",
		SpecType:   reflect.TypeOf(corev1beta1.ComputeImageCapabilitySchemaSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ComputeImageCapabilitySchemaStatus{}),
		SDKStructs: []string{
			"core.CreateComputeImageCapabilitySchemaDetails",
			"core.UpdateComputeImageCapabilitySchemaDetails",
			"core.ComputeImageCapabilitySchema",
			"core.ComputeImageCapabilitySchemaSummary",
		},
	},
	{
		Name:       "CoreConsoleHistory",
		SpecType:   reflect.TypeOf(corev1beta1.ConsoleHistorySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ConsoleHistoryStatus{}),
		SDKStructs: []string{
			"core.UpdateConsoleHistoryDetails",
			"core.ConsoleHistory",
		},
	},
	{
		Name:       "CoreConsoleHistoryContent",
		SpecType:   reflect.TypeOf(corev1beta1.ConsoleHistoryContentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ConsoleHistoryContentStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "CoreCpeDeviceConfigContent",
		SpecType:   reflect.TypeOf(corev1beta1.CpeDeviceConfigContentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CpeDeviceConfigContentStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "CoreCpeDeviceShape",
		SpecType:   reflect.TypeOf(corev1beta1.CpeDeviceShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CpeDeviceShapeStatus{}),
		SDKStructs: []string{
			"core.CpeDeviceShapeSummary",
		},
	},
	{
		Name:       "CoreCrossConnect",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectObservedState{}),
		SDKStructs: []string{
			"core.CreateCrossConnectDetails",
			"core.UpdateCrossConnectDetails",
			"core.CrossConnect",
		},
	},
	{
		Name:       "CoreCrossConnectGroup",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectGroupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectGroupStatus{}),
		SDKStructs: []string{
			"core.CreateCrossConnectGroupDetails",
			"core.UpdateCrossConnectGroupDetails",
			"core.CrossConnectGroup",
		},
	},
	{
		Name:       "CoreCrossConnectLetterOfAuthority",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectLetterOfAuthoritySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectLetterOfAuthorityStatus{}),
		SDKStructs: []string{
			"core.LetterOfAuthority",
		},
	},
	{
		Name:       "CoreCrossConnectLocation",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectLocationSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectLocationStatus{}),
		SDKStructs: []string{
			"core.CrossConnectLocation",
		},
	},
	{
		Name:       "CoreCrossConnectMapping",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectMappingSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectMappingStatus{}),
		SDKStructs: []string{
			"core.CrossConnectMappingDetails",
			"core.CrossConnectMapping",
		},
	},
	{
		Name:       "CoreCrossConnectStatus",
		SpecType:   reflect.TypeOf(corev1beta1.CrossConnectStatusSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossConnectStatusObservedState{}),
		SDKStructs: []string{
			"core.CrossConnectStatus",
		},
	},
	{
		Name:       "CoreCrossconnectPortSpeedShape",
		SpecType:   reflect.TypeOf(corev1beta1.CrossconnectPortSpeedShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.CrossconnectPortSpeedShapeStatus{}),
		SDKStructs: []string{
			"core.CrossConnectPortSpeedShape",
		},
	},
	{
		Name:       "CoreDRG",
		SpecType:   reflect.TypeOf(corev1beta1.DrgSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgStatus{}),
		SDKStructs: []string{
			"core.CreateDrgDetails",
			"core.UpdateDrgDetails",
			"core.Drg",
		},
	},
	{
		Name:       "CoreDRGAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.DrgAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgAttachmentStatus{}),
		SDKStructs: []string{
			"core.CreateDrgAttachmentDetails",
			"core.UpdateDrgAttachmentDetails",
			"core.DrgAttachment",
		},
	},
	{
		Name:       "CoreDRGRouteDistribution",
		SpecType:   reflect.TypeOf(corev1beta1.DrgRouteDistributionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgRouteDistributionStatus{}),
		SDKStructs: []string{
			"core.CreateDrgRouteDistributionDetails",
			"core.UpdateDrgRouteDistributionDetails",
			"core.DrgRouteDistribution",
		},
	},
	{
		Name:       "CoreDRGRouteDistributionStatement",
		SpecType:   reflect.TypeOf(corev1beta1.DrgRouteDistributionStatementSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgRouteDistributionStatementStatus{}),
		SDKStructs: []string{
			"core.UpdateDrgRouteDistributionStatementDetails",
			"core.DrgRouteDistributionStatement",
		},
	},
	{
		Name:       "CoreDRGRouteRule",
		SpecType:   reflect.TypeOf(corev1beta1.DrgRouteRuleSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgRouteRuleStatus{}),
		SDKStructs: []string{
			"core.UpdateDrgRouteRuleDetails",
			"core.DrgRouteRule",
		},
	},
	{
		Name:       "CoreDRGRouteTable",
		SpecType:   reflect.TypeOf(corev1beta1.DrgRouteTableSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgRouteTableStatus{}),
		SDKStructs: []string{
			"core.CreateDrgRouteTableDetails",
			"core.UpdateDrgRouteTableDetails",
			"core.DrgRouteTable",
		},
	},
	{
		Name:       "CoreDedicatedVmHost",
		SpecType:   reflect.TypeOf(corev1beta1.DedicatedVmHostSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DedicatedVmHostStatus{}),
		SDKStructs: []string{
			"core.CreateDedicatedVmHostDetails",
			"core.UpdateDedicatedVmHostDetails",
			"core.DedicatedVmHost",
			"core.DedicatedVmHostSummary",
		},
	},
	{
		Name:       "CoreDedicatedVmHostInstance",
		SpecType:   reflect.TypeOf(corev1beta1.DedicatedVmHostInstanceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DedicatedVmHostInstanceStatus{}),
		SDKStructs: []string{
			"core.DedicatedVmHostInstanceSummary",
		},
	},
	{
		Name:       "CoreDedicatedVmHostInstanceShape",
		SpecType:   reflect.TypeOf(corev1beta1.DedicatedVmHostInstanceShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DedicatedVmHostInstanceShapeStatus{}),
		SDKStructs: []string{
			"core.DedicatedVmHostInstanceShapeSummary",
		},
	},
	{
		Name:       "CoreDedicatedVmHostShape",
		SpecType:   reflect.TypeOf(corev1beta1.DedicatedVmHostShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DedicatedVmHostShapeStatus{}),
		SDKStructs: []string{
			"core.DedicatedVmHostShapeSummary",
		},
	},
	{
		Name:       "CoreDhcpOption",
		SpecType:   reflect.TypeOf(corev1beta1.DhcpOptionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DhcpOptionStatus{}),
		SDKStructs: []string{
			"core.CreateDhcpDetails",
			"core.UpdateDhcpDetails",
			"core.DhcpOptions",
		},
	},
	{
		Name:       "CoreDrgRedundancyStatus",
		SpecType:   reflect.TypeOf(corev1beta1.DrgRedundancyStatusSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.DrgRedundancyStatusObservedState{}),
		SDKStructs: []string{
			"core.DrgRedundancyStatus",
		},
	},
	{
		Name:       "CoreFastConnectProviderService",
		SpecType:   reflect.TypeOf(corev1beta1.FastConnectProviderServiceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.FastConnectProviderServiceStatus{}),
		SDKStructs: []string{
			"core.FastConnectProviderService",
		},
	},
	{
		Name:       "CoreFastConnectProviderServiceKey",
		SpecType:   reflect.TypeOf(corev1beta1.FastConnectProviderServiceKeySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.FastConnectProviderServiceKeyStatus{}),
		SDKStructs: []string{
			"core.FastConnectProviderServiceKey",
		},
	},
	{
		Name:       "CoreFastConnectProviderVirtualCircuitBandwidthShape",
		SpecType:   reflect.TypeOf(corev1beta1.FastConnectProviderVirtualCircuitBandwidthShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.FastConnectProviderVirtualCircuitBandwidthShapeStatus{}),
		SDKStructs: []string{
			"core.VirtualCircuitBandwidthShape",
		},
	},
	{
		Name:       "CoreIPSecConnection",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionStatus{}),
		SDKStructs: []string{
			"core.CreateIpSecConnectionDetails",
			"core.UpdateIpSecConnectionDetails",
			"core.IpSecConnection",
		},
	},
	{
		Name:       "CoreIPSecConnectionDeviceConfig",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionDeviceConfigSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionDeviceConfigStatus{}),
		SDKStructs: []string{
			"core.IpSecConnectionDeviceConfig",
		},
	},
	{
		Name:       "CoreIPSecConnectionDeviceStatus",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionDeviceStatusSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionDeviceStatusObservedState{}),
		SDKStructs: []string{
			"core.IpSecConnectionDeviceStatus",
		},
	},
	{
		Name:       "CoreIPSecConnectionTunnel",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionTunnelSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionTunnelStatus{}),
		SDKStructs: []string{
			"core.CreateIpSecConnectionTunnelDetails",
			"core.UpdateIpSecConnectionTunnelDetails",
			"core.IpSecConnectionTunnel",
		},
	},
	{
		Name:       "CoreIPSecConnectionTunnelError",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionTunnelErrorSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionTunnelErrorStatus{}),
		SDKStructs: []string{
			"core.IpSecConnectionTunnelErrorDetails",
		},
	},
	{
		Name:       "CoreIPSecConnectionTunnelRoute",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionTunnelRouteSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionTunnelRouteStatus{}),
		SDKStructs: []string{
			"core.TunnelRouteSummary",
		},
	},
	{
		Name:       "CoreIPSecConnectionTunnelSecurityAssociation",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionTunnelSecurityAssociationSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionTunnelSecurityAssociationStatus{}),
		SDKStructs: []string{
			"core.TunnelSecurityAssociationSummary",
		},
	},
	{
		Name:       "CoreIPSecConnectionTunnelSharedSecret",
		SpecType:   reflect.TypeOf(corev1beta1.IPSecConnectionTunnelSharedSecretSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IPSecConnectionTunnelSharedSecretStatus{}),
		SDKStructs: []string{
			"core.UpdateIpSecConnectionTunnelSharedSecretDetails",
			"core.IpSecConnectionTunnelSharedSecret",
		},
	},
	{
		Name:       "CoreImage",
		SpecType:   reflect.TypeOf(corev1beta1.ImageSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ImageStatus{}),
		SDKStructs: []string{
			"core.CreateImageDetails",
			"core.UpdateImageDetails",
			"core.Image",
		},
	},
	{
		Name:       "CoreImageShapeCompatibilityEntry",
		SpecType:   reflect.TypeOf(corev1beta1.ImageShapeCompatibilityEntrySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ImageShapeCompatibilityEntryStatus{}),
		SDKStructs: []string{
			"core.ImageShapeCompatibilityEntry",
		},
	},
	{
		Name:       "CoreInstance",
		SpecType:   reflect.TypeOf(corev1beta1.InstanceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstanceStatus{}),
		SDKStructs: []string{
			"core.UpdateInstanceDetails",
			"core.Instance",
			"core.InstanceSummary",
		},
	},
	{
		Name:       "CoreInstanceConfiguration",
		SpecType:   reflect.TypeOf(corev1beta1.InstanceConfigurationSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstanceConfigurationStatus{}),
		SDKStructs: []string{
			"core.CreateInstanceConfigurationDetails",
			"core.UpdateInstanceConfigurationDetails",
			"core.InstanceConfiguration",
			"core.InstanceConfigurationSummary",
		},
	},
	{
		Name:       "CoreInstanceConsoleConnection",
		SpecType:   reflect.TypeOf(corev1beta1.InstanceConsoleConnectionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstanceConsoleConnectionStatus{}),
		SDKStructs: []string{
			"core.CreateInstanceConsoleConnectionDetails",
			"core.UpdateInstanceConsoleConnectionDetails",
			"core.InstanceConsoleConnection",
		},
	},
	{
		Name:       "CoreInstanceDevice",
		SpecType:   reflect.TypeOf(corev1beta1.InstanceDeviceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstanceDeviceStatus{}),
		SDKStructs: []string{
			"core.Device",
		},
	},
	{
		Name:       "CoreInstanceMaintenanceReboot",
		SpecType:   reflect.TypeOf(corev1beta1.InstanceMaintenanceRebootSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstanceMaintenanceRebootStatus{}),
		SDKStructs: []string{
			"core.InstanceMaintenanceReboot",
		},
	},
	{
		Name:       "CoreInstancePool",
		SpecType:   reflect.TypeOf(corev1beta1.InstancePoolSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstancePoolStatus{}),
		SDKStructs: []string{
			"core.CreateInstancePoolDetails",
			"core.UpdateInstancePoolDetails",
			"core.InstancePool",
			"core.InstancePoolSummary",
		},
	},
	{
		Name:       "CoreInstancePoolInstance",
		SpecType:   reflect.TypeOf(corev1beta1.InstancePoolInstanceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstancePoolInstanceStatus{}),
		SDKStructs: []string{
			"core.InstancePoolInstance",
		},
	},
	{
		Name:       "CoreInstancePoolLoadBalancerAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.InstancePoolLoadBalancerAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InstancePoolLoadBalancerAttachmentStatus{}),
		SDKStructs: []string{
			"core.InstancePoolLoadBalancerAttachment",
		},
	},
	{
		Name:       "CoreInternetGateway",
		SpecType:   reflect.TypeOf(corev1beta1.InternetGatewaySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.InternetGatewayStatus{}),
		SDKStructs: []string{
			"core.CreateInternetGatewayDetails",
			"core.UpdateInternetGatewayDetails",
			"core.InternetGateway",
		},
	},
	{
		Name:       "CoreIpsecCpeDeviceConfigContent",
		SpecType:   reflect.TypeOf(corev1beta1.IpsecCpeDeviceConfigContentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.IpsecCpeDeviceConfigContentStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "CoreIpv6",
		SpecType:   reflect.TypeOf(corev1beta1.Ipv6Spec{}),
		StatusType: reflect.TypeOf(corev1beta1.Ipv6Status{}),
		SDKStructs: []string{
			"core.CreateIpv6Details",
			"core.UpdateIpv6Details",
			"core.Ipv6",
		},
	},
	{
		Name:       "CoreLocalPeeringGateway",
		SpecType:   reflect.TypeOf(corev1beta1.LocalPeeringGatewaySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.LocalPeeringGatewayStatus{}),
		SDKStructs: []string{
			"core.CreateLocalPeeringGatewayDetails",
			"core.UpdateLocalPeeringGatewayDetails",
			"core.LocalPeeringGateway",
		},
	},
	{
		Name:       "CoreMeasuredBootReport",
		SpecType:   reflect.TypeOf(corev1beta1.MeasuredBootReportSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.MeasuredBootReportStatus{}),
		SDKStructs: []string{
			"core.MeasuredBootReport",
		},
	},
	{
		Name:       "CoreNATGateway",
		SpecType:   reflect.TypeOf(corev1beta1.NatGatewaySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.NatGatewayStatus{}),
		SDKStructs: []string{
			"core.CreateNatGatewayDetails",
			"core.UpdateNatGatewayDetails",
			"core.NatGateway",
		},
	},
	{
		Name:       "CoreNetworkSecurityGroup",
		SpecType:   reflect.TypeOf(corev1beta1.NetworkSecurityGroupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.NetworkSecurityGroupStatus{}),
		SDKStructs: []string{
			"core.CreateNetworkSecurityGroupDetails",
			"core.UpdateNetworkSecurityGroupDetails",
			"core.NetworkSecurityGroup",
		},
	},
	{
		Name:       "CoreNetworkSecurityGroupSecurityRule",
		SpecType:   reflect.TypeOf(corev1beta1.NetworkSecurityGroupSecurityRuleSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.NetworkSecurityGroupSecurityRuleStatus{}),
		SDKStructs: []string{
			"core.SecurityRule",
		},
	},
	{
		Name:       "CoreNetworkSecurityGroupVnic",
		SpecType:   reflect.TypeOf(corev1beta1.NetworkSecurityGroupVnicSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.NetworkSecurityGroupVnicStatus{}),
		SDKStructs: []string{
			"core.NetworkSecurityGroupVnic",
		},
	},
	{
		Name:       "CoreNetworkingTopology",
		SpecType:   reflect.TypeOf(corev1beta1.NetworkingTopologySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.NetworkingTopologyStatus{}),
		SDKStructs: []string{
			"core.NetworkingTopology",
		},
	},
	{
		Name:       "CorePrivateIP",
		SpecType:   reflect.TypeOf(corev1beta1.PrivateIpSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.PrivateIpStatus{}),
		SDKStructs: []string{
			"core.CreatePrivateIpDetails",
			"core.UpdatePrivateIpDetails",
			"core.PrivateIp",
		},
	},
	{
		Name:       "CorePublicIP",
		SpecType:   reflect.TypeOf(corev1beta1.PublicIpSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.PublicIpStatus{}),
		SDKStructs: []string{
			"core.CreatePublicIpDetails",
			"core.UpdatePublicIpDetails",
			"core.PublicIp",
		},
	},
	{
		Name:       "CorePublicIPPool",
		SpecType:   reflect.TypeOf(corev1beta1.PublicIpPoolSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.PublicIpPoolStatus{}),
		SDKStructs: []string{
			"core.CreatePublicIpPoolDetails",
			"core.UpdatePublicIpPoolDetails",
			"core.PublicIpPool",
			"core.PublicIpPoolCollection",
			"core.PublicIpPoolSummary",
		},
	},
	{
		Name:       "CorePublicIpByIpAddress",
		SpecType:   reflect.TypeOf(corev1beta1.PublicIpByIpAddressSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.PublicIpByIpAddressStatus{}),
		SDKStructs: []string{
			"core.GetPublicIpByIpAddressDetails",
		},
	},
	{
		Name:       "CorePublicIpByPrivateIpId",
		SpecType:   reflect.TypeOf(corev1beta1.PublicIpByPrivateIpIdSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.PublicIpByPrivateIpIdStatus{}),
		SDKStructs: []string{
			"core.GetPublicIpByPrivateIpIdDetails",
		},
	},
	{
		Name:       "CoreRemotePeeringConnection",
		SpecType:   reflect.TypeOf(corev1beta1.RemotePeeringConnectionSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.RemotePeeringConnectionStatus{}),
		SDKStructs: []string{
			"core.CreateRemotePeeringConnectionDetails",
			"core.UpdateRemotePeeringConnectionDetails",
			"core.RemotePeeringConnection",
		},
	},
	{
		Name:       "CoreRouteTable",
		SpecType:   reflect.TypeOf(corev1beta1.RouteTableSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.RouteTableStatus{}),
		SDKStructs: []string{
			"core.CreateRouteTableDetails",
			"core.UpdateRouteTableDetails",
			"core.RouteTable",
		},
	},
	{
		Name:       "CoreSecurityList",
		SpecType:   reflect.TypeOf(corev1beta1.SecurityListSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.SecurityListStatus{}),
		SDKStructs: []string{
			"core.CreateSecurityListDetails",
			"core.UpdateSecurityListDetails",
			"core.SecurityList",
		},
	},
	{
		Name:       "CoreService",
		SpecType:   reflect.TypeOf(corev1beta1.ServiceSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ServiceStatus{}),
		SDKStructs: []string{
			"core.Service",
		},
	},
	{
		Name:       "CoreServiceGateway",
		SpecType:   reflect.TypeOf(corev1beta1.ServiceGatewaySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ServiceGatewayStatus{}),
		SDKStructs: []string{
			"core.CreateServiceGatewayDetails",
			"core.UpdateServiceGatewayDetails",
			"core.ServiceGateway",
		},
	},
	{
		Name:       "CoreShape",
		SpecType:   reflect.TypeOf(corev1beta1.ShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.ShapeStatus{}),
		SDKStructs: []string{
			"core.Shape",
		},
	},
	{
		Name:       "CoreSubnet",
		SpecType:   reflect.TypeOf(corev1beta1.SubnetSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.SubnetStatus{}),
		SDKStructs: []string{
			"core.CreateSubnetDetails",
			"core.UpdateSubnetDetails",
			"core.Subnet",
		},
	},
	{
		Name:       "CoreSubnetTopology",
		SpecType:   reflect.TypeOf(corev1beta1.SubnetTopologySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.SubnetTopologyStatus{}),
		SDKStructs: []string{
			"core.SubnetTopology",
		},
	},
	{
		Name:       "CoreTunnelCPEDeviceConfig",
		SpecType:   reflect.TypeOf(corev1beta1.TunnelCpeDeviceConfigSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.TunnelCpeDeviceConfigStatus{}),
		SDKStructs: []string{
			"core.UpdateTunnelCpeDeviceConfigDetails",
			"core.TunnelCpeDeviceConfig",
		},
	},
	{
		Name:       "CoreTunnelCpeDeviceConfigContent",
		SpecType:   reflect.TypeOf(corev1beta1.TunnelCpeDeviceConfigContentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.TunnelCpeDeviceConfigContentStatus{}),
		SDKStructs: []string{},
	},
	{
		Name:       "CoreUpgradeStatus",
		SpecType:   reflect.TypeOf(corev1beta1.UpgradeStatusSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.UpgradeStatusObservedState{}),
		SDKStructs: []string{
			"core.UpgradeStatus",
		},
	},
	{
		Name:       "CoreVCN",
		SpecType:   reflect.TypeOf(corev1beta1.VcnSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VcnStatus{}),
		SDKStructs: []string{
			"core.CreateVcnDetails",
			"core.UpdateVcnDetails",
			"core.Vcn",
		},
	},
	{
		Name:       "CoreVLAN",
		SpecType:   reflect.TypeOf(corev1beta1.VlanSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VlanStatus{}),
		SDKStructs: []string{
			"core.CreateVlanDetails",
			"core.UpdateVlanDetails",
			"core.Vlan",
		},
	},
	{
		Name:       "CoreVNIC",
		SpecType:   reflect.TypeOf(corev1beta1.VnicSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VnicStatus{}),
		SDKStructs: []string{
			"core.CreateVnicDetails",
			"core.UpdateVnicDetails",
			"core.Vnic",
		},
	},
	{
		Name:       "CoreVcnDnsResolverAssociation",
		SpecType:   reflect.TypeOf(corev1beta1.VcnDnsResolverAssociationSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VcnDnsResolverAssociationStatus{}),
		SDKStructs: []string{
			"core.VcnDnsResolverAssociation",
		},
	},
	{
		Name:       "CoreVcnTopology",
		SpecType:   reflect.TypeOf(corev1beta1.VcnTopologySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VcnTopologyStatus{}),
		SDKStructs: []string{
			"core.VcnTopology",
		},
	},
	{
		Name:       "CoreVirtualCircuit",
		SpecType:   reflect.TypeOf(corev1beta1.VirtualCircuitSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VirtualCircuitStatus{}),
		SDKStructs: []string{
			"core.CreateVirtualCircuitDetails",
			"core.UpdateVirtualCircuitDetails",
			"core.VirtualCircuit",
		},
	},
	{
		Name:       "CoreVirtualCircuitAssociatedTunnel",
		SpecType:   reflect.TypeOf(corev1beta1.VirtualCircuitAssociatedTunnelSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VirtualCircuitAssociatedTunnelStatus{}),
		SDKStructs: []string{
			"core.VirtualCircuitAssociatedTunnelDetails",
		},
	},
	{
		Name:       "CoreVirtualCircuitBandwidthShape",
		SpecType:   reflect.TypeOf(corev1beta1.VirtualCircuitBandwidthShapeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VirtualCircuitBandwidthShapeStatus{}),
		SDKStructs: []string{
			"core.VirtualCircuitBandwidthShape",
		},
	},
	{
		Name:       "CoreVirtualCircuitPublicPrefix",
		SpecType:   reflect.TypeOf(corev1beta1.VirtualCircuitPublicPrefixSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VirtualCircuitPublicPrefixStatus{}),
		SDKStructs: []string{
			"core.CreateVirtualCircuitPublicPrefixDetails",
			"core.VirtualCircuitPublicPrefix",
		},
	},
	{
		Name:       "CoreVnicAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.VnicAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VnicAttachmentStatus{}),
		SDKStructs: []string{
			"core.VnicAttachment",
		},
	},
	{
		Name:       "CoreVolume",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeDetails",
			"core.UpdateVolumeDetails",
			"core.Volume",
		},
	},
	{
		Name:       "CoreVolumeAttachment",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeAttachmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeAttachmentStatus{}),
		SDKStructs: []string{
			"core.UpdateVolumeAttachmentDetails",
		},
	},
	{
		Name:       "CoreVolumeBackup",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeBackupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeBackupStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeBackupDetails",
			"core.UpdateVolumeBackupDetails",
			"core.VolumeBackup",
		},
	},
	{
		Name:       "CoreVolumeBackupPolicy",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeBackupPolicySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeBackupPolicyStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeBackupPolicyDetails",
			"core.UpdateVolumeBackupPolicyDetails",
			"core.VolumeBackupPolicy",
		},
	},
	{
		Name:       "CoreVolumeBackupPolicyAssetAssignment",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeBackupPolicyAssetAssignmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeBackupPolicyAssetAssignmentStatus{}),
		SDKStructs: []string{
			"core.VolumeBackupPolicyAssignment",
		},
	},
	{
		Name:       "CoreVolumeBackupPolicyAssignment",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeBackupPolicyAssignmentSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeBackupPolicyAssignmentStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeBackupPolicyAssignmentDetails",
			"core.VolumeBackupPolicyAssignment",
		},
	},
	{
		Name:       "CoreVolumeGroup",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeGroupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeGroupStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeGroupDetails",
			"core.UpdateVolumeGroupDetails",
			"core.VolumeGroup",
		},
	},
	{
		Name:       "CoreVolumeGroupBackup",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeGroupBackupSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeGroupBackupStatus{}),
		SDKStructs: []string{
			"core.CreateVolumeGroupBackupDetails",
			"core.UpdateVolumeGroupBackupDetails",
			"core.VolumeGroupBackup",
		},
	},
	{
		Name:       "CoreVolumeGroupReplica",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeGroupReplicaSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeGroupReplicaStatus{}),
		SDKStructs: []string{
			"core.VolumeGroupReplicaDetails",
			"core.VolumeGroupReplica",
		},
	},
	{
		Name:       "CoreVolumeKMSKey",
		SpecType:   reflect.TypeOf(corev1beta1.VolumeKmsKeySpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VolumeKmsKeyStatus{}),
		SDKStructs: []string{
			"core.UpdateVolumeKmsKeyDetails",
			"core.VolumeKmsKey",
		},
	},
	{
		Name:       "CoreVtap",
		SpecType:   reflect.TypeOf(corev1beta1.VtapSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.VtapStatus{}),
		SDKStructs: []string{
			"core.CreateVtapDetails",
			"core.UpdateVtapDetails",
			"core.Vtap",
		},
	},
	{
		Name:       "CoreWindowsInstanceInitialCredential",
		SpecType:   reflect.TypeOf(corev1beta1.WindowsInstanceInitialCredentialSpec{}),
		StatusType: reflect.TypeOf(corev1beta1.WindowsInstanceInitialCredentialStatus{}),
		SDKStructs: []string{
			"core.InstanceCredentials",
		},
	},
	{
		Name:       "WorkrequestsWorkRequest",
		SpecType:   reflect.TypeOf(workrequestsv1beta1.WorkRequestSpec{}),
		StatusType: reflect.TypeOf(workrequestsv1beta1.WorkRequestStatus{}),
		SDKStructs: []string{
			"workrequests.WorkRequest",
			"workrequests.WorkRequestSummary",
		},
	},
	{
		Name:       "WorkrequestsWorkRequestError",
		SpecType:   reflect.TypeOf(workrequestsv1beta1.WorkRequestErrorSpec{}),
		StatusType: reflect.TypeOf(workrequestsv1beta1.WorkRequestErrorStatus{}),
		SDKStructs: []string{
			"workrequests.WorkRequestError",
		},
	},
	{
		Name:       "WorkrequestsWorkRequestLog",
		SpecType:   reflect.TypeOf(workrequestsv1beta1.WorkRequestLogSpec{}),
		StatusType: reflect.TypeOf(workrequestsv1beta1.WorkRequestLogStatus{}),
		SDKStructs: []string{
			"workrequests.WorkRequestLogEntry",
		},
	},
}

func Targets() []Target {
	result := make([]Target, len(targets))
	copy(result, targets)
	return result
}
