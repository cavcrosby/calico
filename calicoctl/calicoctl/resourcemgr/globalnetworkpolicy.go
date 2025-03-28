// Copyright (c) 2017-2024 Tigera, Inc. All rights reserved.

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

package resourcemgr

import (
	"context"
	"strings"

	api "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	client "github.com/projectcalico/calico/libcalico-go/lib/clientv3"
	cerrors "github.com/projectcalico/calico/libcalico-go/lib/errors"
	"github.com/projectcalico/calico/libcalico-go/lib/names"
	"github.com/projectcalico/calico/libcalico-go/lib/options"
)

func init() {
	registerResource(
		api.NewGlobalNetworkPolicy(),
		newGlobalNetworkPolicyList(),
		false,
		[]string{"globalnetworkpolicy", "globalnetworkpolicies", "gnp", "gnps"},
		[]string{"NAME", "TIER"},
		[]string{"NAME", "TIER", "ORDER", "SELECTOR"},
		map[string]string{
			"NAME":     "{{.ObjectMeta.Name}}",
			"ORDER":    "{{.Spec.Order}}",
			"SELECTOR": "{{.Spec.Selector}}",
			"TIER":     "{{.Spec.Tier}}",
		},
		func(ctx context.Context, client client.Interface, resource ResourceObject) (ResourceObject, error) {
			r := resource.(*api.GlobalNetworkPolicy)
			if policyIsANP(r) {
				return nil, cerrors.ErrorOperationNotSupported{
					Operation:  "create or apply",
					Identifier: resource,
					Reason:     "kubernetes admin network policies must be managed through the kubernetes API",
				}
			}
			return client.GlobalNetworkPolicies().Create(ctx, r, options.SetOptions{})
		},
		func(ctx context.Context, client client.Interface, resource ResourceObject) (ResourceObject, error) {
			r := resource.(*api.GlobalNetworkPolicy)
			if policyIsANP(r) {
				return nil, cerrors.ErrorOperationNotSupported{
					Operation:  "create or apply",
					Identifier: resource,
					Reason:     "kubernetes admin network policies must be managed through the kubernetes API",
				}
			}
			return client.GlobalNetworkPolicies().Update(ctx, r, options.SetOptions{})
		},
		func(ctx context.Context, client client.Interface, resource ResourceObject) (ResourceObject, error) {
			r := resource.(*api.GlobalNetworkPolicy)
			if policyIsANP(r) {
				return nil, cerrors.ErrorOperationNotSupported{
					Operation:  "create or apply",
					Identifier: resource,
					Reason:     "kubernetes admin network policies must be managed through the kubernetes API",
				}
			}
			return client.GlobalNetworkPolicies().Delete(ctx, r.Name, options.DeleteOptions{ResourceVersion: r.ResourceVersion})
		},
		func(ctx context.Context, client client.Interface, resource ResourceObject) (ResourceObject, error) {
			r := resource.(*api.GlobalNetworkPolicy)
			return client.GlobalNetworkPolicies().Get(ctx, r.Name, options.GetOptions{ResourceVersion: r.ResourceVersion})
		},
		func(ctx context.Context, client client.Interface, resource ResourceObject) (ResourceListObject, error) {
			r := resource.(*api.GlobalNetworkPolicy)
			return client.GlobalNetworkPolicies().List(ctx, options.ListOptions{ResourceVersion: r.ResourceVersion, Name: r.Name})
		},
	)
}

func policyIsANP(r *api.GlobalNetworkPolicy) bool {
	return strings.HasPrefix(r.Name, names.K8sAdminNetworkPolicyNamePrefix) ||
		strings.HasPrefix(r.Name, names.K8sBaselineAdminNetworkPolicyNamePrefix)
}

// newGlobalNetworkPolicyList creates a new (zeroed) GlobalNetworkPolicyList struct with the TypeMetadata initialised to the current
// version.
func newGlobalNetworkPolicyList() *api.GlobalNetworkPolicyList {
	return &api.GlobalNetworkPolicyList{
		TypeMeta: metav1.TypeMeta{
			Kind:       api.KindGlobalNetworkPolicyList,
			APIVersion: api.GroupVersionCurrent,
		},
	}
}
