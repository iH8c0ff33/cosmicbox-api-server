package datastore

import (
	"testing"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/model"
	"github.com/franela/goblin"
)

func TestUsers(t *testing.T) {
	s := newTest()
	defer s.Close()

	g := goblin.Goblin(t)
	g.Describe("User", func() {
		// delete every table
		g.BeforeEach(func() {
			s.Exec("DELETE FROM users")
			s.Exec("DELETE FROM events")
		})

		g.It("should update a user", func() {
			user := model.User{
				Login: "username",
				Email: "email@email.it",
				Token: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
			}
			err1 := s.CreateUser(&user)
			err2 := s.UpdateUser(&user)
			getuser, err3 := s.GetUser(user.ID)
			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsTrue()
			g.Assert(err3 == nil).IsTrue()
			g.Assert(user.ID).Equal(getuser.ID)
		})

		g.It("should add a new user", func() {
			user := model.User{
				Login: "random",
				Token: "a9d285f2e8b61836ea6eb025991c12fba20de001",
				Email: "e@mail.it",
			}
			err := s.CreateUser(&user)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID != 0).IsTrue()
		})

		g.It("should get a user", func() {
			user := model.User{
				Login:  "loginname",
				Token:  "a9d285f2e8b61836ea6eb025991c12fba20de001",
				Secret: "f634c25ee39995676855ab60bd66b721d17767cf",
				Email:  "e@mail.it",
			}
			s.CreateUser(&user)
			getuser, err := s.GetUser(user.ID)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID).Equal(getuser.ID)
			g.Assert(user.Login).Equal(getuser.Login)
			g.Assert(user.Token).Equal(getuser.Token)
			g.Assert(user.Secret).Equal(getuser.Secret)
			g.Assert(user.Email).Equal(getuser.Email)
		})

		g.It("should get a user by login name", func() {
			user := model.User{
				Login: "qwerty123",
				Token: "9e0845c7abf7942d001f5694604a0f5abedfbe92",
				Email: "querty@gmail.com",
			}
			s.CreateUser(&user)
			getuser, err := s.GetUserByLogin(user.Login)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID).Equal(getuser.ID)
			g.Assert(user.Login).Equal(getuser.Login)
		})

		g.It("should not create more users with same login name", func() {
			user1 := model.User{
				Login: "random",
				Token: "a9d285f2e8b61836ea6eb025991c12fba20de001",
				Email: "e@mail.it",
			}
			user2 := model.User{
				Login: "random",
				Token: "f363630e269950a7e407c1686925b78b01583278",
				Email: "e@mail.it",
			}

			err1 := s.CreateUser(&user1)
			err2 := s.CreateUser(&user2)
			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsFalse()
		})

		g.It("should get the user list", func() {
			user1 := model.User{
				Login: "jane",
				Email: "foo@bar.com",
				Token: "ab20g0ddaf012c744e136da16aa21ad9",
			}
			user2 := model.User{
				Login: "joe",
				Email: "foo@bar.com",
				Token: "e42080dddf012c718e476da161d21ad5",
			}
			s.CreateUser(&user1)
			s.CreateUser(&user2)
			users, err := s.GetUserList()
			g.Assert(err == nil).IsTrue()
			g.Assert(len(users)).Equal(2)
			g.Assert(users[0].Login).Equal(user1.Login)
			g.Assert(users[0].Email).Equal(user1.Email)
			g.Assert(users[0].Token).Equal(user1.Token)
		})

		g.It("should get the user count", func() {
			user1 := model.User{
				Login: "jane",
				Email: "foo@bar.com",
				Token: "ab20g0ddaf012c744e136da16aa21ad9",
			}
			user2 := model.User{
				Login: "joe",
				Email: "foo@bar.com",
				Token: "e42080dddf012c718e476da161d21ad5",
			}
			s.CreateUser(&user1)
			s.CreateUser(&user2)
			count, err := s.GetUserCount()
			g.Assert(err == nil).IsTrue()
			if s.driver != "postgres" {
				// we have to skip this check for postgres because it uses
				// an estimate which may not be updated.
				g.Assert(count).Equal(2)
			}
		})

		g.It("should get 0 as user count", func() {
			count, err := s.GetUserCount()
			g.Assert(err == nil).IsTrue()
			g.Assert(count).Equal(0)
		})

		g.It("should delete a user", func() {
			user := model.User{
				Login: "joe",
				Email: "foo@bar.com",
				Token: "e42080dddf012c718e476da161d21ad5",
			}
			s.CreateUser(&user)
			_, err1 := s.GetUser(user.ID)
			err2 := s.DeleteUser(&user)
			_, err3 := s.GetUser(user.ID)
			g.Assert(err1 == nil).IsTrue()
			g.Assert(err2 == nil).IsTrue()
			g.Assert(err3 == nil).IsFalse()
		})
	})
}
