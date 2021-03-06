package apitest

import (
	"math/rand"
	"testing"
	"time"

	"github.com/romshark/dgraph_graphql_go/api/graph/gqlmod"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int, runes []rune) string {
	if runes == nil {
		runes = []rune("a")
	}
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func compareUsers(
	t *testing.T,
	expected *gqlmod.User,
	actual *gqlmod.User,
) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)

	// id
	require.NotNil(t, expected.ID)
	require.NotNil(t, actual.ID)
	require.Equal(t, *expected.ID, *actual.ID)

	// email
	require.NotNil(t, expected.Email)
	require.NotNil(t, actual.Email)
	require.Equal(t, *expected.Email, *actual.Email)

	// displayName
	require.NotNil(t, expected.DisplayName)
	require.NotNil(t, actual.DisplayName)
	require.Equal(t, *expected.DisplayName, *actual.DisplayName)

	// creation
	require.NotNil(t, expected.Creation)
	require.NotNil(t, actual.Creation)
	require.Equal(t, *expected.Creation, *actual.Creation)
}

func comparePosts(
	t *testing.T,
	expected *gqlmod.Post,
	actual *gqlmod.Post,
) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)

	// id
	require.NotNil(t, expected.ID)
	require.NotNil(t, actual.ID)
	require.Equal(t, *expected.ID, *actual.ID)

	// title
	require.NotNil(t, expected.Title)
	require.NotNil(t, actual.Title)
	require.Equal(t, *expected.Title, *actual.Title)

	// contents
	require.NotNil(t, expected.Contents)
	require.NotNil(t, actual.Contents)
	require.Equal(t, *expected.Contents, *actual.Contents)

	// creation
	require.NotNil(t, expected.Creation)
	require.NotNil(t, actual.Creation)
	require.Equal(t, *expected.Creation, *actual.Creation)
}

func compareReactions(
	t *testing.T,
	expected *gqlmod.Reaction,
	actual *gqlmod.Reaction,
) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)

	// id
	require.NotNil(t, expected.ID)
	require.NotNil(t, actual.ID)
	require.Equal(t, *expected.ID, *actual.ID)

	// emotion
	require.NotNil(t, expected.Emotion)
	require.NotNil(t, actual.Emotion)
	require.Equal(t, *expected.Emotion, *actual.Emotion)

	// message
	require.NotNil(t, expected.Message)
	require.NotNil(t, actual.Message)
	require.Equal(t, *expected.Message, *actual.Message)

	// creation
	require.NotNil(t, expected.Creation)
	require.NotNil(t, actual.Creation)
	require.Equal(t, *expected.Creation, *actual.Creation)
}
