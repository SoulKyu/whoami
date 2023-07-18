package kubernetes

import (
	"context"
	"log"
	"whoami/models"
	"whoami/pkg/auth"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListNamespaces returns a list of namespaces in the current Kubernetes cluster.
func ListNamespaces() ([]corev1.Namespace, error) {
	// Create clientset from the config.
	clientset, err := auth.GetFakeOrLiveKubernetesClient()
	if err != nil {
		log.Printf("Get Kubernetes Client: %v", err)
		return nil, err
	}

	// Retrieve the list of namespaces.
	namespacesList, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Unable to get namespacesList from Kubernetes: %v", err)
		return nil, err
	}

	return namespacesList.Items, nil
}

// ListPodsInNamespace returns a list of pods in the specified Kubernetes namespace.
func ListPodsInNamespace(namespace string) ([]string, error) {

	// Create clientset from the config.
	clientset, err := auth.GetFakeOrLiveKubernetesClient()
	if err != nil {
		log.Printf("Get Kubernetes Client: %v", err)
		return nil, err
	}

	// Retrieve the list of pods in the specified namespace.
	podsList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Unable to get podsList from Kubernetes: %v", err)
		return nil, err
	}

	var pods []string

	for _, pod := range podsList.Items {
		pods = append(pods, pod.Name)
	}

	return pods, nil
}

// ListIngressInNamespace returns a list of ingress in the specified Kubernetes namespace.
func ListIngressInNamespace(namespace string) ([]string, error) {
	// Create clientset from the config.
	clientset, err := auth.GetFakeOrLiveKubernetesClient()
	if err != nil {
		log.Printf("Get Kubernetes Client: %v", err)
		return nil, err
	}

	// Retrieve the list of ingresses in the specified namespace.
	ingressList, err := clientset.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Unable to get ingressList from Kubernetes: %v", err)
		return nil, err
	}
	var ingresses []string

	for _, ingress := range ingressList.Items {
		ingresses = append(ingresses, ingress.Spec.Rules[0].Host)
	}

	return ingresses, nil
}

// ListServicesInNamespace returns a list of services in the specified namespace.
func ListServicesInNamespace(namespace string) ([]string, error) {
	clientset, err := auth.GetFakeOrLiveKubernetesClient()
	if err != nil {
		log.Printf("Get Kubernetes Client: %v", err)
		return nil, err
	}

	servicesList, err := clientset.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Unable to get services from Kubernetes: %v", err)
		return nil, err
	}

	var services []string

	for _, service := range servicesList.Items {
		services = append(services, service.Name)
	}

	return services, nil
}

func GetNamespaceResources() ([]models.NamespaceResources, error) {

	namespaces, err := ListNamespaces()
	if err != nil {
		log.Printf("Unable to get Namespace with error : ", err)
		return nil, err
	}

	var namespaceResources []models.NamespaceResources
	for _, namespace := range namespaces {
		pods, err := ListPodsInNamespace(namespace.Name)
		if err != nil {
			log.Printf("Unable to get pods with error : %v", err)
			return nil, err
		}
		ingresses, err := ListIngressInNamespace(namespace.Name)
		if err != nil {
			log.Printf("Unable to get ingresses with error : %v", err)
			return nil, err
		}
		services, err := ListServicesInNamespace(namespace.Name)
		if err != nil {
			log.Printf("Unable to get services with error : %v", err)
			return nil, err
		}
		namespaceResources = append(namespaceResources, models.NamespaceResources{
			Namespace: namespace.Name,
			Pods:      pods,
			Ingresses: ingresses,
			Services:  services,
		})
	}
	return namespaceResources, nil
}
