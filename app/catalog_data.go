package app

import (
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/config"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultURL = "https://github.com/ibuildthecloud/rancher-charts"
)

func addCatalogs(management *config.ManagementContext) error {
	catalogClient := management.Management.Catalogs("")
	_, err := catalogClient.Get("helm", metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	} else if errors.IsNotFound(err) {
		obj := &v3.Catalog{
			ObjectMeta: metav1.ObjectMeta{
				Name: "helm",
			},
			Spec: v3.CatalogSpec{
				URL:         defaultURL,
				CatalogKind: "helm",
			},
		}
		if _, err := catalogClient.Create(obj); err != nil {
			return err
		}
	}
	return nil
}
