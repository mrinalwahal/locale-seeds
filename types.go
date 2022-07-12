package main

type (

	// Object struct
	Object struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// Role struct
	Role struct {
		ID          string       `json:"id"`
		Name        string       `json:"name"`
		Permissions []Permission `json:"permissions"`
	}

	// Operations struct
	Operation struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}

	// Permission struct
	Permission struct {
		ID          string    `json:"id"`
		ObjectID    string    `json:"object_id"`
		Object      Object    `json:"object"`
		Operation   Operation `json:"operation"`
		OperationID string    `json:"operation_id"`
	}
)
