wrk.method = "POST"

wrk.headers["Content-Type"] = "application/json"
wrk.headers["Cookie"] = "session_token=;"

wrk.headers["X-CSRF-Token"] = ""

request = function()
  local body = string.format(
    '{"title":"perf playlist %d","description":"load test"}',
    math.random(1, 1000000000)
  )
  return wrk.format(nil, "/api/v1/playlists", nil, body)
end