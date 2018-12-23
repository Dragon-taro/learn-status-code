package status

const (
	StatusContinue = 100

	StatusOK           = 200
	StatusCreated      = 201
	StatusNoContent    = 204
	StatusResetContent = 205

	StatusMovedPermanently = 301

	StatusBadRequest            = 400
	StatusUnauthorized          = 401
	StatusPaymentRequired       = 402
	StatusForbidden             = 403
	StatusNotFound              = 404
	StatusMethodNotAllowed      = 405
	StatusRequestEntityTooLarge = 413
	StatusRequestURITooLong     = 414
	StatusTooManyRequests       = 429

	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
)

var statusText = map[int]string{
	StatusContinue: "うん。",

	StatusOK:           "あ、いいよ♡",
	StatusCreated:      "できちゃった♡",
	StatusNoContent:    "いかなかった",
	StatusResetContent: "やり直しましょ",

	StatusMovedPermanently: "あっち行って",

	StatusBadRequest:            "いや、頼まれてもw",
	StatusUnauthorized:          "私たち、付き合ってないよね？",
	StatusPaymentRequired:       "１時間2万円で〜す。",
	StatusForbidden:             "今日は無理です。",
	StatusNotFound:              "お相手が見つかりませんでした。",
	StatusMethodNotAllowed:      "本番は禁止です。",
	StatusRequestEntityTooLarge: "めっちゃ出るやんw",
	StatusRequestURITooLong:     "巨根すぎわろたw",
	StatusTooManyRequests:       "激しすぎw",

	StatusInternalServerError: "EDおつ。",
	StatusBadGateway:          "穴違いですよ〜",
	StatusServiceUnavailable:  "待ってまだパンツ履いてるってw",
	StatusGatewayTimeout:      "遅漏すぎん？笑",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
