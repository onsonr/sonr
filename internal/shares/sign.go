package shares

// func signV1(curve *curves.Curve, aliceDkgResultMessage *protocol.Message, bobDkgResultMessage *protocol.Message) {
// 	// New DklsSign
// 	msg := []byte("As soon as you trust yourself, you will know how to live.")
// 	aliceSign, err := NewAliceSign(curve, sha3.New256(), msg, aliceDkgResultMessage, protocol.Version1)
// 	require.NoError(t, err)
// 	bobSign, err := NewBobSign(curve, sha3.New256(), msg, bobDkgResultMessage, protocol.Version1)
// 	require.NoError(t, err)

// 	// Sign
// 	var aErr error
// 	var bErr error
// 	t.Run("sign", func(t *testing.T) {
// 		aErr, bErr = require.ErrorIs(t, aErr, protocol.ErrProtocolFinished)
// 		require.ErrorIs(t, bErr, protocol.ErrProtocolFinished)
// 	})
// 	// Don't continue to verifying results if sign didn't run to completion.
// 	require.ErrorIs(t, aErr, protocol.ErrProtocolFinished)
// 	require.ErrorIs(t, bErr, protocol.ErrProtocolFinished)

// 	// Output
// 	var result *curves.EcdsaSignature
// 	t.Run("bob produces result of correct type", func(t *testing.T) {
// 		resultMessage, err := bobSign.Result(protocol.Version1)
// 		require.NoError(t, err)
// 		result, err = DecodeSignature(resultMessage)
// 		require.NoError(t, err)
// 	})
// 	require.NotNil(t, result)

// 	aliceDkg, err := DecodeAliceDkgResult(aliceDkgResultMessage)
// 	require.NoError(t, err)

// 	t.Run("valid signature", func(t *testing.T) {
// 		hash := sha3.New256()
// 		_, err = hash.Write(msg)
// 		require.NoError(t, err)
// 		digest := hash.Sum(nil)
// 		unCompressedAffinePublicKey := aliceDkg.PublicKey.ToAffineUncompressed()
// 		require.Equal(t, 65, len(unCompressedAffinePublicKey))
// 		x := new(big.Int).SetBytes(unCompressedAffinePublicKey[1:33])
// 		y := new(big.Int).SetBytes(unCompressedAffinePublicKey[33:])
// 		ecCurve, err := curve.ToEllipticCurve()
// 		require.NoError(t, err)
// 		publicKey := &curves.EcPoint{
// 			Curve: ecCurve,
// 			X:     x,
// 			Y:     y,
// 		}
// 		require.True(t,
// 			curves.VerifyEcdsa(publicKey,
// 				digest[:],
// 				result,
// 			),
// 			"signature failed verification",
// 		)
// 	})
// }

// func refreshV1(t *testing.T, curve *curves.Curve, aliceDkgResultMessage, bobDkgResultMessage *protocol.Message) (aliceRefreshResultMessage, bobRefreshResultMessage *protocol.Message) {
// 	t.Helper()
// 	aliceRefresh, err := NewAliceRefresh(curve, aliceDkgResultMessage, protocol.Version1)
// 	require.NoError(t, err)
// 	bobRefresh, err := NewBobRefresh(curve, bobDkgResultMessage, protocol.Version1)
// 	require.NoError(t, err)

// 	aErr, bErr := runIteratedProtocol(aliceRefresh, bobRefresh)
// 	require.ErrorIs(t, aErr, protocol.ErrProtocolFinished)
// 	require.ErrorIs(t, bErr, protocol.ErrProtocolFinished)

// 	aliceRefreshResultMessage, err = aliceRefresh.Result(protocol.Version1)
// 	require.NoError(t, err)
// 	require.NotNil(t, aliceRefreshResultMessage)
// 	_, err = DecodeAliceRefreshResult(aliceRefreshResultMessage)
// 	require.NoError(t, err)

// 	bobRefreshResultMessage, err = bobRefresh.Result(protocol.Version1)
// 	require.NoError(t, err)
// 	require.NotNil(t, bobRefreshResultMessage)
// 	_, err = DecodeBobRefreshResult(bobRefreshResultMessage)
// 	require.NoError(t, err)

// 	return aliceRefreshResultMessage, bobRefreshResultMessage
// }
