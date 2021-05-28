/*
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

package resource

import (
	"context"
	"fmt"

	"github.com/prateekgogia/kit/pkg/apis/infrastructure/v1alpha1"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VPC struct {
	KubeClient client.Client
}

func (v *VPC) Create(ctx context.Context, controlPlane *v1alpha1.ControlPlane) error {
	if err := v.exists(ctx, controlPlane.Namespace, controlPlane.Name); err != nil {
		if errors.IsNotFound(err) {
			if err := v.create(ctx, controlPlane); err != nil {
				return fmt.Errorf("creating kube object, %w", err)
			}
			return nil
		}
		return fmt.Errorf("getting VPC object, %w", err)
	}
	// TODO verify existing object matches the desired else update
	return nil
}

func (v *VPC) create(ctx context.Context, controlPlane *v1alpha1.ControlPlane) error {
	if err := v.KubeClient.Create(ctx, &v1alpha1.VPC{
		ObjectMeta: ObjectMeta(controlPlane),
		Spec:       v1alpha1.VPCSpec{},
	}); err != nil {
		return fmt.Errorf("creating kube object, %w", err)
	}
	zap.S().Debugf("Successfully created VPC object for cluster %v", controlPlane.Name)
	return nil
}

func (v *VPC) exists(ctx context.Context, ns, objName string) error {
	result := &v1alpha1.VPC{}
	if err := v.KubeClient.Get(ctx, NamespacedName(ns, objName), result); err != nil {
		return err
	}
	return nil
}
