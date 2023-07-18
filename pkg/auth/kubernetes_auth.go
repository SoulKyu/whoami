/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package auth

import (
	"context"
	"os"
	"log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes/fake"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedbatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	typedbatchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	typednetworkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

type K8sClientSetInterface interface {
    // Here you can list the methods you need to use on the clientsets
    // For instance:
	CoreV1() typedcorev1.CoreV1Interface
	AppsV1() typedappsv1.AppsV1Interface
	BatchV1() typedbatchv1.BatchV1Interface
	BatchV1beta1() typedbatchv1beta1.BatchV1beta1Interface
	NetworkingV1() typednetworkingv1.NetworkingV1Interface
}

func GetInClusterConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetKubernetesClient() (K8sClientSetInterface, error) {
	config, err := GetInClusterConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func GetFakeKubernetesClient() (K8sClientSetInterface, error) {

	clientset := fake.NewSimpleClientset()

	return clientset, nil
}

func CreateFakeKubernetesApplication(client K8sClientSetInterface) error {

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-namespace",
		},
	}
	_, _ = client.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: ns.Name,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "my-container",
					Image: "nginx:latest",
				},
			},
		},
	}
	_, err := client.CoreV1().Pods(ns.Name).Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Une erreur est survenu lors du fake création du pod : %v", err)
		return err
	}


	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-service",
			Namespace: ns.Name,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "my-app",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port:     80,
				},
			},
		},
	}
	_, _ = client.CoreV1().Services(ns.Name).Create(context.Background(), svc, metav1.CreateOptions{})

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-configmap",
			Namespace: ns.Name,
		},
		Data: map[string]string{
			"key": "value",
		},
	}
	_, _ = client.CoreV1().ConfigMaps(ns.Name).Create(context.Background(), cm, metav1.CreateOptions{})

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-deployment",
			Namespace: ns.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "my-container",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
	_, _ = client.AppsV1().Deployments(ns.Name).Create(context.Background(), deploy, metav1.CreateOptions{})

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-job",
			Namespace: ns.Name,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "job-container",
							Image: "busybox",
							Command: []string{
								"/bin/sh",
								"-c",
								"date; echo Hello from the Kubernetes cluster",
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}
	_, _ = client.BatchV1().Jobs(ns.Name).Create(context.Background(), job, metav1.CreateOptions{})

	cronjob := &batchv1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-cronjob",
			Namespace: ns.Name,
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: "*/1 * * * *",
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "cronjob-container",
									Image: "busybox",
									Command: []string{
										"/bin/sh",
										"-c",
										"date; echo Hello from the Kubernetes cron job",
									},
								},
							},
							RestartPolicy: corev1.RestartPolicyOnFailure,
						},
					},
				},
			},
		},
	}
	_, _ = client.BatchV1beta1().CronJobs(ns.Name).Create(context.Background(), cronjob, metav1.CreateOptions{})

	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-ingress",
			Namespace: ns.Name,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "my-app.my-domain.com",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path: "/",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "my-service",
											Port: networkingv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_, _ = client.NetworkingV1().Ingresses(ns.Name).Create(context.Background(), ingress, metav1.CreateOptions{})

	
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: ns.Name,
		},
		Data: map[string][]byte{
			"username": []byte("my-username"),
			"password": []byte("my-password"),
		},
	}
	_, _ = client.CoreV1().Secrets(ns.Name).Create(context.Background(), secret, metav1.CreateOptions{})

	statefulset := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-statefulset",
			Namespace: ns.Name,
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "my-container",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
	_, _ = client.AppsV1().StatefulSets(ns.Name).Create(context.Background(), statefulset, metav1.CreateOptions{})

	daemonset := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-daemonset",
			Namespace: ns.Name,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "my-container",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
	_, _ = client.AppsV1().DaemonSets(ns.Name).Create(context.Background(), daemonset, metav1.CreateOptions{})

	return nil
}

func GetFakeOrLiveKubernetesClient() (K8sClientSetInterface, error) {

	_, exists := os.LookupEnv("KUBERNETES_SERVICE_HOST")
	if exists {
		log.Println("Running inside a Kubernetes cluster")
		client, err := GetKubernetesClient()
		if err != nil {
			log.Printf("Get Kubernetes Client: %v", err)
			return nil, err
		}
		return client, nil
	} else {
		log.Println("Not running inside a Kubernetes cluster, initialize Fake application")
		client, err := GetFakeKubernetesClient()
		if err != nil {
			return nil, err
		}
		CreateFakeKubernetesApplication(client)
		return client, nil
	}
	return nil, nil
}