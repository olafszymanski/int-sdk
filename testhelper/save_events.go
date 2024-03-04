package testhelper

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/olafszymanski/int-sdk/integration/pb"
	"github.com/stretchr/testify/require"
)

func SaveEvents(t *testing.T, events []*pb.Event, file *os.File) {
	b, err := json.MarshalIndent(events, "", "   ")
	require.NoError(t, err)

	_, err = file.Write(b)
	require.NoError(t, err)
}
