// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx-api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx-api/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PluginInformer provides access to a shared informer and lister for
// Plugins.
type PluginInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PluginLister
}

type pluginInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPluginInformer constructs a new informer for Plugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPluginInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPluginInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPluginInformer constructs a new informer for Plugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPluginInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Plugins(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Plugins(namespace).Watch(options)
			},
		},
		&jenkinsiov1.Plugin{},
		resyncPeriod,
		indexers,
	)
}

func (f *pluginInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPluginInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pluginInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.Plugin{}, f.defaultInformer)
}

func (f *pluginInformer) Lister() v1.PluginLister {
	return v1.NewPluginLister(f.Informer().GetIndexer())
}
