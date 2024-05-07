package email

import (
	"context"

	"encore.app/integrations/sendmail"
	"encore.app/services/user"

	"encore.dev/pubsub"
)

var _ = pubsub.NewSubscription(
	user.Signups, "send-welcome-email",
	pubsub.SubscriptionConfig[*user.SignupEvent]{
		Handler: SendWelcomeEmail,
	},
)

// SendWelcomeEmail sends a welcome email to new users
func SendWelcomeEmail(ctx context.Context, event *user.SignupEvent) error {
	u, err := user.GetUser(ctx, event.UserID)
	if err != nil {
		return err
	}
	return sendmail.SendEmail(u.Email, "Welcome to Encore!", "Welcome to Encore! We're excited to have you on board.")
}
