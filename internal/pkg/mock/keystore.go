// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: keystore.go

package mock

import (
	sync "sync"

	github_com_confluentinc_cli_internal_pkg_config_v1 "github.com/confluentinc/cli/internal/pkg/config/v1"
)

// KeyStore is a mock of KeyStore interface
type KeyStore struct {
	lockHasAPIKey sync.Mutex
	HasAPIKeyFunc func(key, clusterId string) (bool, error)

	lockStoreAPIKey sync.Mutex
	StoreAPIKeyFunc func(key *github_com_confluentinc_cli_internal_pkg_config_v1.APIKeyPair, clusterId string) error

	lockDeleteAPIKey sync.Mutex
	DeleteAPIKeyFunc func(key string) error

	calls struct {
		HasAPIKey []struct {
			Key       string
			ClusterId string
		}
		StoreAPIKey []struct {
			Key       *github_com_confluentinc_cli_internal_pkg_config_v1.APIKeyPair
			ClusterId string
		}
		DeleteAPIKey []struct {
			Key string
		}
	}
}

// HasAPIKey mocks base method by wrapping the associated func.
func (m *KeyStore) HasAPIKey(key, clusterId string) (bool, error) {
	m.lockHasAPIKey.Lock()
	defer m.lockHasAPIKey.Unlock()

	if m.HasAPIKeyFunc == nil {
		panic("mocker: KeyStore.HasAPIKeyFunc is nil but KeyStore.HasAPIKey was called.")
	}

	call := struct {
		Key       string
		ClusterId string
	}{
		Key:       key,
		ClusterId: clusterId,
	}

	m.calls.HasAPIKey = append(m.calls.HasAPIKey, call)

	return m.HasAPIKeyFunc(key, clusterId)
}

// HasAPIKeyCalled returns true if HasAPIKey was called at least once.
func (m *KeyStore) HasAPIKeyCalled() bool {
	m.lockHasAPIKey.Lock()
	defer m.lockHasAPIKey.Unlock()

	return len(m.calls.HasAPIKey) > 0
}

// HasAPIKeyCalls returns the calls made to HasAPIKey.
func (m *KeyStore) HasAPIKeyCalls() []struct {
	Key       string
	ClusterId string
} {
	m.lockHasAPIKey.Lock()
	defer m.lockHasAPIKey.Unlock()

	return m.calls.HasAPIKey
}

// StoreAPIKey mocks base method by wrapping the associated func.
func (m *KeyStore) StoreAPIKey(key *github_com_confluentinc_cli_internal_pkg_config_v1.APIKeyPair, clusterId string) error {
	m.lockStoreAPIKey.Lock()
	defer m.lockStoreAPIKey.Unlock()

	if m.StoreAPIKeyFunc == nil {
		panic("mocker: KeyStore.StoreAPIKeyFunc is nil but KeyStore.StoreAPIKey was called.")
	}

	call := struct {
		Key       *github_com_confluentinc_cli_internal_pkg_config_v1.APIKeyPair
		ClusterId string
	}{
		Key:       key,
		ClusterId: clusterId,
	}

	m.calls.StoreAPIKey = append(m.calls.StoreAPIKey, call)

	return m.StoreAPIKeyFunc(key, clusterId)
}

// StoreAPIKeyCalled returns true if StoreAPIKey was called at least once.
func (m *KeyStore) StoreAPIKeyCalled() bool {
	m.lockStoreAPIKey.Lock()
	defer m.lockStoreAPIKey.Unlock()

	return len(m.calls.StoreAPIKey) > 0
}

// StoreAPIKeyCalls returns the calls made to StoreAPIKey.
func (m *KeyStore) StoreAPIKeyCalls() []struct {
	Key       *github_com_confluentinc_cli_internal_pkg_config_v1.APIKeyPair
	ClusterId string
} {
	m.lockStoreAPIKey.Lock()
	defer m.lockStoreAPIKey.Unlock()

	return m.calls.StoreAPIKey
}

// DeleteAPIKey mocks base method by wrapping the associated func.
func (m *KeyStore) DeleteAPIKey(key string) error {
	m.lockDeleteAPIKey.Lock()
	defer m.lockDeleteAPIKey.Unlock()

	if m.DeleteAPIKeyFunc == nil {
		panic("mocker: KeyStore.DeleteAPIKeyFunc is nil but KeyStore.DeleteAPIKey was called.")
	}

	call := struct {
		Key string
	}{
		Key: key,
	}

	m.calls.DeleteAPIKey = append(m.calls.DeleteAPIKey, call)

	return m.DeleteAPIKeyFunc(key)
}

// DeleteAPIKeyCalled returns true if DeleteAPIKey was called at least once.
func (m *KeyStore) DeleteAPIKeyCalled() bool {
	m.lockDeleteAPIKey.Lock()
	defer m.lockDeleteAPIKey.Unlock()

	return len(m.calls.DeleteAPIKey) > 0
}

// DeleteAPIKeyCalls returns the calls made to DeleteAPIKey.
func (m *KeyStore) DeleteAPIKeyCalls() []struct {
	Key string
} {
	m.lockDeleteAPIKey.Lock()
	defer m.lockDeleteAPIKey.Unlock()

	return m.calls.DeleteAPIKey
}

// Reset resets the calls made to the mocked methods.
func (m *KeyStore) Reset() {
	m.lockHasAPIKey.Lock()
	m.calls.HasAPIKey = nil
	m.lockHasAPIKey.Unlock()
	m.lockStoreAPIKey.Lock()
	m.calls.StoreAPIKey = nil
	m.lockStoreAPIKey.Unlock()
	m.lockDeleteAPIKey.Lock()
	m.calls.DeleteAPIKey = nil
	m.lockDeleteAPIKey.Unlock()
}
