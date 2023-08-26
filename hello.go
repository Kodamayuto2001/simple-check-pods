package main

import (
	"context"
	"fmt"

	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	//  kubeconfigのパスを取得。デフォルトは~/.kube/config
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	//  kubeconfigを使用してconfigをビルド
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//  Kubernetesクライアントを作成
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//  Podのりすとを取得
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pods found in namespace %v:\n", "default")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
