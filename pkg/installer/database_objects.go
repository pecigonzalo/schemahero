package installer

var generatedDatabaseCRDV1Beta1 = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.8
  creationTimestamp: null
  name: databases.databases.schemahero.io
spec:
  group: databases.schemahero.io
  names:
    kind: Database
    listKind: DatabaseList
    plural: databases
    singular: database
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: Database is the Schema for the databases API
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
            connection:
              description: DatabaseConnection defines connection parameters for the
                database driver
              properties:
                cockroachdb:
                  properties:
                    uri:
                      properties:
                        value:
                          type: string
                        valueFrom:
                          properties:
                            secretKeyRef:
                              properties:
                                key:
                                  type: string
                                name:
                                  type: string
                              required:
                              - key
                              - name
                              type: object
                          type: object
                      type: object
                  type: object
                mysql:
                  properties:
                    uri:
                      properties:
                        value:
                          type: string
                        valueFrom:
                          properties:
                            secretKeyRef:
                              properties:
                                key:
                                  type: string
                                name:
                                  type: string
                              required:
                              - key
                              - name
                              type: object
                          type: object
                      type: object
                  type: object
                postgres:
                  properties:
                    uri:
                      properties:
                        value:
                          type: string
                        valueFrom:
                          properties:
                            secretKeyRef:
                              properties:
                                key:
                                  type: string
                                name:
                                  type: string
                              required:
                              - key
                              - name
                              type: object
                          type: object
                      type: object
                  type: object
              type: object
            immediateDeploy:
              type: boolean
            schemahero:
              properties:
                image:
                  type: string
                nodeSelector:
                  additionalProperties:
                    type: string
                  type: object
              type: object
          type: object
        status:
          description: DatabaseStatus defines the observed state of Database
          properties:
            isConnected:
              type: boolean
            lastPing:
              type: string
          required:
          - isConnected
          - lastPing
          type: object
      required:
      - spec
      type: object
  version: v1alpha3
  versions:
  - name: v1alpha3
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
var generatedDatabaseCRDV1 = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.8
  creationTimestamp: null
  name: databases.databases.schemahero.io
spec:
  group: databases.schemahero.io
  names:
    kind: Database
    listKind: DatabaseList
    plural: databases
    singular: database
  scope: Namespaced
  versions:
  - name: v1alpha3
    schema:
      openAPIV3Schema:
        description: Database is the Schema for the databases API
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
              connection:
                description: DatabaseConnection defines connection parameters for
                  the database driver
                properties:
                  cockroachdb:
                    properties:
                      uri:
                        properties:
                          value:
                            type: string
                          valueFrom:
                            properties:
                              secretKeyRef:
                                properties:
                                  key:
                                    type: string
                                  name:
                                    type: string
                                required:
                                - key
                                - name
                                type: object
                            type: object
                        type: object
                    type: object
                  mysql:
                    properties:
                      uri:
                        properties:
                          value:
                            type: string
                          valueFrom:
                            properties:
                              secretKeyRef:
                                properties:
                                  key:
                                    type: string
                                  name:
                                    type: string
                                required:
                                - key
                                - name
                                type: object
                            type: object
                        type: object
                    type: object
                  postgres:
                    properties:
                      uri:
                        properties:
                          value:
                            type: string
                          valueFrom:
                            properties:
                              secretKeyRef:
                                properties:
                                  key:
                                    type: string
                                  name:
                                    type: string
                                required:
                                - key
                                - name
                                type: object
                            type: object
                        type: object
                    type: object
                type: object
              immediateDeploy:
                type: boolean
              schemahero:
                properties:
                  image:
                    type: string
                  nodeSelector:
                    additionalProperties:
                      type: string
                    type: object
                type: object
            type: object
          status:
            description: DatabaseStatus defines the observed state of Database
            properties:
              isConnected:
                type: boolean
              lastPing:
                type: string
            required:
            - isConnected
            - lastPing
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
