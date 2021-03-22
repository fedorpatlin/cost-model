package cloud

import (
	"fmt"
	"io"

	v1 "k8s.io/api/core/v1"
)

type YandexProvider struct{}

func (yp *YandexProvider) ClusterInfo() (map[string]string, error) {
	return map[string]string{"": ""}, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetAddresses() ([]byte, error) {
	return []byte(""), fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetDisks() ([]byte, error) {
	return []byte(""), fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) NodePricing(Key) (*Node, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) PVPricing(PVKey) (*PV, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) NetworkPricing() (*Network, error) {
	return nil, fmt.Errorf("Not yet implemented")
} // TODO: add key interface arg for dynamic price fetching
func (yp *YandexProvider) LoadBalancerPricing() (*LoadBalancer, error) {
	return nil, fmt.Errorf("Not yet implemented")
} // TODO: add key interface arg for dynamic price fetching
func (yp *YandexProvider) AllNodePricing() (interface{}, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) DownloadPricingData() error {
	return fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetKey(map[string]string, *v1.Node) Key {
	return nil
}
func (yp *YandexProvider) GetPVKey(*v1.PersistentVolume, map[string]string, string) PVKey {
	return nil
}
func (yp *YandexProvider) UpdateConfig(r io.Reader, updateType string) (*CustomPricing, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) UpdateConfigFromConfigMap(map[string]string) (*CustomPricing, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetConfig() (*CustomPricing, error) {
	return nil, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetManagementPlatform() (string, error) {
	return "", fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) GetLocalStorageQuery(string, string, bool, bool) string {
	return ""
}
func (yp *YandexProvider) ExternalAllocations(string, string, []string, string, string, bool) ([]*OutOfClusterAllocation, error) {
	return []*OutOfClusterAllocation{}, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) ApplyReservedInstancePricing(map[string]*Node) {}
func (yp *YandexProvider) ServiceAccountStatus() *ServiceAccountStatus {
	return nil
}
func (yp *YandexProvider) PricingSourceStatus() map[string]*PricingSource {
	return map[string]*PricingSource{"": nil}
}
func (yp *YandexProvider) ClusterManagementPricing() (string, float64, error) {
	return "", 0, fmt.Errorf("Not yet implemented")
}
func (yp *YandexProvider) CombinedDiscountForNode(string, bool, float64, float64) float64 {
	return 0
}
func (yp *YandexProvider) ParseID(string) string {
	return ""
}
func (yp *YandexProvider) ParsePVID(string) string {
	return ""
}
