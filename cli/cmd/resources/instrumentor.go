package resources

import (
	"fmt"

	"github.com/keyval-dev/odigos/cli/pkg/labels"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	instrumentorImage = "ghcr.io/keyval-dev/odigos/instrumentor"
	langDetectorImage = "ghcr.io/keyval-dev/odigos/lang-detector"
)

func NewInstrumentorServiceAccount() *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-instrumentor",
			Labels: labels.OdigosSystem,
		},
	}
}

func NewInstrumentorRoleBinding() *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-instrumentor-leader-election",
			Labels: labels.OdigosSystem,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind: "ServiceAccount",
				Name: "vision-instrumentor",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "odigos-leader-election-role",
		},
	}
}

func NewInstrumentorClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-instrumentor",
			Labels: labels.OdigosSystem,
		},
		Rules: []rbacv1.PolicyRule{
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{""},
				Resources: []string{
					"pods",
				},
			},
			{
				Verbs: []string{
					"list",
					"watch",
					"get",
				},
				APIGroups: []string{""},
				Resources: []string{
					"nodes",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{""},
				Resources: []string{
					"pods/status",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"deployments",
				},
			},
			{
				Verbs: []string{
					"update",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"deployments/finalizers",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"deployments/status",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"statefulsets",
				},
			},
			{
				Verbs: []string{
					"update",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"statefulsets/finalizers",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"statefulsets/status",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"collectorsgroups",
				},
			},
			{
				Verbs: []string{
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"collectorsgroups/finalizers",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"collectorsgroups/status",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"instrumentedapplications",
				},
			},
			{
				Verbs: []string{
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"instrumentedapplications/finalizers",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"instrumentedapplications/status",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"odigosconfigurations",
				},
			},
			{
				Verbs: []string{
					"create",
					"delete",
					"get",
					"list",
					"patch",
					"update",
					"watch",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"destinations",
				},
			},
			{
				Verbs: []string{
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"destinations/finalizers",
				},
			},
			{
				Verbs: []string{
					"get",
					"patch",
					"update",
				},
				APIGroups: []string{
					"vision.middleware.io",
				},
				Resources: []string{
					"destinations/status",
				},
			},
		},
	}
}

func NewInstrumentorClusterRoleBinding(ns string) *rbacv1.ClusterRoleBinding {
	return &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-instrumentor",
			Labels: labels.OdigosSystem,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "vision-instrumentor",
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "vision-instrumentor",
		},
	}
}

func NewInstrumentorDeployment(version string, telemetryEnabled bool, ignoredNamespaces []string) *appsv1.Deployment {
	args := []string{
		"--health-probe-bind-address=:8081",
		"--metrics-bind-address=127.0.0.1:8080",
		"--leader-elect",
		fmt.Sprintf("--lang-detector-tag=%s", version),
		fmt.Sprintf("--lang-detector-image=%s", langDetectorImage),
	}
	for _, v := range ignoredNamespaces {
		args = append(args, fmt.Sprintf("--ignore-namespace=%s", v))
	}

	if !telemetryEnabled {
		args = append(args, "--telemetry-disabled")
	}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "vision-instrumentor",
			Labels: map[string]string{
				"app":                       "vision-instrumentor",
				labels.OdigosSystemLabelKey: labels.OdigosSystemLabelValue,
			},
			Annotations: map[string]string{
				"vision.middleware.io/skip": "true",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrint32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "vision-instrumentor",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "vision-instrumentor",
					},
					Annotations: map[string]string{
						"kubectl.kubernetes.io/default-container": "manager",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "apis-rbac-proxy",
							Image: "gcr.io/kubebuilder/kube-rbac-proxy:v0.11.0",
							Args: []string{
								"--secure-listen-address=0.0.0.0:8443",
								"--upstream=http://127.0.0.1:8080/",
								"--logtostderr=true",
								"--v=0",
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "https",
									HostPort:      0,
									ContainerPort: 8443,
									Protocol:      "TCP",
								},
							},
							Resources: corev1.ResourceRequirements{},
						},
						{
							Name:  "manager",
							Image: fmt.Sprintf("%s:%s", instrumentorImage, version),
							Command: []string{
								"/app",
							},
							Args: args,
							Env: []corev1.EnvVar{
								{
									Name: "CURRENT_NS",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									"cpu":    resource.MustParse("500m"),
									"memory": *resource.NewQuantity(134217728, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									"cpu":    resource.MustParse("10m"),
									"memory": *resource.NewQuantity(67108864, resource.BinarySI),
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											Type:   intstr.Type(0),
											IntVal: 8081,
										},
									},
								},
								InitialDelaySeconds: 15,
								TimeoutSeconds:      0,
								PeriodSeconds:       20,
								SuccessThreshold:    0,
								FailureThreshold:    0,
							},
							SecurityContext: &corev1.SecurityContext{},
						},
					},
					TerminationGracePeriodSeconds: ptrint64(10),
					ServiceAccountName:            "vision-instrumentor",
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: ptrbool(true),
					},
				},
			},
			Strategy:        appsv1.DeploymentStrategy{},
			MinReadySeconds: 0,
		},
	}
}

func ptrint32(i int32) *int32 {
	return &i
}

func ptrint64(i int64) *int64 {
	return &i
}

func ptrbool(b bool) *bool {
	return &b
}
