package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "namespace", "name", "toDir")
		return
	}
	namespace := os.Args[1]
	name := os.Args[2]
	toDir := os.Args[3]

	config, err := rest.InClusterConfig()
	checkError(err)

	clientset, err := kubernetes.NewForConfig(config)
	checkError(err)

	configMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	checkError(err)

	os.MkdirAll(toDir, 0644)
	checkError(err)

	for key, value := range configMap.Data {
		f, err := os.Create(filepath.Join(toDir, key))
		checkError(err)
		defer f.Close()

		f.WriteString(value)
		f.Sync()
	}

	for key, value := range configMap.BinaryData {
		f, err := os.Create(filepath.Join(toDir, key))
		checkError(err)
		defer f.Close()

		f.Write(value)
		f.Sync()
	}
}
