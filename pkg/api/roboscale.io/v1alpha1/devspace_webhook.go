package v1alpha1

import (
	"errors"

	"github.com/robolaunch/devspace-operator/internal"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var devspacelog = logf.Log.WithName("devspace-resource")

func (r *DevSpace) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-dev-roboscale-io-v1alpha1-devspace,mutating=true,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspaces,verbs=create;update,versions=v1alpha1,name=mdevspace.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &DevSpace{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *DevSpace) Default() {
	devspacelog.Info("default", "name", r.Name)

	DefaultRepositoryPaths(r)
	_ = r.setRepositoryInfo()
}

func DefaultRepositoryPaths(r *DevSpace) {
	for wsKey := range r.Spec.WorkspaceManagerTemplate.Workspaces {
		ws := r.Spec.WorkspaceManagerTemplate.Workspaces[wsKey]
		for repoKey := range ws.Repositories {
			repo := ws.Repositories[repoKey]
			repo.Path = r.Spec.WorkspaceManagerTemplate.WorkspacesPath + "/" + ws.Name + "/src/" + repoKey
			ws.Repositories[repoKey] = repo
		}
		r.Spec.WorkspaceManagerTemplate.Workspaces[wsKey] = ws
	}
}

//+kubebuilder:webhook:path=/validate-dev-roboscale-io-v1alpha1-devspace,mutating=false,failurePolicy=fail,sideEffects=None,groups=dev.roboscale.io,resources=devspaces,verbs=create;update,versions=v1alpha1,name=vdevspace.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &DevSpace{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpace) ValidateCreate() error {
	devspacelog.Info("validate create", "name", r.Name)

	err := r.checkTenancyLabels()
	if err != nil {
		return err
	}

	err = r.checkDevSuite()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpace) ValidateUpdate(old runtime.Object) error {
	devspacelog.Info("validate update", "name", r.Name)

	err := r.checkTenancyLabels()
	if err != nil {
		return err
	}

	err = r.checkDevSuite()
	if err != nil {
		return err
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *DevSpace) ValidateDelete() error {
	devspacelog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *DevSpace) checkTenancyLabels() error {
	labels := r.GetLabels()

	if _, ok := labels[internal.ORGANIZATION_LABEL_KEY]; !ok {
		return errors.New("organization label should be added with key " + internal.ORGANIZATION_LABEL_KEY)
	}

	if _, ok := labels[internal.TEAM_LABEL_KEY]; !ok {
		return errors.New("team label should be added with key " + internal.TEAM_LABEL_KEY)
	}

	if _, ok := labels[internal.REGION_LABEL_KEY]; !ok {
		return errors.New("super cluster label should be added with key " + internal.REGION_LABEL_KEY)
	}

	if _, ok := labels[internal.CLOUD_INSTANCE_LABEL_KEY]; !ok {
		return errors.New("cloud instance label should be added with key " + internal.CLOUD_INSTANCE_LABEL_KEY)
	}

	if _, ok := labels[internal.CLOUD_INSTANCE_ALIAS_LABEL_KEY]; !ok {
		return errors.New("cloud instance alias label should be added with key " + internal.CLOUD_INSTANCE_ALIAS_LABEL_KEY)
	}
	return nil
}

func (r *DevSpace) checkDevSuite() error {

	dst := r.Spec.DevSuiteTemplate

	if dst.IDEEnabled && dst.DevSpaceIDETemplate.Display && !dst.VDIEnabled {
		return errors.New("cannot open an ide with a display when vdi disabled")
	}

	return nil
}

func (r *DevSpace) setRepositoryInfo() error {

	for k1, ws := range r.Spec.WorkspaceManagerTemplate.Workspaces {
		for k2, repo := range ws.Repositories {
			owner, repoName, err := getPathVariables(repo.URL)
			if err != nil {
				return err
			}

			repo.Owner = owner
			repo.Repo = repoName

			lastCommitHash, err := getLastCommitHash(repo)
			if err != nil {
				return err
			}

			repo.Hash = lastCommitHash

			ws.Repositories[k2] = repo
		}
		r.Spec.WorkspaceManagerTemplate.Workspaces[k1] = ws
	}

	return nil

}
