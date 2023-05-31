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

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-devspaceide,mutating=true,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=devspaceides,verbs=create;update,versions=v1alpha1,name=mdevspaceide.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &DevSpaceIDE{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *DevSpaceIDE) Default() {
	devspaceidelog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-devspaceide,mutating=false,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=devspaceides,verbs=create;update,versions=v1alpha1,name=vdevspaceide.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DevSpaceIDE{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceIDE) ValidateCreate() error {
	devspaceidelog.Info("validate create", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetRobotVDILabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpaceIDE) ValidateUpdate(old runtime.Object) error {
	devspaceidelog.Info("validate update", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	err = r.checkTargetRobotVDILabel()
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

func (r *DevSpaceIDE) checkTargetRobotLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_ROBOT_LABEL_KEY)
	}

	return nil
}

func (r *DevSpaceIDE) checkTargetRobotVDILabel() error {
	labels := r.GetLabels()

	if r.Spec.Display {
		if _, ok := labels[internal.TARGET_VDI_LABEL_KEY]; !ok {
			return errors.New("target robot vdi label should be added with key " + internal.TARGET_VDI_LABEL_KEY)
		}
	}

	return nil
}

// ********************************
// RobotVDI webhooks
// ********************************

// log is for logging in this package.
var robotvdilog = logf.Log.WithName("robotvdi-resource")

func (r *RobotVDI) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-robot-roboscale-io-v1alpha1-robotvdi,mutating=true,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=robotvdis,verbs=create;update,versions=v1alpha1,name=mrobotvdi.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &RobotVDI{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *RobotVDI) Default() {
	robotvdilog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:path=/validate-robot-roboscale-io-v1alpha1-robotvdi,mutating=false,failurePolicy=fail,sideEffects=None,groups=robot.roboscale.io,resources=robotvdis,verbs=create;update,versions=v1alpha1,name=vrobotvdi.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &RobotVDI{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *RobotVDI) ValidateCreate() error {
	robotvdilog.Info("validate create", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *RobotVDI) ValidateUpdate(old runtime.Object) error {
	robotvdilog.Info("validate update", "name", r.Name)

	err := r.checkTargetRobotLabel()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *RobotVDI) ValidateDelete() error {
	robotvdilog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *RobotVDI) checkTargetRobotLabel() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.TARGET_ROBOT_LABEL_KEY]; !ok {
		return errors.New("target robot label should be added with key " + internal.TARGET_VDI_LABEL_KEY)
	}

	return nil
}
