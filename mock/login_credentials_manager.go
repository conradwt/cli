// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: login_credentials_manager.go

package mock

import (
	sync "sync"

	github_com_confluentinc_ccloud_sdk_go "github.com/confluentinc/ccloud-sdk-go"
	github_com_spf13_cobra "github.com/spf13/cobra"

	github_com_confluentinc_cli_internal_pkg_auth "github.com/confluentinc/cli/internal/pkg/auth"
	github_com_confluentinc_cli_internal_pkg_netrc "github.com/confluentinc/cli/internal/pkg/netrc"
)

// MockLoginCredentialsManager is a mock of LoginCredentialsManager interface
type MockLoginCredentialsManager struct {
	lockGetCCloudCredentialsFromEnvVar sync.Mutex
	GetCCloudCredentialsFromEnvVarFunc func(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetCCloudCredentialsFromPrompt sync.Mutex
	GetCCloudCredentialsFromPromptFunc func(cmd *github_com_spf13_cobra.Command, client *github_com_confluentinc_ccloud_sdk_go.Client) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetConfluentCredentialsFromEnvVar sync.Mutex
	GetConfluentCredentialsFromEnvVarFunc func(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetConfluentCredentialsFromPrompt sync.Mutex
	GetConfluentCredentialsFromPromptFunc func(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetCredentialsFromNetrc sync.Mutex
	GetCredentialsFromNetrcFunc func(cmd *github_com_spf13_cobra.Command, filterParams github_com_confluentinc_cli_internal_pkg_netrc.GetMatchingNetrcMachineParams) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetConfluentPrerunCredentialsFromEnvVar sync.Mutex
	GetConfluentPrerunCredentialsFromEnvVarFunc func(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	lockGetConfluentPrerunCredentialsFromNetrc sync.Mutex
	GetConfluentPrerunCredentialsFromNetrcFunc func(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error)

	calls struct {
		GetCCloudCredentialsFromEnvVar []struct {
			Cmd *github_com_spf13_cobra.Command
		}
		GetCCloudCredentialsFromPrompt []struct {
			Cmd    *github_com_spf13_cobra.Command
			Client *github_com_confluentinc_ccloud_sdk_go.Client
		}
		GetConfluentCredentialsFromEnvVar []struct {
			Cmd *github_com_spf13_cobra.Command
		}
		GetConfluentCredentialsFromPrompt []struct {
			Cmd *github_com_spf13_cobra.Command
		}
		GetCredentialsFromNetrc []struct {
			Cmd          *github_com_spf13_cobra.Command
			FilterParams github_com_confluentinc_cli_internal_pkg_netrc.GetMatchingNetrcMachineParams
		}
		GetConfluentPrerunCredentialsFromEnvVar []struct {
			Cmd *github_com_spf13_cobra.Command
		}
		GetConfluentPrerunCredentialsFromNetrc []struct {
			Cmd *github_com_spf13_cobra.Command
		}
	}
}

// GetCCloudCredentialsFromEnvVar mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromEnvVar(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetCCloudCredentialsFromEnvVar.Lock()
	defer m.lockGetCCloudCredentialsFromEnvVar.Unlock()

	if m.GetCCloudCredentialsFromEnvVarFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetCCloudCredentialsFromEnvVarFunc is nil but MockLoginCredentialsManager.GetCCloudCredentialsFromEnvVar was called.")
	}

	call := struct {
		Cmd *github_com_spf13_cobra.Command
	}{
		Cmd: cmd,
	}

	m.calls.GetCCloudCredentialsFromEnvVar = append(m.calls.GetCCloudCredentialsFromEnvVar, call)

	return m.GetCCloudCredentialsFromEnvVarFunc(cmd)
}

// GetCCloudCredentialsFromEnvVarCalled returns true if GetCCloudCredentialsFromEnvVar was called at least once.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromEnvVarCalled() bool {
	m.lockGetCCloudCredentialsFromEnvVar.Lock()
	defer m.lockGetCCloudCredentialsFromEnvVar.Unlock()

	return len(m.calls.GetCCloudCredentialsFromEnvVar) > 0
}

// GetCCloudCredentialsFromEnvVarCalls returns the calls made to GetCCloudCredentialsFromEnvVar.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromEnvVarCalls() []struct {
	Cmd *github_com_spf13_cobra.Command
} {
	m.lockGetCCloudCredentialsFromEnvVar.Lock()
	defer m.lockGetCCloudCredentialsFromEnvVar.Unlock()

	return m.calls.GetCCloudCredentialsFromEnvVar
}

