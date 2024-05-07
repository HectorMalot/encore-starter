package user

import (
	"encore.dev/pubsub"
	"encore.dev/types/uuid"
)

type SignupEvent struct{ UserID uuid.UUID }

var Signups = pubsub.NewTopic[*SignupEvent]("signups", pubsub.TopicConfig{
	DeliveryGuarantee: pubsub.AtLeastOnce,
})
