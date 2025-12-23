wrk.method = "GET"

wrk.headers["Cookie"] = "session_token=;"

return wrk.format(
	nil,
	"/api/v1/playlists/my?limit=10&offset=5000",
	nil,
	nil
)
end