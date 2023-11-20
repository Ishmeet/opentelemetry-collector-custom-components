# Opentelemetry collector custom components

 Hi there,
 This repository develop and host components and include it in custom build of the opentelemetry collector using the collector builder

## Current components

 - sflowreceiver

## Build and Run

 Install the opentelemetry collector builder 
 ```
 GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@latest
 ```
 
 this will generate the collector binary `otelcol`

 ```
 ./otelcol --config config.yaml
 ```

## Config

 I have used a default config (as mentioned below), please change as per your needs. I have specified all of the receivers, processors, exporters and connectors in the otelcol-builder.yaml as per version v0.89.0 of opentelemetry-collector-contrib package. You can enable or disable them as per your needs.
 
```yaml
receivers:
  sflow:
    endpoint: 0.0.0.0:9995
    labels:
      mylabel1: value1
      mylabel2: value2

exporters:
  logging:
    verbosity: detailed

service:
  pipelines:
    logs:
      receivers:
        - sflow
      exporters:
        - logging
```