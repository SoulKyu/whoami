package models

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

type Users struct {
	Username    string `json:"username"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Password    string `json:"password,omitempty"`
	Token       string `json:"token,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}
type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}

type Page struct {
	ID      int
	Title   string
	Content string
}

type Application struct {
	Name      string
	Namespace corev1.Namespace
	Pods      []corev1.Pod
	Ingress   []networkingv1.Ingress
}

type NamespaceResources struct {
	Namespace    corev1.Namespace
	Pods         []corev1.Pod
	Ingresses    []networkingv1.Ingress
	Services     []corev1.Service
	Deployments  []appsv1.Deployment
	StatefulSets []appsv1.StatefulSet
	DaemonSets   []appsv1.DaemonSet
	Endpoints    []corev1.Endpoints
}
