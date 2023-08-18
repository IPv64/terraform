package dynrpcserver

import (
	"context"
	"sync"

	tf1 "github.com/hashicorp/terraform/internal/rpcapi/terraform1"
)

type Dependencies struct {
	impl tf1.DependenciesServer
	mu   sync.RWMutex
}

var _ tf1.DependenciesServer = (*Dependencies)(nil)

func NewDependenciesStub() *Dependencies {
	return &Dependencies{}
}

func (s *Dependencies) BuildProviderPluginCache(a0 *tf1.BuildProviderPluginCache_Request, a1 tf1.Dependencies_BuildProviderPluginCacheServer) error {
	impl, err := s.realRPCServer()
	if err != nil {
		return err
	}
	return impl.BuildProviderPluginCache(a0, a1)
}

func (s *Dependencies) CloseDependencyLocks(a0 context.Context, a1 *tf1.CloseDependencyLocks_Request) (*tf1.CloseDependencyLocks_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.CloseDependencyLocks(a0, a1)
}

func (s *Dependencies) CloseProviderPluginCache(a0 context.Context, a1 *tf1.CloseProviderPluginCache_Request) (*tf1.CloseProviderPluginCache_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.CloseProviderPluginCache(a0, a1)
}

func (s *Dependencies) CloseSourceBundle(a0 context.Context, a1 *tf1.CloseSourceBundle_Request) (*tf1.CloseSourceBundle_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.CloseSourceBundle(a0, a1)
}

func (s *Dependencies) GetBuiltInProviders(a0 context.Context, a1 *tf1.GetBuiltInProviders_Request) (*tf1.GetBuiltInProviders_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.GetBuiltInProviders(a0, a1)
}

func (s *Dependencies) GetCachedProviders(a0 context.Context, a1 *tf1.GetCachedProviders_Request) (*tf1.GetCachedProviders_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.GetCachedProviders(a0, a1)
}

func (s *Dependencies) GetLockedProviderDependencies(a0 context.Context, a1 *tf1.GetLockedProviderDependencies_Request) (*tf1.GetLockedProviderDependencies_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.GetLockedProviderDependencies(a0, a1)
}

func (s *Dependencies) GetProviderSchema(a0 context.Context, a1 *tf1.GetProviderSchema_Request) (*tf1.GetProviderSchema_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.GetProviderSchema(a0, a1)
}

func (s *Dependencies) OpenDependencyLockFile(a0 context.Context, a1 *tf1.OpenDependencyLockFile_Request) (*tf1.OpenDependencyLockFile_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.OpenDependencyLockFile(a0, a1)
}

func (s *Dependencies) OpenProviderPluginCache(a0 context.Context, a1 *tf1.OpenProviderPluginCache_Request) (*tf1.OpenProviderPluginCache_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.OpenProviderPluginCache(a0, a1)
}

func (s *Dependencies) OpenSourceBundle(a0 context.Context, a1 *tf1.OpenSourceBundle_Request) (*tf1.OpenSourceBundle_Response, error) {
	impl, err := s.realRPCServer()
	if err != nil {
		return nil, err
	}
	return impl.OpenSourceBundle(a0, a1)
}

func (s *Dependencies) ActivateRPCServer(impl tf1.DependenciesServer) {
	s.mu.Lock()
	s.impl = impl
	s.mu.Unlock()
}

func (s *Dependencies) realRPCServer() (tf1.DependenciesServer, error) {
	s.mu.RLock()
	impl := s.impl
	s.mu.RUnlock()
	if impl == nil {
		return nil, unavailableErr
	}
	return impl, nil
}
