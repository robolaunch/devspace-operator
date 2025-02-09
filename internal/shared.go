package internal

import corev1 "k8s.io/api/core/v1"

// Platform related labels
const (
	PLATFORM_VERSION_LABEL_KEY = "robolaunch.io/platform"
)

// Tenancy labels
const (
	ORGANIZATION_LABEL_KEY         = "robolaunch.io/organization"
	TEAM_LABEL_KEY                 = "robolaunch.io/team"
	REGION_LABEL_KEY               = "robolaunch.io/region"
	CLOUD_INSTANCE_LABEL_KEY       = "robolaunch.io/cloud-instance"
	CLOUD_INSTANCE_ALIAS_LABEL_KEY = "robolaunch.io/cloud-instance-alias"
	PHYSICAL_INSTANCE_LABEL_KEY    = "robolaunch.io/physical-instance"
)

// Ready devspace label
const (
	DEVSPACE_IMAGE_USER       = "robolaunch.io/devspace-image-user"
	DEVSPACE_IMAGE_REPOSITORY = "robolaunch.io/devspace-image-repository"
	DEVSPACE_IMAGE_TAG        = "robolaunch.io/devspace-image-tag"
)

// Target resource labels
const (
	TARGET_DEVSPACE_LABEL_KEY = "robolaunch.io/target-devspace"
	TARGET_VDI_LABEL_KEY      = "robolaunch.io/target-vdi"
)

// Special escape labels
const (
	DEVSPACE_DEV_SUITE_OWNED = "robolaunch.io/dev-suite-owned"
)

// DevSpace owned resources' postfixes
const (
	PVC_VAR_POSTFIX            = "-pvc-var"
	PVC_ETC_POSTFIX            = "-pvc-etc"
	PVC_OPT_POSTFIX            = "-pvc-opt"
	PVC_USR_POSTFIX            = "-pvc-usr"
	PVC_DISPLAY_POSTFIX        = "-pvc-display"
	PVC_WORKSPACE_POSTFIX      = "-pvc-workspace"
	DISCOVERY_SERVER_POSTFIX   = "-discovery"
	JOB_LOADER_POSTFIX         = "-loader"
	ROS_BRIDGE_POSTFIX         = "-bridge"
	DEVSPACE_DEV_SUITE_POSTFIX = "-dev"
	WORKSPACE_MANAGER_POSTFIX  = "-ws"
)

// WorkspaceManager owned resources' postfixes
const (
	JOB_CLONER_POSTFIX  = "-cloner"
	JOB_CLEANUP_POSTFIX = "-cleanup"
)

// DevSpaceVDI owned resources' postfixes
const (
	PVC_VDI_POSTFIX     = "-display"
	SVC_TCP_VDI_POSTFIX = "-tcp"
	SVC_UDP_VDI_POSTFIX = "-udp"
	POD_VDI_POSTFIX     = ""
	INGRESS_VDI_POSTFIX = ""
)

// DevSpaceIDE owned resources' postfixes
const (
	SVC_IDE_POSTFIX     = ""
	POD_IDE_POSTFIX     = ""
	INGRESS_IDE_POSTFIX = ""
)

// DevSpaceIDE owned resources' postfixes
const (
	SVC_JUPYTER_POSTFIX     = ""
	POD_JUPYTER_POSTFIX     = ""
	INGRESS_JUPYTER_POSTFIX = ""
)

// DevSuite owned resources' postfixes
const (
	DEVSPACE_VDI_POSTFIX     = "-vdi"
	DEVSPACE_IDE_POSTFIX     = "-ide"
	DEVSPACE_JUPYTER_POSTFIX = "-jupyter"
)

// Paths

const (
	CUSTOM_SCRIPTS_PATH = "/etc/custom"
	HELPERS_PATH        = "/var/lib/robolaunch-helpers/"
	X11_UNIX_PATH       = "/tmp/.X11-unix"
)

// Ingress annotations
const (
	INGRESS_AUTH_URL_KEY                  = "nginx.ingress.kubernetes.io/auth-url"
	INGRESS_AUTH_URL_VAL                  = "https://%s.%s/oauth2/auth"
	INGRESS_AUTH_SIGNIN_KEY               = "nginx.ingress.kubernetes.io/auth-signin"
	INGRESS_AUTH_SIGNIN_VAL               = "https://%s.%s/oauth2/start?rd=$scheme://$best_http_host$request_uri"
	INGRESS_AUTH_RESPONSE_HEADERS_KEY     = "nginx.ingress.kubernetes.io/auth-response-headers"
	INGRESS_AUTH_RESPONSE_HEADERS_VAL     = "x-auth-request-user, x-auth-request-email, x-auth-request-access-token"
	INGRESS_CONFIGURATION_SNIPPET_KEY     = "nginx.ingress.kubernetes.io/configuration-snippet"
	INGRESS_VDI_CONFIGURATION_SNIPPET_VAL = "" +
		"        #proxy_set_header Host $host;\n" +
		"		proxy_set_header X-Real-IP $remote_addr;\n" +
		"		proxy_set_header X-Forwarded-For $remote_addr;\n" +
		"		proxy_set_header X-Forwarded-Host $host;\n" +
		"		proxy_set_header X-Forwarded-Port $server_port;\n" +
		"		proxy_set_header X-Forwarded-Protocol $scheme;\n"
	INGRESS_CERT_MANAGER_KEY                   = "acme.cert-manager.io/http01-edit-in-place"
	INGRESS_CERT_MANAGER_VAL                   = "true"
	INGRESS_NGINX_PROXY_BUFFER_SIZE_KEY        = "nginx.ingress.kubernetes.io/proxy-buffer-size"
	INGRESS_NGINX_PROXY_BUFFER_SIZE_VAL        = "16k"
	INGRESS_NGINX_PROXY_BUFFERS_NUMBER_KEY     = "nginx.ingress.kubernetes.io/proxy-buffers-number"
	INGRESS_VDI_NGINX_PROXY_BUFFERS_NUMBER_VAL = "4"
	INGRESS_NGINX_REWRITE_TARGET_KEY           = "nginx.ingress.kubernetes.io/rewrite-target"
	INGRESS_NGINX_REWRITE_TARGET_VAL           = "/$2"

	INGRESS_IDE_CONFIGURATION_SNIPPET_VAL = "" +
		"auth_request_set $name_upstream_1 $upstream_cookie_name_1;" +
		"access_by_lua_block {" +
		"  if ngx.var.name_upstream_1 ~= \"\" then" +
		"	ngx.header[\"Set-Cookie\"] = \"name_1=\" .. ngx.var.name_upstream_1 .. ngx.var.auth_cookie:match(\"(; .*)\")" +
		"  end" +
		"}"
)

// Commands for collecting metrics
const (
	CMD_GET_CPU          = "cat /sys/fs/cgroup/cpu/cpuacct.usage"
	CMD_GET_MEMORY       = "cat /sys/fs/cgroup/memory/memory.usage_in_bytes"
	CMD_GET_NETWORK_LOAD = "cat /proc/net/dev | awk -F ' ' '{print $1 $2 \":\" $10}' | tail -n+3"
)

func Bash(command string) []string {
	return []string{
		"/bin/bash",
		"-c",
		command,
	}
}

func BashWithUser(command, user string) []string {
	return []string{
		"sudo",
		"-H",
		"-u",
		user,
		"bash",
		"-c",
		command,
	}
}

func Env(key string, value string) corev1.EnvVar {
	return corev1.EnvVar{
		Name:  key,
		Value: value,
	}
}
