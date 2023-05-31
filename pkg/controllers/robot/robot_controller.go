package robot

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/go-logr/logr"
	"github.com/robolaunch/devspace-operator/internal/label"
	robotv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
)

// RobotReconciler reconciles a Robot object
type RobotReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots/finalizers,verbs=update

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=workspacemanagers,verbs=get;list;watch;create;update;patch;delete

var logger logr.Logger

func (r *RobotReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = log.FromContext(ctx)

	instance, err := r.reconcileGetInstance(ctx, req.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	_, err = r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	// err = r.reconcileCheckDeletion(ctx, instance)
	// if err != nil {

	// 	if errors.IsNotFound(err) {
	// 		return ctrl.Result{}, nil
	// 	}

	// 	return ctrl.Result{}, err
	// }

	err = r.reconcileCheckImage(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckResources(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *RobotReconciler) reconcileCheckStatus(ctx context.Context, instance *robotv1alpha1.Robot) error {
	switch instance.Status.VolumeStatuses.Var.Created &&
		instance.Status.VolumeStatuses.Opt.Created &&
		instance.Status.VolumeStatuses.Etc.Created &&
		instance.Status.VolumeStatuses.Usr.Created &&
		instance.Status.VolumeStatuses.Workspace.Created {
	case true:

		switch instance.Status.LoaderJobStatus.Created {
		case true:

			switch instance.Status.LoaderJobStatus.Phase {
			case string(robotv1alpha1.JobSucceeded):

				switch instance.Spec.DevSuiteTemplate.IDEEnabled || instance.Spec.DevSuiteTemplate.VDIEnabled {
				case true:

					switch instance.Status.DevSuiteStatus.Resource.Created {
					case true:

						switch instance.Status.DevSuiteStatus.Status.Phase {
						case robotv1alpha1.DevSuitePhaseRunning:

							instance.Status.Phase = robotv1alpha1.RobotPhaseEnvironmentReady

						}

					case false:

						instance.Status.Phase = robotv1alpha1.RobotPhaseCreatingDevelopmentSuite
						err := r.createDevSuite(ctx, instance, instance.GetDevSuiteMetadata())
						if err != nil {
							return err
						}
						instance.Status.DevSuiteStatus.Resource.Created = true

					}

				case false:

					instance.Status.Phase = robotv1alpha1.RobotPhaseEnvironmentReady

				}

			case string(robotv1alpha1.JobActive):

				instance.Status.Phase = robotv1alpha1.RobotPhaseConfiguringEnvironment

			case string(robotv1alpha1.JobFailed):

				// TODO: add reason
				instance.Status.Phase = robotv1alpha1.RobotPhaseFailed

			}

		case false:

			instance.Status.Phase = robotv1alpha1.RobotPhaseConfiguringEnvironment
			err := r.createJob(ctx, instance, instance.GetLoaderJobMetadata())
			if err != nil {
				return err
			}
			instance.Status.LoaderJobStatus.Created = true
		}

	case false:

		instance.Status.Phase = robotv1alpha1.RobotPhaseCreatingEnvironment

		if !instance.Status.VolumeStatuses.Var.Created {
			err := r.createPVC(ctx, instance, instance.GetPVCVarMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatuses.Var.Created = true
		}

		if !instance.Status.VolumeStatuses.Opt.Created {
			err := r.createPVC(ctx, instance, instance.GetPVCOptMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatuses.Opt.Created = true
		}

		if !instance.Status.VolumeStatuses.Etc.Created {
			err := r.createPVC(ctx, instance, instance.GetPVCEtcMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatuses.Etc.Created = true
		}

		if !instance.Status.VolumeStatuses.Usr.Created {
			err := r.createPVC(ctx, instance, instance.GetPVCUsrMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatuses.Usr.Created = true
		}

		if !instance.Status.VolumeStatuses.Workspace.Created {
			err := r.createPVC(ctx, instance, instance.GetPVCWorkspaceMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatuses.Workspace.Created = true
		}
	}

	return nil
}

func (r *RobotReconciler) reconcileCheckResources(ctx context.Context, instance *robotv1alpha1.Robot) error {

	err := r.reconcileCheckPVCs(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckLoaderJob(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckDevSuite(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckWorkspaceManager(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RobotReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&robotv1alpha1.Robot{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Owns(&batchv1.Job{}).
		Owns(&robotv1alpha1.WorkspaceManager{}).
		Watches(
			&source.Kind{Type: &robotv1alpha1.DevSuite{}},
			handler.EnqueueRequestsFromMapFunc(r.watchAttachedDevSuites),
		).
		Complete(r)
}

func (r *RobotReconciler) watchAttachedDevSuites(o client.Object) []reconcile.Request {

	obj := o.(*robotv1alpha1.DevSuite)

	robot := &robotv1alpha1.Robot{}
	err := r.Get(context.TODO(), types.NamespacedName{
		Name:      label.GetTargetRobot(obj),
		Namespace: obj.Namespace,
	}, robot)
	if err != nil {
		return []reconcile.Request{}
	}

	return []reconcile.Request{
		{
			NamespacedName: types.NamespacedName{
				Name:      robot.Name,
				Namespace: robot.Namespace,
			},
		},
	}
}
