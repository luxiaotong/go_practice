package main

import (
	"testing"

	"github.com/gavv/httpexpect"
)

var e *httpexpect.Expect

func init() {
}

func TestAll(t *testing.T) {
	e = httpexpect.New(t, "http://127.0.0.1:8096")
	// t.Run("testSendCode", testSendCode)
	// t.Run("testSignUp", testSignUp)
	// defer clearUser()
	t.Run("testSignIn", testSignIn)
	// t.Run("testUserPass", testUserPass)
	// t.Run("testUploadFile", testUploadFile)
	// t.Run("testUpdateUser", testUpdateUser)
	// t.Run("testUpdateUser_Logo", testUpdateUser_Logo)
	// t.Run("testCertifyUser", testCertifyUser)
	t.Run("testGetUser", testGetUser)

	t.Run("testAdminSignIn", testAdminSignIn)
	t.Run("testGetUsers", testGetUsers)
	// t.Run("testAuditUser", testAuditUser)
	// t.Run("testAdminUpdateUser", testAdminUpdateUser)
	// t.Run("testFreezeUser", testFreezeUser)
	// t.Run("testAdminGetUser", testAdminGetUser)
	// t.Run("testGetUsers", testGetUsers)

	t.Run("testAddTag", testAddTag)
	t.Run("testGetTags", testGetTags)
	t.Run("testOpTag", testOpTag)
	defer clearTags()

	// t.Run("testUploadFile_Definition", testUploadFile_Definition)
	// t.Run("testGetDictAttach_Definition", testGetDictAttach_Definition)
	// t.Run("testUploadFile_Vote", testUploadFile_Vote)
	// t.Run("testGetDictAttach_Vote", testGetDictAttach_Vote)
	t.Run("testAddDict", testAddDict)
	t.Run("testGetDicts", testGetDicts)
	t.Run("testGetDict", testGetDict)
	t.Run("testGetFields", testGetFields)
	t.Run("testOpDict", testOpDict)
	t.Run("testEditDict", testEditDict)
	t.Run("testAuditDict", testAuditDict)
	t.Run("testSearchDicts", testSearchDicts)
	// t.Run("testDeleteDict", testDeleteDict)
}
