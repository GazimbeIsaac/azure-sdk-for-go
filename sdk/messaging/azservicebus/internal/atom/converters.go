// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package atom

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/internal/auth"
)

func WrapWithQueueEnvelope(qd *QueueDescription, tokenProvider auth.TokenProvider) (*QueueEnvelope, []MiddlewareFunc) {
	qd.ServiceBusSchema = to.Ptr(serviceBusSchema)

	qe := &QueueEnvelope{
		Entry: &Entry{
			AtomSchema: atomSchema,
		},
		Content: &queueContent{
			Type:             applicationXML,
			QueueDescription: *qd,
		},
	}

	var mw []MiddlewareFunc
	if qd.ForwardTo != nil {
		mw = append(mw, addSupplementalAuthorization(*qd.ForwardTo, tokenProvider))
	}

	if qd.ForwardDeadLetteredMessagesTo != nil {
		mw = append(mw, addDeadLetterSupplementalAuthorization(*qd.ForwardDeadLetteredMessagesTo, tokenProvider))
	}

	return qe, mw
}

func WrapWithTopicEnvelope(td *TopicDescription) *TopicEnvelope {
	td.ServiceBusSchema = to.Ptr(serviceBusSchema)

	return &TopicEnvelope{
		Entry: &Entry{
			AtomSchema: atomSchema,
		},
		Content: &topicContent{
			Type:             applicationXML,
			TopicDescription: *td,
		},
	}
}

func WrapWithSubscriptionEnvelope(sd *SubscriptionDescription) *SubscriptionEnvelope {
	sd.ServiceBusSchema = to.Ptr(serviceBusSchema)

	return &SubscriptionEnvelope{
		Entry: &Entry{
			AtomSchema: atomSchema,
		},
		Content: &subscriptionContent{
			Type:                    applicationXML,
			SubscriptionDescription: *sd,
		},
	}
}

func WrapWithRuleEnvelope(rd *RuleDescription) *RuleEnvelope {
	rd.XMLNS = "http://schemas.microsoft.com/netservices/2010/10/servicebus/connect"
	rd.XMLNSI = "http://www.w3.org/2001/XMLSchema-instance"

	return &RuleEnvelope{
		Entry: &Entry{
			AtomSchema: atomSchema,
		},
		Content: &RuleContent{
			Type:            applicationXML,
			RuleDescription: *rd,
		},
	}
}
