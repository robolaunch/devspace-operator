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
	RobotDevSuitePhaseRobotNotFound    RobotDevSuitePhase = "RobotNotFound"
	RobotDevSuitePhaseCreatingRobotVDI RobotDevSuitePhase = "CreatingRobotVDI"
	RobotDevSuitePhaseCreatingRobotIDE RobotDevSuitePhase = "CreatingRobotIDE"
	RobotDevSuitePhaseRunning          RobotDevSuitePhase = "Running"
	RobotDevSuitePhaseDeactivating     RobotDevSuitePhase = "Deactivating"
	RobotDevSuitePhaseInactive         RobotDevSuitePhase = "Inactive"
)

type RobotIDEPhase string

const (
	RobotIDEPhaseCreatingService RobotIDEPhase = "CreatingService"
	RobotIDEPhaseCreatingPod     RobotIDEPhase = "CreatingPod"
	RobotIDEPhaseCreatingIngress RobotIDEPhase = "CreatingIngress"
	RobotIDEPhaseRunning         RobotIDEPhase = "Running"
)

type RobotVDIPhase string

const (
	RobotVDIPhaseCreatingPVC        RobotVDIPhase = "CreatingPVC"
	RobotVDIPhaseCreatingTCPService RobotVDIPhase = "CreatingTCPService"
	RobotVDIPhaseCreatingUDPService RobotVDIPhase = "CreatingUDPService"
	RobotVDIPhaseCreatingPod        RobotVDIPhase = "CreatingPod"
	RobotVDIPhaseCreatingIngress    RobotVDIPhase = "CreatingIngress"
	RobotVDIPhaseRunning            RobotVDIPhase = "Running"
)
