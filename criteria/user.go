package criteria

type CreateUser struct {
	Email string `json:"email"`
}

type FetchUser struct {
	Email string `json:"email"`
}
