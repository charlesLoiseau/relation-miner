package model

type Follower struct {
	Users             []User `json:"users"`
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

type FollowerId struct {
	Ids               []int  `json:"ids"`
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}
