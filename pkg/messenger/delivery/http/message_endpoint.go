package http

import (
	"time"

	"github.com/jasonsoft/napnap"
	"github.com/jasonsoft/wakanda/internal/identity"
	"github.com/jasonsoft/wakanda/internal/pagination"
	"github.com/jasonsoft/wakanda/internal/types"
	"github.com/jasonsoft/wakanda/pkg/messenger"
)

func (h *MessengerHandler) messageMeListEndpoint(c *napnap.Context) {
	ctx := c.StdContext()
	pagination := pagination.FromContext(c)

	claim, found := identity.FromContext(ctx)
	if found == false {
		c.SetStatus(403)
		return
	}

	listMsgOpts := &messenger.FindMessagesOptions{
		MemberID: claim.UserID,
		Skip:     pagination.Skip(),
		PerPage:  pagination.PerPage,
	}

	anchorUpdatedAtStr := c.Query("anchor_updated_at")
	if len(anchorUpdatedAtStr) > 0 {
		anchorUpdatedAt, err := time.Parse(time.RFC3339, anchorUpdatedAtStr)
		if err != nil {
			panic(types.AppError{ErrorCode: "invalid_input", Message: "anchor_updated_at field was invalid"})
		}
		listMsgOpts.AnchorUpdatedAt = &anchorUpdatedAt
	}

	messages, err := h.messageService.Messages(ctx, listMsgOpts)
	if err != nil {
		panic(err)
	}

	c.JSON(200, messages)
}
