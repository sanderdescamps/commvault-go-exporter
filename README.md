# commvault-go-exporter
Prometheus exporter for Commvault


## Run Exporter

    Usage:
      commvault-go-exporter [flags]

    Flags:
      -e, --endpoint string   Commvault endpoint
      -h, --help              help for commvault-go-exporter
          --insecure          Skip SSL certificate verification
      -p, --password string   Password to connect to Commvault
      -l, --port int16        Port the metricserver will listen on (default 2112)
      -u, --user string       Username to connect to Commvault
  
#### example

    commvault-go-exporter --endpoint https://cs01.example.com --user testuser
    commvault-go-exporter --port 3000 --endpoint https://cs01.example.com --user testuser --password supersecure

## Metrics

    curl localhost:2112/metrics