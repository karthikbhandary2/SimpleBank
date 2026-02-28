package mail

import (
	"testing"

	"github.com/karthikbhandary2/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	// skips test
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
		<h1>Hello World</h1>
		<p>This is a test message from Simple bank app by karthik</p>
	`
	to := []string{"karthikbhandary2@gmail.com"}
	// attachFiles := []string{"../.gitignore"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}