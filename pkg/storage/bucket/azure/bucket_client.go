package azure

import (
	"github.com/go-kit/log"
	"github.com/thanos-io/thanos/pkg/httpconfig"
	"github.com/thanos-io/thanos/pkg/objstore"
	"github.com/thanos-io/thanos/pkg/objstore/azure"
	yaml "gopkg.in/yaml.v2"
)

func NewBucketClient(cfg Config, name string, logger log.Logger) (objstore.Bucket, error) {
	bucketConfig := azure.Config{
		StorageAccountName: cfg.StorageAccountName,
		StorageAccountKey:  cfg.StorageAccountKey.Value,
		ContainerName:      cfg.ContainerName,
		Endpoint:           cfg.Endpoint,
		MaxRetries:         cfg.MaxRetries,
		HTTPConfig: azure.HTTPConfig{
			TransportConfig: httpconfig.TransportConfig{
				IdleConnTimeout:       int64(cfg.IdleConnTimeout),
				ResponseHeaderTimeout: int64(cfg.ResponseHeaderTimeout),
				TLSHandshakeTimeout:   int64(cfg.TLSHandshakeTimeout),
				MaxIdleConns:          cfg.MaxIdleConns,
				MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
				MaxConnsPerHost:       cfg.MaxConnsPerHost,
			},
			TLSConfig: httpconfig.TLSConfig{
				InsecureSkipVerify: cfg.InsecureSkipVerify,
			},
		},
	}

	// Thanos currently doesn't support passing the config as is, but expects a YAML,
	// so we're going to serialize it.
	serialized, err := yaml.Marshal(bucketConfig)
	if err != nil {
		return nil, err
	}

	return azure.NewBucket(logger, serialized, name)
}
