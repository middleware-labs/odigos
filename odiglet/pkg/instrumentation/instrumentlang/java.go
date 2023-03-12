package instrumentlang

import (
	"fmt"
	"github.com/keyval-dev/odigos/visioncart/pkg/env"
	"github.com/keyval-dev/odigos/visioncart/pkg/instrumentation/consts"
	"k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	otelResourceAttributesEnvVar = "OTEL_RESOURCE_ATTRIBUTES"
	otelResourceAttrPatteern     = "service.name=%s,odigos.device=java"
	javaToolOptionsEnvVar        = "JAVA_TOOL_OPTIONS"
	javaOptsEnvVar               = "JAVA_OPTS"
	javaToolOptionsPattern       = "-javaagent:/odigos/java/javaagent.jar " +
		"-Dotel.traces.sampler=always_on -Dotel.exporter.otlp.endpoint=http://%s:%d"
)

func Java(deviceId string) *v1beta1.ContainerAllocateResponse {
	javaOpts := fmt.Sprintf(javaToolOptionsPattern, env.Current.NodeIP, consts.OTLPPort)
	return &v1beta1.ContainerAllocateResponse{
		Envs: map[string]string{
			otelResourceAttributesEnvVar: fmt.Sprintf(otelResourceAttrPatteern, deviceId),
			javaToolOptionsEnvVar:        javaOpts,
			javaOptsEnvVar:               javaOpts,
		},
		Mounts: []*v1beta1.Mount{
			{
				ContainerPath: "/odigos/java",
				HostPath:      "/odigos/java",
				ReadOnly:      true,
			},
		},
	}
}
