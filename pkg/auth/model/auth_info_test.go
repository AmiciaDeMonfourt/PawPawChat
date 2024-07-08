package model_test

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ti := time.Now()
	ti = ti.UTC()
	fmt.Println(ti.Format(time.RFC3339))
}

// func TestUserAuth_UnmarshalJSON(t *testing.T) {
// 	credentials := &model.AuthInfo{
// 		UserID:    1,
// 		Email:     "mail@com",
// 		HashPass:  "121",
// 		LastLogin: time.Now(),
// 	}

// 	json, err := credentials.MarshalJSON()
// 	assert.NoError(t, err)

// 	fmt.Println(string(json))

// 	var credentialsCopy model.AuthInfo
// 	assert.NoError(t, credentialsCopy.UnmarshalJSON(json))

// 	assert.Equal(t, credentials.UserID, credentialsCopy.UserID)
// 	assert.Equal(t, credentials.Email, credentialsCopy.Email)
// 	assert.Equal(t, credentials.HashPass, credentialsCopy.HashPass)
// 	assert.True(t, credentials.LastLogin.Equal(credentialsCopy.LastLogin))

// }
