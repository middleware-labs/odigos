package resources

import (
	"github.com/keyval-dev/odigos/cli/pkg/labels"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewDataCollectionServiceAccount() *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-data-collection",
			Labels: labels.OdigosSystem,
		},
	}
}

func NewDataCollectionClusterRole() *rbacv1.ClusterRole {
	return &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-data-collection",
			Labels: labels.OdigosSystem,
		},
		Rules: []rbacv1.PolicyRule{
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{""},
				Resources: []string{
					"nodes",
					"nodes/stats",
					"namespaces",
					"pods",
					"serviceaccounts",
					"services",
					"configmaps",
					"endpoints",
					"persistentvolumeclaims",
					"replicationcontrollers",
					"replicationcontrollers/scale",
					"persistentvolumeclaims",
					"persistentvolumes",
					"bindings",
					"events",
					"limitranges",
					"namespaces/status",
					"pods/log",
					"pods/status",
					"replicationcontrollers/status",
					"resourcequotas",
					"resourcequotas/status",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"metrics.k8s.io"},
				Resources: []string{
					"nodes",
					"pods",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
					"patch",
				},
				APIGroups: []string{"apps"},
				Resources: []string{
					"daemonsets",
					"deployments",
					"deployments/scale",
					"replicasets",
					"replicasets/scale",
					"statefulsets",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"autoscaling"},
				Resources: []string{
					"horizontalpodautoscalers",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"batch"},
				Resources: []string{
					"cronjobs",
					"jobs",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"extensions"},
				Resources: []string{
					"daemonsets",
					"deployments",
					"deployments/scale",
					"networkpolicies",
					"replicasets",
					"replicasets/scale",
					"replicationcontrollers/scale",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"networking.k8s.io"},
				Resources: []string{
					"ingresses",
					"networkpolicies",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"policy"},
				Resources: []string{
					"poddisruptionbudgets",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"storage.k8s.io"},
				Resources: []string{
					"storageclasses",
					"volumeattachments",
				},
			},
			{
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
				APIGroups: []string{"rbac.authorization.k8s.io"},
				Resources: []string{
					"clusterrolebindings",
					"clusterroles",
					"roles",
					"rolebindings",
				},
			},
		},
	}
}

func NewDataCollectionClusterRoleBinding(ns string) *rbacv1.ClusterRoleBinding {
	return &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "vision-data-collection",
			Labels: labels.OdigosSystem,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "vision-data-collection",
				Namespace: ns,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "vision-data-collection",
		},
	}
}
