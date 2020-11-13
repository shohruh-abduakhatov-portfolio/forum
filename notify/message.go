package notify

type Format string

var (
	fmtLikedPost   = Format(`User %s liked your post:"%s"`)
	fmtDisikedPost = Format(`User %s disliked your post:"%s"`)
)
