wrk.method = "GET"

wrk.headers["Cookie"] = "session_token=;"
request = function()
	return wrk.format(
		nil,
		"/api/v1/playlists/my?limit=10&offset=0",
		nil,
		nil
	)
end