// GetCCloudCredentialsFromPrompt mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromPrompt(cmd *github_com_spf13_cobra.Command, client *github_com_confluentinc_ccloud_sdk_go.Client) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetCCloudCredentialsFromPrompt.Lock()
	defer m.lockGetCCloudCredentialsFromPrompt.Unlock()

	if m.GetCCloudCredentialsFromPromptFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetCCloudCredentialsFromPromptFunc is nil but MockLoginCredentialsManager.GetCCloudCredentialsFromPrompt was called.")
	}

	call := struct {
		Cmd    *github_com_spf13_cobra.Command
		Client *github_com_confluentinc_ccloud_sdk_go.Client
	}{
		Cmd:    cmd,
		Client: client,
	}

	m.calls.GetCCloudCredentialsFromPrompt = append(m.calls.GetCCloudCredentialsFromPrompt, call)

	return m.GetCCloudCredentialsFromPromptFunc(cmd, client)
}

// GetCCloudCredentialsFromPromptCalled returns true if GetCCloudCredentialsFromPrompt was called at least once.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromPromptCalled() bool {
	m.lockGetCCloudCredentialsFromPrompt.Lock()
	defer m.lockGetCCloudCredentialsFromPrompt.Unlock()

	return len(m.calls.GetCCloudCredentialsFromPrompt) > 0
}

// GetCCloudCredentialsFromPromptCalls returns the calls made to GetCCloudCredentialsFromPrompt.
func (m *MockLoginCredentialsManager) GetCCloudCredentialsFromPromptCalls() []struct {
	Cmd    *github_com_spf13_cobra.Command
	Client *github_com_confluentinc_ccloud_sdk_go.Client
} {
	m.lockGetCCloudCredentialsFromPrompt.Lock()
	defer m.lockGetCCloudCredentialsFromPrompt.Unlock()

	return m.calls.GetCCloudCredentialsFromPrompt
}

// GetConfluentCredentialsFromEnvVar mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromEnvVar(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetConfluentCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentCredentialsFromEnvVar.Unlock()

	if m.GetConfluentCredentialsFromEnvVarFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetConfluentCredentialsFromEnvVarFunc is nil but MockLoginCredentialsManager.GetConfluentCredentialsFromEnvVar was called.")
	}

	call := struct {
		Cmd *github_com_spf13_cobra.Command
	}{
		Cmd: cmd,
	}

	m.calls.GetConfluentCredentialsFromEnvVar = append(m.calls.GetConfluentCredentialsFromEnvVar, call)

	return m.GetConfluentCredentialsFromEnvVarFunc(cmd)
}

// GetConfluentCredentialsFromEnvVarCalled returns true if GetConfluentCredentialsFromEnvVar was called at least once.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromEnvVarCalled() bool {
	m.lockGetConfluentCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentCredentialsFromEnvVar.Unlock()

	return len(m.calls.GetConfluentCredentialsFromEnvVar) > 0
}

// GetConfluentCredentialsFromEnvVarCalls returns the calls made to GetConfluentCredentialsFromEnvVar.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromEnvVarCalls() []struct {
	Cmd *github_com_spf13_cobra.Command
} {
	m.lockGetConfluentCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentCredentialsFromEnvVar.Unlock()

	return m.calls.GetConfluentCredentialsFromEnvVar
}

// GetConfluentCredentialsFromPrompt mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromPrompt(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetConfluentCredentialsFromPrompt.Lock()
	defer m.lockGetConfluentCredentialsFromPrompt.Unlock()

	if m.GetConfluentCredentialsFromPromptFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetConfluentCredentialsFromPromptFunc is nil but MockLoginCredentialsManager.GetConfluentCredentialsFromPrompt was called.")
	}

	call := struct {
		Cmd *github_com_spf13_cobra.Command
	}{
		Cmd: cmd,
	}

	m.calls.GetConfluentCredentialsFromPrompt = append(m.calls.GetConfluentCredentialsFromPrompt, call)

	return m.GetConfluentCredentialsFromPromptFunc(cmd)
}

