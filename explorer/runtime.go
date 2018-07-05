// Copyright © 2018 Kris Nova <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package explorer

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Required to authenticate against GKE clusters.
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	k8sclient *kubernetes.Clientset
	options   *RuntimeOptions
)

type RuntimeOptions struct {
	KubeconfigPath string
	ShellExecImage string
}

func Init(opt *RuntimeOptions) error {

	s := `  _
 | |
 | | ___ __ _____  __
 | |/ / '__/ _ \ \/ /
 |   <| | |  __/>  <
 |_|\_\_|  \___/_/\_\`

	fmt.Println(s)
	fmt.Println()
	fmt.Println("Kubernetes Resource Explorer")
	fmt.Println()

	config, err := clientcmd.BuildConfigFromFlags("", opt.KubeconfigPath)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	k8sclient = clientset
	options = opt
	return nil
}
