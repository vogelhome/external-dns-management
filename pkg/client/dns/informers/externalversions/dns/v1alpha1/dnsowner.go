// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	dnsv1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	versioned "github.com/gardener/external-dns-management/pkg/client/dns/clientset/versioned"
	internalinterfaces "github.com/gardener/external-dns-management/pkg/client/dns/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/external-dns-management/pkg/client/dns/listers/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSOwnerInformer provides access to a shared informer and lister for
// DNSOwners.
type DNSOwnerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DNSOwnerLister
}

type dNSOwnerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSOwnerInformer constructs a new informer for DNSOwner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSOwnerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSOwnerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSOwnerInformer constructs a new informer for DNSOwner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSOwnerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSOwners(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSOwners(namespace).Watch(context.TODO(), options)
			},
		},
		&dnsv1alpha1.DNSOwner{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSOwnerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSOwnerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSOwnerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dnsv1alpha1.DNSOwner{}, f.defaultInformer)
}

func (f *dNSOwnerInformer) Lister() v1alpha1.DNSOwnerLister {
	return v1alpha1.NewDNSOwnerLister(f.Informer().GetIndexer())
}
