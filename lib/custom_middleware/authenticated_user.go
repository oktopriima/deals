package custom_middleware

import (
	"context"
	"fmt"
	"strconv"
)

func GetAuthenticatedUser(ctx context.Context) (int64, error) {
	ctxUser := ctx.Value(AuthUser)
	userIdStr, ok := ctxUser.(string)
	if !ok {
		return 0, fmt.Errorf("user not found")
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid type user id")
	}

	return userId, nil
}
