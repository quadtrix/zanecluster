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
	cl.secure = secure
	if cl.secure {
		secconfig, err := NewSecureConfig(secConfigFile)
		if err != nil {
			return nil, err
		}
		cl.secureConfig = secconfig
	}
	cl.clusterNet = clusterNet
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
	scc.trustedCertPool = tcp
	scc.clusterNet = types.Network{
		net:  clusterConfig.Cluster.Network,
		mask: clusterConfig.Cluster.NetMask,
		port: clusterConfig.Cluster.Port,
	}
	scc.clusterPassword = types.EncryptedClusterPassword{
		encryption: clusterConfig.Cluster.Encryption,
		pwd:        clusterConfig.Cluster.Password,
	}
	scc.serverCert = clusterConfig.Cluster.SSLCertFile
	scc.serverKey = clusterConfig.Cluster.SSLKeyFile
	return scc, nil
}

func (cl types.Cluster) PopulateCluster() error {

}
