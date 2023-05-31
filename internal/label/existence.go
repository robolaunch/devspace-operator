package label

import (
	devspaceErr "github.com/robolaunch/devspace-operator/internal/error"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CheckLabelExistence(objMeta metav1.ObjectMeta, typeMeta metav1.TypeMeta, labelKeys []string) error {
	labels := objMeta.GetLabels()
	for _, key := range labelKeys {
		if _, ok := labels[key]; !ok {
			return &devspaceErr.LabelNotFoundError{
				LabelKey:          key,
				ResourceKind:      typeMeta.Kind,
				ResourceName:      objMeta.Name,
				ResourceNamespace: objMeta.Namespace,
			}
		}
	}
	return nil
}