// GetConfluentCredentialsFromPromptCalled returns true if GetConfluentCredentialsFromPrompt was called at least once.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromPromptCalled() bool {
	m.lockGetConfluentCredentialsFromPrompt.Lock()
	defer m.lockGetConfluentCredentialsFromPrompt.Unlock()

	return len(m.calls.GetConfluentCredentialsFromPrompt) > 0
}

// GetConfluentCredentialsFromPromptCalls returns the calls made to GetConfluentCredentialsFromPrompt.
func (m *MockLoginCredentialsManager) GetConfluentCredentialsFromPromptCalls() []struct {
	Cmd *github_com_spf13_cobra.Command
} {
	m.lockGetConfluentCredentialsFromPrompt.Lock()
	defer m.lockGetConfluentCredentialsFromPrompt.Unlock()

	return m.calls.GetConfluentCredentialsFromPrompt
}

// GetCredentialsFromNetrc mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetCredentialsFromNetrc(cmd *github_com_spf13_cobra.Command, filterParams github_com_confluentinc_cli_internal_pkg_netrc.GetMatchingNetrcMachineParams) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetCredentialsFromNetrc.Lock()
	defer m.lockGetCredentialsFromNetrc.Unlock()

	if m.GetCredentialsFromNetrcFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetCredentialsFromNetrcFunc is nil but MockLoginCredentialsManager.GetCredentialsFromNetrc was called.")
	}

	call := struct {
		Cmd          *github_com_spf13_cobra.Command
		FilterParams github_com_confluentinc_cli_internal_pkg_netrc.GetMatchingNetrcMachineParams
	}{
		Cmd:          cmd,
		FilterParams: filterParams,
	}

	m.calls.GetCredentialsFromNetrc = append(m.calls.GetCredentialsFromNetrc, call)

	return m.GetCredentialsFromNetrcFunc(cmd, filterParams)
}

// GetCredentialsFromNetrcCalled returns true if GetCredentialsFromNetrc was called at least once.
func (m *MockLoginCredentialsManager) GetCredentialsFromNetrcCalled() bool {
	m.lockGetCredentialsFromNetrc.Lock()
	defer m.lockGetCredentialsFromNetrc.Unlock()

	return len(m.calls.GetCredentialsFromNetrc) > 0
}

// GetCredentialsFromNetrcCalls returns the calls made to GetCredentialsFromNetrc.
func (m *MockLoginCredentialsManager) GetCredentialsFromNetrcCalls() []struct {
	Cmd          *github_com_spf13_cobra.Command
	FilterParams github_com_confluentinc_cli_internal_pkg_netrc.GetMatchingNetrcMachineParams
} {
	m.lockGetCredentialsFromNetrc.Lock()
	defer m.lockGetCredentialsFromNetrc.Unlock()

	return m.calls.GetCredentialsFromNetrc
}

// GetConfluentPrerunCredentialsFromEnvVar mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromEnvVar(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetConfluentPrerunCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromEnvVar.Unlock()

	if m.GetConfluentPrerunCredentialsFromEnvVarFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetConfluentPrerunCredentialsFromEnvVarFunc is nil but MockLoginCredentialsManager.GetConfluentPrerunCredentialsFromEnvVar was called.")
	}

	call := struct {
		Cmd *github_com_spf13_cobra.Command
	}{
		Cmd: cmd,
	}

	m.calls.GetConfluentPrerunCredentialsFromEnvVar = append(m.calls.GetConfluentPrerunCredentialsFromEnvVar, call)

	return m.GetConfluentPrerunCredentialsFromEnvVarFunc(cmd)
}

// GetConfluentPrerunCredentialsFromEnvVarCalled returns true if GetConfluentPrerunCredentialsFromEnvVar was called at least once.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromEnvVarCalled() bool {
	m.lockGetConfluentPrerunCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromEnvVar.Unlock()

	return len(m.calls.GetConfluentPrerunCredentialsFromEnvVar) > 0
}

