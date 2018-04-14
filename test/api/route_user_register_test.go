package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var registerTestCasses = []apiTestCase{
	apiTestCase{
		requests: []func(*testing.T, []response) request{
			func(t *testing.T, r []response) request {
				return request{
					method: "POST",
					path:   "/auth/register",
					body: map[string]interface{}{
						"username": "username",
						"password": "password",
						"email":    "email",
					},
				}
			},
		},
		validate: func(t *testing.T, responses []response, session *mgo.Session) {
			r := responses[0]
			t.Logf("%+v", r)

			body, bodyOk := r.body.(map[string]interface{})
			require.True(t, bodyOk)
			assert.Equal(t, http.StatusOK, r.code)

			assert.Equal(t, "username", body["username"])
			assert.Equal(t, "email", body["email"])
			assertMongoID(t, body["id"])

			var result map[string]interface{}
			if err := session.DB("").C("user").Find(bson.M{"username": "username"}).One(&result); err != nil {
				t.Fatal(err)
			}

			if result["username"] != "username" || result["email"] != "email" {
				t.Fatalf("wrong user data in db %v", result)
			}
		},
	},
}

func TestUserRegister(t *testing.T) {
	runTestCases(t, registerTestCasses)
}
