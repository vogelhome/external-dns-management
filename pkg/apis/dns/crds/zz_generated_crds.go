// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package crds

import (
	"github.com/gardener/controller-manager-library/pkg/resources/apiextensions"
	"github.com/gardener/controller-manager-library/pkg/utils"
)

var registry = apiextensions.NewRegistry()

func init() {
	var data string
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnsannotations.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSAnnotation
    listKind: DNSAnnotationList
    plural: dnsannotations
    shortNames:
    - dnsa
    singular: dnsannotation
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.resourceRef.apiVersion
      name: RefGroup
      type: string
    - jsonPath: .spec.resourceRef.kind
      name: RefKind
      type: string
    - jsonPath: .spec.resourceRef.name
      name: RefName
      type: string
    - jsonPath: .spec.resourceRef.namespace
      name: RefNamespace
      type: string
    - jsonPath: .status.active
      name: Active
      type: boolean
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              annotations:
                additionalProperties:
                  type: string
                type: object
              resourceRef:
                properties:
                  apiVersion:
                    description: API Version of the annotated object
                    type: string
                  kind:
                    description: 'Kind of the annotated object More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: Name of the annotated object
                    type: string
                  namespace:
                    description: Namspace of the annotated object Defaulted by the
                      namespace of the containing resource.
                    type: string
                required:
                - apiVersion
                - kind
                type: object
            required:
            - annotations
            - resourceRef
            type: object
          status:
            properties:
              active:
                description: Indicates that annotation is observed by a DNS sorce
                  controller
                type: boolean
              message:
                description: In case of a configuration problem this field describes
                  the reason
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnsentries.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSEntry
    listKind: DNSEntryList
    plural: dnsentries
    shortNames:
    - dnse
    singular: dnsentry
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: FQDN of DNS Entry
      jsonPath: .spec.dnsName
      name: DNS
      type: string
    - description: provider type
      jsonPath: .status.providerType
      name: TYPE
      type: string
    - description: assigned provider (namespace/name)
      jsonPath: .status.provider
      name: PROVIDER
      type: string
    - description: entry status
      jsonPath: .status.state
      name: STATUS
      type: string
    - description: entry creation timestamp
      jsonPath: .metadata.creationTimestamp
      name: AGE
      type: date
    - description: effective targets
      jsonPath: .status.targets
      name: TARGETS
      type: string
    - description: owner id used to tag entries in external DNS system
      jsonPath: .spec.ownerId
      name: OWNERID
      type: string
    - description: time to live
      jsonPath: .status.ttl
      name: TTL
      priority: 2000
      type: integer
    - description: zone id
      jsonPath: .status.zone
      name: ZONE
      priority: 2000
      type: string
    - description: routing policy type
      jsonPath: .status.routingPolicy.type
      name: POLICY_TYPE
      priority: 2000
      type: string
    - description: routing policy set identifier
      jsonPath: .status.routingPolicy.setIdentifier
      name: POLICY_SETID
      priority: 2000
      type: string
    - description: routing policy parameters
      jsonPath: .status.routingPolicy.parameters
      name: POLICY_PARAMS
      priority: 2000
      type: string
    - description: message describing the reason for the state
      jsonPath: .status.message
      name: MESSAGE
      priority: 2000
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              cnameLookupInterval:
                description: lookup interval for CNAMEs that must be resolved to IP
                  addresses
                format: int64
                type: integer
              dnsName:
                description: full qualified domain name
                type: string
              ownerId:
                description: owner id used to tag entries in external DNS system
                type: string
              reference:
                description: reference to base entry used to inherit attributes from
                properties:
                  name:
                    description: name of the referenced DNSEntry object
                    type: string
                  namespace:
                    description: namespace of the referenced DNSEntry object
                    type: string
                required:
                - name
                type: object
              routingPolicy:
                description: optional routing policy
                properties:
                  parameters:
                    additionalProperties:
                      type: string
                    description: Policy specific parameters
                    type: object
                  setIdentifier:
                    description: SetIdentifier is the identifier of the record set
                    type: string
                  type:
                    description: Policy is the policy type. Allowed values are provider
                      dependent, e.g. ` + "`" + `weighted` + "`" + `
                    type: string
                required:
                - parameters
                - setIdentifier
                - type
                type: object
              targets:
                description: target records (CNAME or A records), either text or targets
                  must be specified
                items:
                  type: string
                type: array
              text:
                description: text records, either text or targets must be specified
                items:
                  type: string
                type: array
              ttl:
                description: time to live for records in external DNS system
                format: int64
                type: integer
            required:
            - dnsName
            type: object
          status:
            properties:
              lastUpdateTime:
                description: lastUpdateTime contains the timestamp of the last status
                  update
                format: date-time
                type: string
              message:
                description: message describing the reason for the state
                type: string
              observedGeneration:
                format: int64
                type: integer
              provider:
                description: assigned provider
                type: string
              providerType:
                description: provider type used for the entry
                type: string
              routingPolicy:
                description: effective routing policy
                properties:
                  parameters:
                    additionalProperties:
                      type: string
                    description: Policy specific parameters
                    type: object
                  setIdentifier:
                    description: SetIdentifier is the identifier of the record set
                    type: string
                  type:
                    description: Policy is the policy type. Allowed values are provider
                      dependent, e.g. ` + "`" + `weighted` + "`" + `
                    type: string
                required:
                - parameters
                - setIdentifier
                - type
                type: object
              state:
                description: entry state
                type: string
              targets:
                description: effective targets generated for the entry
                items:
                  type: string
                type: array
              ttl:
                description: time to live used for the entry
                format: int64
                type: integer
              zone:
                description: zone used for the entry
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnshostedzonepolicies.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSHostedZonePolicy
    listKind: DNSHostedZonePolicyList
    plural: dnshostedzonepolicies
    shortNames:
    - dnshzp
    singular: dnshostedzonepolicy
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .status.count
      name: Zone Count
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              policy:
                description: ZonePolicy specifies zone specific policy
                properties:
                  zoneStateCacheTTL:
                    description: ZoneStateCacheTTL specifies the TTL for the zone
                      state cache
                    type: string
                type: object
              selector:
                description: ZoneSelector specifies the selector for the DNS hosted
                  zones
                properties:
                  domainNames:
                    description: DomainNames selects by base domain name of hosted
                      zone. Policy will be applied to zones with matching base domain
                    items:
                      type: string
                    type: array
                  providerTypes:
                    description: ProviderTypes selects by provider types
                    items:
                      type: string
                    type: array
                  zoneIDs:
                    description: ZoneIDs selects by provider dependent zone ID
                    items:
                      type: string
                    type: array
                type: object
            required:
            - policy
            - selector
            type: object
          status:
            properties:
              count:
                description: Number of zones this policy is applied to
                type: integer
              lastStatusUpdateTime:
                description: LastStatusUpdateTime contains the timestamp of the last
                  status update
                format: date-time
                type: string
              message:
                description: In case of a configuration problem this field describes
                  the reason
                type: string
              zones:
                description: Indicates that annotation is observed by a DNS sorce
                  controller
                items:
                  properties:
                    domainName:
                      description: Domain name of the zone
                      type: string
                    providerType:
                      description: Provider type of the zone
                      type: string
                    zoneID:
                      description: ID of the zone
                      type: string
                  required:
                  - domainName
                  - providerType
                  - zoneID
                  type: object
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnslocks.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSLock
    listKind: DNSLockList
    plural: dnslocks
    shortNames:
    - dnsl
    singular: dnslock
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: FQDN of DNS Entry
      jsonPath: .spec.dnsName
      name: DNS
      type: string
    - description: provider type
      jsonPath: .status.providerType
      name: TYPE
      type: string
    - description: assigned provider (namespace/name)
      jsonPath: .status.provider
      name: PROVIDER
      type: string
    - description: entry status
      jsonPath: .status.state
      name: STATUS
      type: string
    - description: entry creation timestamp
      jsonPath: .metadata.creationTimestamp
      name: AGE
      type: date
    - description: owner group id used to tag entries in external DNS system
      jsonPath: .spec.ownerGroupId
      name: OWNERID
      type: string
    - description: time to live
      jsonPath: .status.ttl
      name: TTL
      priority: 2000
      type: integer
    - description: zone id
      jsonPath: .status.zone
      name: ZONE
      priority: 2000
      type: string
    - description: message describing the reason for the state
      jsonPath: .status.message
      name: MESSAGE
      priority: 2000
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              attributes:
                additionalProperties:
                  type: string
                description: attribute values (must be compatible with DNS TXT records)
                type: object
              dnsName:
                description: full qualified domain name
                type: string
              lockId:
                description: owner group for collaboration of multiple controller
                type: string
              timestamp:
                description: Activation time stamp
                format: date-time
                type: string
              ttl:
                description: time to live for records in external DNS system
                format: int64
                type: integer
            required:
            - dnsName
            - timestamp
            - ttl
            type: object
          status:
            properties:
              attributes:
                additionalProperties:
                  type: string
                description: attribute values found in DNS
                type: object
              firstFailedDNSLookup:
                description: First failed DNS looup
                format: date-time
                type: string
              lastUpdateTime:
                description: lastUpdateTime contains the timestamp of the last status
                  update
                format: date-time
                type: string
              lockId:
                description: owner group for collaboration of multiple controller
                  found in DNS
                type: string
              message:
                description: message describing the reason for the state
                type: string
              observedGeneration:
                format: int64
                type: integer
              provider:
                description: assigned provider
                type: string
              providerType:
                description: provider type used for the entry
                type: string
              state:
                description: entry state
                type: string
              timestamp:
                description: Activation time stamp found in DNS
                format: date-time
                type: string
              ttl:
                description: time to live used for the entry
                format: int64
                type: integer
              zone:
                description: zone used for the entry
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnsowners.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSOwner
    listKind: DNSOwnerList
    plural: dnsowners
    shortNames:
    - dnso
    singular: dnsowner
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.ownerId
      name: OwnerId
      type: string
    - jsonPath: .status.active
      name: Active
      type: boolean
    - jsonPath: .status.entries.amount
      name: Usages
      type: integer
    - description: expiration date
      format: date-time
      jsonPath: .spec.validUntil
      name: Valid
      type: string
    - description: creation timestamp
      jsonPath: .metadata.creationTimestamp
      name: AGE
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              active:
                description: state of the ownerid for the DNS controller observing
                  entry using this owner id (default:true)
                type: boolean
              dnsActivation:
                description: Optional activation info for controlling the owner activation
                  remotely via DNS TXT record
                properties:
                  dnsName:
                    description: DNS name for controlling the owner activation remotely
                      via DNS TXT record
                    type: string
                  value:
                    description: Optional value for the DNS activation record used
                      to activate this owner The default is the id of the cluster
                      used to read the owner object
                    type: string
                required:
                - dnsName
                type: object
              ownerId:
                description: owner id used to tag entries in external DNS system
                type: string
              validUntil:
                description: optional time this owner should be active if active flag
                  is not false
                format: date-time
                type: string
            required:
            - ownerId
            type: object
          status:
            properties:
              active:
                description: state of the ownerid for the DNS controller observing
                  entry using this owner id
                type: boolean
              entries:
                description: Entry statistic for this owner id
                properties:
                  amount:
                    description: number of entries using this owner id
                    type: integer
                  types:
                    additionalProperties:
                      type: integer
                    description: number of entries per provider type
                    type: object
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: dnsproviders.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: DNSProvider
    listKind: DNSProviderList
    plural: dnsproviders
    shortNames:
    - dnspr
    singular: dnsprovider
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.type
      name: TYPE
      type: string
    - jsonPath: .status.state
      name: STATUS
      type: string
    - description: creation timestamp
      jsonPath: .metadata.creationTimestamp
      name: AGE
      type: date
    - description: included domains
      jsonPath: .status.domains.included
      name: INCLUDED_DOMAINS
      type: string
    - description: included zones
      jsonPath: .status.zones.included
      name: INCLUDED_ZONES
      priority: 2000
      type: string
    - description: message describing the reason for the state
      jsonPath: .status.message
      name: MESSAGE
      priority: 2000
      type: string
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              defaultTTL:
                description: default TTL used for DNS entries if not specified explicitly
                format: int64
                type: integer
              domains:
                description: desired selection of usable domains (by default all zones
                  and domains in those zones will be served)
                properties:
                  exclude:
                    description: values that should be ignored (domains or zones)
                    items:
                      type: string
                    type: array
                  include:
                    description: values that should be observed (domains or zones)
                    items:
                      type: string
                    type: array
                type: object
              providerConfig:
                description: optional additional provider specific configuration values
                type: object
                x-kubernetes-preserve-unknown-fields: true
              rateLimit:
                description: rate limit for create/update operations on DNSEntries
                  assigned to this provider
                properties:
                  burst:
                    description: Burst allows bursts of up to 'burst' to exceed the
                      rate defined by 'RequestsPerDay', while still maintaining a
                      smoothed rate of 'RequestsPerDay'
                    type: integer
                  requestsPerDay:
                    description: RequestsPerDay is create/update request rate per
                      DNS entry given by requests per day
                    type: integer
                required:
                - burst
                - requestsPerDay
                type: object
              secretRef:
                description: access credential for the external DNS system of the
                  given type
                properties:
                  name:
                    description: name is unique within a namespace to reference a
                      secret resource.
                    type: string
                  namespace:
                    description: namespace defines the space within which the secret
                      name must be unique.
                    type: string
                type: object
                x-kubernetes-map-type: atomic
              type:
                description: type of the provider (selecting the responsible type
                  of DNS controller)
                type: string
              zones:
                description: desired selection of usable domains the domain selection
                  is used for served zones, only (by default all zones will be served)
                properties:
                  exclude:
                    description: values that should be ignored (domains or zones)
                    items:
                      type: string
                    type: array
                  include:
                    description: values that should be observed (domains or zones)
                    items:
                      type: string
                    type: array
                type: object
            type: object
          status:
            properties:
              defaultTTL:
                description: actually used default TTL for DNS entries
                format: int64
                type: integer
              domains:
                description: actually served domain selection
                properties:
                  excluded:
                    description: Excluded values (domains or zones)
                    items:
                      type: string
                    type: array
                  included:
                    description: included values (domains or zones)
                    items:
                      type: string
                    type: array
                type: object
              lastUpdateTime:
                description: lastUpdateTime contains the timestamp of the last status
                  update
                format: date-time
                type: string
              message:
                description: message describing the reason for the actual state of
                  the provider
                type: string
              observedGeneration:
                format: int64
                type: integer
              rateLimit:
                description: actually used rate limit for create/update operations
                  on DNSEntries assigned to this provider
                properties:
                  burst:
                    description: Burst allows bursts of up to 'burst' to exceed the
                      rate defined by 'RequestsPerDay', while still maintaining a
                      smoothed rate of 'RequestsPerDay'
                    type: integer
                  requestsPerDay:
                    description: RequestsPerDay is create/update request rate per
                      DNS entry given by requests per day
                    type: integer
                required:
                - burst
                - requestsPerDay
                type: object
              state:
                description: state of the provider
                type: string
              zones:
                description: actually served zones
                properties:
                  excluded:
                    description: Excluded values (domains or zones)
                    items:
                      type: string
                    type: array
                  included:
                    description: included values (domains or zones)
                    items:
                      type: string
                    type: array
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
	data = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.13.0
  name: remoteaccesscertificates.dns.gardener.cloud
spec:
  group: dns.gardener.cloud
  names:
    kind: RemoteAccessCertificate
    listKind: RemoteAccessCertificateList
    plural: remoteaccesscertificates
    shortNames:
    - remotecert
    singular: remoteaccesscertificate
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.type
      name: Type
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - jsonPath: .status.notBefore
      name: SecretAge
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              days:
                description: Number of days the certificate should be valid
                type: integer
              domainName:
                description: Domain name, used for building subject and DNS name
                type: string
              recreate:
                description: Indicates if certificate should be recreated and replaced
                  in the secret
                type: boolean
              secretName:
                description: Name of the secret to store the client certificate
                type: string
              type:
                description: Certificate type (client or server)
                type: string
            required:
            - days
            - domainName
            - secretName
            - type
            type: object
          status:
            properties:
              message:
                description: In case of a configuration problem this field describes
                  the reason
                type: string
              notAfter:
                description: Expiration timestamp of the certificate
                format: date-time
                type: string
              notBefore:
                description: Creation timestamp of the certificate
                format: date-time
                type: string
              recreating:
                description: Indicates if certificate should be recreated and replaced
                  in the secret
                type: boolean
              serialNumber:
                description: Serial number of the certificate
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  `
	utils.Must(registry.RegisterCRD(data))
}

func AddToRegistry(r apiextensions.Registry) {
	registry.AddToRegistry(r)
}
