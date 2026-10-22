package mongodb

// MongoDB collection names used by free5GC NFs and Webconsole.
//
// Keep the existing names and capitalization unchanged in Phase 1.
// Collection names belong only in the MongoDB implementation layer;
// dbif interfaces and NF code must not depend on them.
const (
	// Application data (UDR).
	collApplicationInfluenceData         = "applicationData.influenceData"
	collApplicationInfluenceSubsToNotify = "applicationData.influenceData.subsToNotify"
	collApplicationPfds                  = "applicationData.pfds"

	// Policy data (UDR, PCF, CHF, Webconsole).
	collPolicyBdtData                 = "policyData.bdtData"
	collPolicyPlmnsUePolicySet        = "policyData.plmns.uePolicySet"
	collPolicySponsorConnectivityData = "policyData.sponsorConnectivityData"
	collPolicyUesAmData               = "policyData.ues.amData"
	collPolicyUesChargingData         = "policyData.ues.chargingData"
	collPolicyUesFlowRule             = "policyData.ues.flowRule"
	collPolicyUesOperatorSpecificData = "policyData.ues.operatorSpecificData"
	collPolicyUesQosFlow              = "policyData.ues.qosFlow"
	collPolicyUesSmData               = "policyData.ues.smData"
	collPolicyUesUsageMonData         = "policyData.ues.smData.usageMonData"
	collPolicyUesUePolicySet          = "policyData.ues.uePolicySet"

	// Subscription data (UDR, SCP, Webconsole).
	collSubscriptionAuthenticationStatus          = "subscriptionData.authenticationData.authenticationStatus"
	collSubscriptionAuthenticationSubscription    = "subscriptionData.authenticationData.authenticationSubscription"
	collSubscriptionWebAuthenticationSubscription = "subscriptionData.authenticationData.webAuthenticationSubscription"
	collSubscriptionAmf3gppAccess                 = "subscriptionData.contextData.amf3gppAccess"
	collSubscriptionAmfNon3gppAccess              = "subscriptionData.contextData.amfNon3gppAccess"
	collSubscriptionSmfRegistrations              = "subscriptionData.contextData.smfRegistrations"
	collSubscriptionSmsf3gppAccess                = "subscriptionData.contextData.smsf3gppAccess"
	collSubscriptionSmsfNon3gppAccess             = "subscriptionData.contextData.smsfNon3gppAccess"
	collSubscriptionEeProfileData                 = "subscriptionData.eeProfileData"
	collSubscriptionIdentityData                  = "subscriptionData.identityData"
	collSubscriptionOperatorDeterminedBarringData = "subscriptionData.operatorDeterminedBarringData"
	collSubscriptionOperatorSpecificData          = "subscriptionData.operatorSpecificData"
	collSubscriptionPpData                        = "subscriptionData.ppData"
	collSubscriptionProvisionedAmData             = "subscriptionData.provisionedData.amData"
	collSubscriptionProvisionedSmData             = "subscriptionData.provisionedData.smData"
	collSubscriptionProvisionedSmfSelectionData   = "subscriptionData.provisionedData.smfSelectionSubscriptionData"
	collSubscriptionProvisionedSmsData            = "subscriptionData.provisionedData.smsData"
	collSubscriptionProvisionedSmsMngData         = "subscriptionData.provisionedData.smsMngData"
	collSubscriptionProvisionedTraceData          = "subscriptionData.provisionedData.traceData"
	collSubscriptionSharedData                    = "subscriptionData.sharedData"
	collSubscriptionSorData                       = "subscriptionData.ueUpdateConfirmationData.sorData"

	// NRF, BSF, and Webconsole-specific collections.
	collNfProfile     = "NfProfile"
	collPcfBindings   = "pcfBindings"
	collProfileData   = "profileData"
	collSubscriptions = "Subscriptions"
	collTenantData    = "tenantData"
	collUriList       = "urilist"
	collUserData      = "userData"
)