// GetConfluentPrerunCredentialsFromEnvVarCalls returns the calls made to GetConfluentPrerunCredentialsFromEnvVar.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromEnvVarCalls() []struct {
	Cmd *github_com_spf13_cobra.Command
} {
	m.lockGetConfluentPrerunCredentialsFromEnvVar.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromEnvVar.Unlock()

	return m.calls.GetConfluentPrerunCredentialsFromEnvVar
}

// GetConfluentPrerunCredentialsFromNetrc mocks base method by wrapping the associated func.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromNetrc(cmd *github_com_spf13_cobra.Command) func() (*github_com_confluentinc_cli_internal_pkg_auth.Credentials, error) {
	m.lockGetConfluentPrerunCredentialsFromNetrc.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromNetrc.Unlock()

	if m.GetConfluentPrerunCredentialsFromNetrcFunc == nil {
		panic("mocker: MockLoginCredentialsManager.GetConfluentPrerunCredentialsFromNetrcFunc is nil but MockLoginCredentialsManager.GetConfluentPrerunCredentialsFromNetrc was called.")
	}

	call := struct {
		Cmd *github_com_spf13_cobra.Command
	}{
		Cmd: cmd,
	}

	m.calls.GetConfluentPrerunCredentialsFromNetrc = append(m.calls.GetConfluentPrerunCredentialsFromNetrc, call)

	return m.GetConfluentPrerunCredentialsFromNetrcFunc(cmd)
}

// GetConfluentPrerunCredentialsFromNetrcCalled returns true if GetConfluentPrerunCredentialsFromNetrc was called at least once.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromNetrcCalled() bool {
	m.lockGetConfluentPrerunCredentialsFromNetrc.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromNetrc.Unlock()

	return len(m.calls.GetConfluentPrerunCredentialsFromNetrc) > 0
}

// GetConfluentPrerunCredentialsFromNetrcCalls returns the calls made to GetConfluentPrerunCredentialsFromNetrc.
func (m *MockLoginCredentialsManager) GetConfluentPrerunCredentialsFromNetrcCalls() []struct {
	Cmd *github_com_spf13_cobra.Command
} {
	m.lockGetConfluentPrerunCredentialsFromNetrc.Lock()
	defer m.lockGetConfluentPrerunCredentialsFromNetrc.Unlock()

	return m.calls.GetConfluentPrerunCredentialsFromNetrc
}

// Reset resets the calls made to the mocked methods.
func (m *MockLoginCredentialsManager) Reset() {
	m.lockGetCCloudCredentialsFromEnvVar.Lock()
	m.calls.GetCCloudCredentialsFromEnvVar = nil
	m.lockGetCCloudCredentialsFromEnvVar.Unlock()
	m.lockGetCCloudCredentialsFromPrompt.Lock()
	m.calls.GetCCloudCredentialsFromPrompt = nil
	m.lockGetCCloudCredentialsFromPrompt.Unlock()
	m.lockGetConfluentCredentialsFromEnvVar.Lock()
	m.calls.GetConfluentCredentialsFromEnvVar = nil
	m.lockGetConfluentCredentialsFromEnvVar.Unlock()
	m.lockGetConfluentCredentialsFromPrompt.Lock()
	m.calls.GetConfluentCredentialsFromPrompt = nil
	m.lockGetConfluentCredentialsFromPrompt.Unlock()
	m.lockGetCredentialsFromNetrc.Lock()
	m.calls.GetCredentialsFromNetrc = nil
	m.lockGetCredentialsFromNetrc.Unlock()
	m.lockGetConfluentPrerunCredentialsFromEnvVar.Lock()
	m.calls.GetConfluentPrerunCredentialsFromEnvVar = nil
	m.lockGetConfluentPrerunCredentialsFromEnvVar.Unlock()
	m.lockGetConfluentPrerunCredentialsFromNetrc.Lock()
	m.calls.GetConfluentPrerunCredentialsFromNetrc = nil
	m.lockGetConfluentPrerunCredentialsFromNetrc.Unlock()
}
