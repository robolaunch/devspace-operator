package v1alpha1

type RobotPhase string

const (
	RobotPhaseCreatingEnvironment      RobotPhase = "CreatingEnvironment"
	RobotPhaseConfiguringEnvironment   RobotPhase = "ConfiguringEnvironment"
	RobotPhaseCreatingDevelopmentSuite RobotPhase = "CreatingDevelopmentSuite"
	RobotPhaseConfiguringWorkspaces    RobotPhase = "ConfiguringWorkspaces"
	RobotPhaseEnvironmentReady         RobotPhase = "EnvironmentReady"
	RobotPhaseBuilding                 RobotPhase = "Building"
	RobotPhaseBuilt                    RobotPhase = "Built"
	RobotPhaseLaunching                RobotPhase = "Launching"
	RobotPhaseRunning                  RobotPhase = "Running"
	RobotPhaseDeletingLoaderJob        RobotPhase = "DeletingLoaderJob"
	RobotPhaseDeletingVolumes          RobotPhase = "DeletingVolumes"

	RobotPhaseFailed RobotPhase = "Failed"
)

type WorkspaceManagerPhase string

const (
	WorkspaceManagerPhaseConfiguringWorkspaces WorkspaceManagerPhase = "ConfiguringWorkspaces"
	WorkspaceManagerPhaseReady                 WorkspaceManagerPhase = "Ready"
	WorkspaceManagerPhaseFailed                WorkspaceManagerPhase = "Failed"
)

type RobotDevSuitePhase string

const (
	RobotDevSuitePhaseRobotNotFound       RobotDevSuitePhase = "RobotNotFound"
	RobotDevSuitePhaseCreatingDevSpaceVDI RobotDevSuitePhase = "CreatingDevSpaceVDI"
	RobotDevSuitePhaseCreatingDevSpaceIDE RobotDevSuitePhase = "CreatingDevspaceIDE"
	RobotDevSuitePhaseRunning             RobotDevSuitePhase = "Running"
	RobotDevSuitePhaseDeactivating        RobotDevSuitePhase = "Deactivating"
	RobotDevSuitePhaseInactive            RobotDevSuitePhase = "Inactive"
)

type DevSpaceIDEPhase string

const (
	DevSpaceIDEPhaseCreatingService DevSpaceIDEPhase = "CreatingService"
	DevSpaceIDEPhaseCreatingPod     DevSpaceIDEPhase = "CreatingPod"
	DevSpaceIDEPhaseCreatingIngress DevSpaceIDEPhase = "CreatingIngress"
	DevSpaceIDEPhaseRunning         DevSpaceIDEPhase = "Running"
)

type DevSpaceVDIPhase string

const (
	DevSpaceVDIPhaseCreatingPVC        DevSpaceVDIPhase = "CreatingPVC"
	DevSpaceVDIPhaseCreatingTCPService DevSpaceVDIPhase = "CreatingTCPService"
	DevSpaceVDIPhaseCreatingUDPService DevSpaceVDIPhase = "CreatingUDPService"
	DevSpaceVDIPhaseCreatingPod        DevSpaceVDIPhase = "CreatingPod"
	DevSpaceVDIPhaseCreatingIngress    DevSpaceVDIPhase = "CreatingIngress"
	DevSpaceVDIPhaseRunning            DevSpaceVDIPhase = "Running"
)
