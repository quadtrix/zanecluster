// package cluster implements backend cluster logic
package cluster

import (
	"crypto/x509"

	"github.com/quadtrix/zanecluster/backend/types"
)

const (
	ENC_MD5   types.Encryption = 0
	ENC_RSA   types.Encryption = 1
	ENC_CRYPT types.Encryption = 2
	ENC_PLAIN types.Encryption = 99
)

func NewCluster(secure bool, secConfigFile string, clusterNet types.Network) (*types.Cluster, error) {
	cl := new(types.Cluster)
	cl.Secure = secure
	if cl.Secure {
		secconfig, err := NewSecureConfig(secConfigFile)
		if err != nil {
			return nil, err
		}
		cl.SecureConfig = secconfig
	}
	cl.ClusterNet = clusterNet
	return cl, nil
}

func loadConfig(cfg string) (*types.ClusterConfig, error) {
	return new(types.ClusterConfig), nil
}

func NewSecureConfig(configfile string) (*types.SecureClusterConfig, error) {
	clusterConfig, err := loadConfig(configfile)
	if err != nil {
		return nil, err
	}
	scc := new(types.SecureClusterConfig)
	tcp, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	scc.TrustedCertPool = tcp
	scc.ClusterNet = types.Network{
		Net:  clusterConfig.Cluster.Network,
		Mask: clusterConfig.Cluster.NetMask,
		Port: clusterConfig.Cluster.Port,
	}
	scc.ClusterPassword = types.EncryptedClusterPassword{
		Encryption: clusterConfig.Cluster.Encryption,
		Pwd:        clusterConfig.Cluster.Password,
	}
	scc.ServerCert = clusterConfig.Cluster.SSLCertFile
	scc.ServerKey = clusterConfig.Cluster.SSLKeyFile
	return scc, nil
}

func (cl types.Cluster) PopulateCluster() error {

}
