/*
Copyright k0s authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/k0sproject/k0s/pkg/apis/etcd/v1beta1"
	scheme "github.com/k0sproject/k0s/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// EtcdMembersGetter has a method to return a EtcdMemberInterface.
// A group's client should implement this interface.
type EtcdMembersGetter interface {
	EtcdMembers() EtcdMemberInterface
}

// EtcdMemberInterface has methods to work with EtcdMember resources.
type EtcdMemberInterface interface {
	Create(ctx context.Context, etcdMember *v1beta1.EtcdMember, opts v1.CreateOptions) (*v1beta1.EtcdMember, error)
	Update(ctx context.Context, etcdMember *v1beta1.EtcdMember, opts v1.UpdateOptions) (*v1beta1.EtcdMember, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, etcdMember *v1beta1.EtcdMember, opts v1.UpdateOptions) (*v1beta1.EtcdMember, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.EtcdMember, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.EtcdMemberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.EtcdMember, err error)
	EtcdMemberExpansion
}

// etcdMembers implements EtcdMemberInterface
type etcdMembers struct {
	*gentype.ClientWithList[*v1beta1.EtcdMember, *v1beta1.EtcdMemberList]
}

// newEtcdMembers returns a EtcdMembers
func newEtcdMembers(c *EtcdV1beta1Client) *etcdMembers {
	return &etcdMembers{
		gentype.NewClientWithList[*v1beta1.EtcdMember, *v1beta1.EtcdMemberList](
			"etcdmembers",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.EtcdMember { return &v1beta1.EtcdMember{} },
			func() *v1beta1.EtcdMemberList { return &v1beta1.EtcdMemberList{} }),
	}
}
