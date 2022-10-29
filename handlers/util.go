package handlers

func (h *Handler) shouldRateLimit(requestingUserIp, path string) bool {
	return !h.RateLimiters[path].GetLimiter(requestingUserIp).Allow()
}
