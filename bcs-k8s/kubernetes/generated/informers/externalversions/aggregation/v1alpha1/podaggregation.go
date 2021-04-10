/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	aggregationv1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/apis/aggregation/v1alpha1"
	clientset "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/generated/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/generated/listers/aggregation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodAggregationInformer provides access to a shared informer and lister for
// PodAggregations.
type PodAggregationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodAggregationLister
}

type podAggregationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodAggregationInformer constructs a new informer for PodAggregation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodAggregationInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodAggregationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodAggregationInformer constructs a new informer for PodAggregation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodAggregationInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AggregationV1alpha1().PodAggregations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AggregationV1alpha1().PodAggregations(namespace).Watch(context.TODO(), options)
			},
		},
		&aggregationv1alpha1.PodAggregation{},
		resyncPeriod,
		indexers,
	)
}

func (f *podAggregationInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodAggregationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podAggregationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aggregationv1alpha1.PodAggregation{}, f.defaultInformer)
}

func (f *podAggregationInformer) Lister() v1alpha1.PodAggregationLister {
	return v1alpha1.NewPodAggregationLister(f.Informer().GetIndexer())
}
