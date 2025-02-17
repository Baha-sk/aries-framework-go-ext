/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package jwk_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go-ext/component/vdr/jwk"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
)

var _ vdr.VDR = (*jwk.VDR)(nil) // verify interface compliance

func TestAccept(t *testing.T) {
	t.Run("jwk method", func(t *testing.T) {
		v := jwk.New()
		require.NotNil(t, v)

		accept := v.Accept("jwk")
		require.True(t, accept)
	})

	t.Run("other method", func(t *testing.T) {
		v := jwk.New()
		require.NotNil(t, v)

		accept := v.Accept("other")
		require.False(t, accept)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test update", func(t *testing.T) {
		v := jwk.New()
		err := v.Update(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not supported")
	})
}

func TestDeactivate(t *testing.T) {
	t.Run("test deactivate", func(t *testing.T) {
		v := jwk.New()
		err := v.Deactivate("")
		require.Error(t, err)
		require.Contains(t, err.Error(), "not supported")
	})
}

func TestClose(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		v := jwk.New()
		require.NotNil(t, v)
		require.NoError(t, v.Close())
	})
}
