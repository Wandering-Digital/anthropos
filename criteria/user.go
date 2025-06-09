package criteria

type CreateUser struct {
	UserName string `json:"user_name"`
}

type FetchUser struct {
	UserName string `json:"user_name"`
}
