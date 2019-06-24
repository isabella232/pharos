package api

import (
	"fmt"
	"net/http"

	"github.com/lob/pharos/pkg/util/model"
	"github.com/pkg/errors"
)

// ListClusters sends a GET request to the clusters endpoint of the Pharos API
// and returns an array of Clusters. Can also be called with query to retrieve
// a certain subset of clusters.
func (c *Client) ListClusters(query map[string]string) ([]model.Cluster, error) {
	var clusters []model.Cluster
	err := c.send(http.MethodGet, "clusters", query, nil, &clusters)
	if err != nil {
		return clusters, errors.Wrap(err, "failed to list clusters")
	}

	return clusters, nil
}

// GetCluster sends a GET request to the clusters/id endpoint of the Pharos API
// and returns a Cluster.
func (c *Client) GetCluster(clusterID string) (model.Cluster, error) {
	var cluster model.Cluster
	err := c.send(http.MethodGet, fmt.Sprintf("clusters/%s", clusterID), nil, nil, &cluster)
	if err != nil {
		return cluster, errors.Wrap(err, "failed to get cluster")
	}

	return cluster, nil
}