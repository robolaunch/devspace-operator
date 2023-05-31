package v1alpha1

import (
	"errors"

	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// ********************************
// DevSpaceIDE webhooks
// ********************************

// log is for logging in this package.
var devspaceidelog = logf.Log.WithName("devspaceide-resource")

func (r *DevSpaceIDE) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-devspaceide,mutating=true,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspaceides,verbs=create;update,versions=v1alpha1,name=mdevspaceide.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &DevSpaceIDE{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *DevSpaceIDE) Default() {
	devspaceidelog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-devspaceide,mutating=false,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspaceides,verbs=create;update,versions=v1alpha1,name=vdevspaceide.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DevSpaceIDE{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceIDE) ValidateCreate() error {
	devspaceidelog.Info("validate create", "name", r.Name)

	err := r.checkTargetDevspaceLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetDevSpaceVDILabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceIDE) ValidateUpdate(old runtime.Object) error {
	devspaceidelog.Info("validate update", "name", r.Name)

	err := r.checkTargetDevspaceLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetDevSpaceVDILabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceIDE) ValidateDelete() error {
	devspaceidelog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *DevSpaceIDE) checkTargetDevspaceLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_ROBOT_LABEL_KEY)
	}

	return nil
}

func (r *DevSpaceIDE) checkTargetDevSpaceVDILabel() error {
	labels := r.GetLabels()

	if r.Spec.Display {
		if _, ok := labels[internal.TARGET_VDI_LABEL_KEY]; !ok {
			return errors.New("target devspace vdi label should be added with key " + internal.TARGET_VDI_LABEL_KEY)
		}
	}

	return nil
}

// ********************************
// DevSpaceVDI webhooks
// ********************************

// log is for logging in this package.
var devspacevdilog = logf.Log.WithName("devspacevdi-resource")

func (r *DevSpaceVDI) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-devspacevdi,mutating=true,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspacevdis,verbs=create;update,versions=v1alpha1,name=mdevspacevdi.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &DevSpaceVDI{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *DevSpaceVDI) Default() {
	devspacevdilog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-devspacevdi,mutating=false,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspacevdis,verbs=create;update,versions=v1alpha1,name=vdevspacevdi.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DevSpaceVDI{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceVDI) ValidateCreate() error {
	devspacevdilog.Info("validate create", "name", r.Name)

	err := r.checkTargetDevspaceLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceVDI) ValidateUpdate(old runtime.Object) error {
	devspacevdilog.Info("validate update", "name", r.Name)

	err := r.checkTargetDevspaceLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceVDI) ValidateDelete() error {
	devspacevdilog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *DevSpaceVDI) checkTargetDevspaceLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_VDI_LABEL_KEY)
	}

	return nil
}
