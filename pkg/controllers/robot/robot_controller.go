package robot

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
)

// RobotReconciler reconciles a Robot object
type RobotReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=robots/finalizers,verbs=update

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=discoveryservers,verbs=get;list;watch;create;update;patch;delete

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
	switch instance.Status.VolumeStatus.Var &&
		instance.Status.VolumeStatus.Opt &&
		instance.Status.VolumeStatus.Etc &&
		instance.Status.VolumeStatus.Usr &&
		instance.Status.VolumeStatus.Display &&
		instance.Status.VolumeStatus.Workspace {
	case true:

		switch instance.Spec.DiscoveryServerTemplate.Attached {
		case true:

			switch instance.Status.DiscoveryServerStatus.Created {
			case true:

				switch instance.Status.DiscoveryServerStatus.Status.Phase {
				case robotv1alpha1.DiscoveryServerPhaseReady:

					instance.Status.Phase = robotv1alpha1.RobotPhaseConfiguringEnvironment

				}

			case false:

				// create discovery server

				instance.Status.Phase = robotv1alpha1.RobotPhaseCreatingDiscoveryServer
				err := r.createDiscoveryServer(ctx, instance, instance.GetDiscoveryServerMetadata())
				if err != nil {
					return err
				}
				instance.Status.DiscoveryServerStatus.Created = true

			}

		case false:

			instance.Status.Phase = robotv1alpha1.RobotPhaseConfiguringEnvironment

		}

	case false:

		instance.Status.Phase = robotv1alpha1.RobotPhaseCreatingEnvironment

		if !instance.Status.VolumeStatus.Var {
			err := r.createPVC(ctx, instance, instance.GetPVCVarMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Var = true
		}

		if !instance.Status.VolumeStatus.Opt {
			err := r.createPVC(ctx, instance, instance.GetPVCOptMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Opt = true
		}

		if !instance.Status.VolumeStatus.Etc {
			err := r.createPVC(ctx, instance, instance.GetPVCEtcMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Etc = true
		}

		if !instance.Status.VolumeStatus.Usr {
			err := r.createPVC(ctx, instance, instance.GetPVCUsrMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Usr = true
		}

		if !instance.Status.VolumeStatus.Display {
			err := r.createPVC(ctx, instance, instance.GetPVCDisplayMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Display = true
		}

		if !instance.Status.VolumeStatus.Workspace {
			err := r.createPVC(ctx, instance, instance.GetPVCWorkspaceMetadata())
			if err != nil {
				return err
			}
			instance.Status.VolumeStatus.Workspace = true
		}
	}

	return nil
}

func (r *RobotReconciler) reconcileCheckResources(ctx context.Context, instance *robotv1alpha1.Robot) error {

	err := r.reconcileCheckPVCs(ctx, instance)
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
		Owns(&robotv1alpha1.DiscoveryServer{}).
		Complete(r)
}
