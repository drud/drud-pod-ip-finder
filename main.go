package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
)

var kubeClient *client.Client

// get list of pods that matches the passed in lavels
func getPodsByLabel(kubeClient *client.Client, matches map[string]string) ([]api.Pod, error) {
	namespace := api.NamespaceAll

	if _, ok := matches["namespace"]; ok {
		namespace = matches["namespace"]
		delete(matches, "namespace")
	}

	selector := labels.SelectorFromSet(labels.Set(matches))
	listMatches := api.ListOptions{LabelSelector: selector}

	pods, err := kubeClient.Pods(namespace).List(listMatches)
	if err != nil {
		return []api.Pod{}, err
	}

	return pods.Items, nil

}

// getPodIps gets a list of pods matches the passed in params and returns a list of their ips
func getPodIps(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	matches := map[string]string{}
	for key, val := range params {
		matches[key] = val[0]
	}

	pods, err := getPodsByLabel(kubeClient, matches)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var ips []string
	for _, pod := range pods {
		if pod.Status.Phase == "Running" {
			ips = append(ips, pod.Status.PodIP)
		}
	}

	b, err := json.Marshal(ips)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b)) // send data to client side
}

func main() {
	var err error
	kubeClient, err = client.NewInCluster()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	http.HandleFunc("/", getPodIps)         // set router
	err = http.ListenAndServe(":3000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
