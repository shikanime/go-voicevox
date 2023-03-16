/*
VOICEVOX Engine

Testing AudioLibraryManagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package voicevox

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github/infinity-blackhole/go-voicevox"
)

func Test_voicevox_AudioLibraryManagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AudioLibraryManagementApiService DownloadableLibrariesDownloadableLibrariesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AudioLibraryManagementApi.DownloadableLibrariesDownloadableLibrariesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioLibraryManagementApiService InstallLibraryInstallLibraryLibraryUuidPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryUuid string

		httpRes, err := apiClient.AudioLibraryManagementApi.InstallLibraryInstallLibraryLibraryUuidPost(context.Background(), libraryUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioLibraryManagementApiService InstalledLibrariesInstalledLibrariesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AudioLibraryManagementApi.InstalledLibrariesInstalledLibrariesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
