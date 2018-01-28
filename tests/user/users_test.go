package user_test

import (
	"strconv"
	"testing"
	"time"

	"gitlab.com/iH8c0ff33/cosmicbox-api-server/model"
	. "gitlab.com/iH8c0ff33/cosmicbox-api-server/store/datastore"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db *Datastore
var testUsers []*model.User

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Model")
}

func populateTestUsers() {
	By("populating test users")
	testUsers = []*model.User{}
	for i := 0; i < 5; i++ {
		a := strconv.Itoa(i)
		testUsers = append(testUsers, &model.User{
			Email:   "foobar" + a + "@example.com",
			Expiry:  time.Now().Add(time.Duration(i) * time.Second),
			Hash:    "foo.bar." + a,
			Login:   "foo_bar" + a,
			Refresh: "foo.Refresh.bar." + a,
			Token:   "some.serious.stuff." + a,
		})
	}
}

func addTestUsersToDb() {
	populateTestUsers()
	By("adding test users to the db")
	for _, user := range testUsers {
		db.CreateUser(user)
	}
}

var _ = Describe("User", func() {
	BeforeSuite(func() {
		db = NewTestDb()
	})

	BeforeEach(func() {
		By("deleting \"users\" table")
		db.Exec("DELETE FROM users")
	})

	Context("given a list of users", func() {
		BeforeEach(populateTestUsers)

		It("should add them to the database", func() {
			for _, user := range testUsers {
				err := db.CreateUser(user)
				Expect(err).NotTo(HaveOccurred())
				Expect(user.ID).NotTo(Equal(0))
			}
		})
	})

	Context("given a list of users in the database", func() {
		BeforeEach(addTestUsersToDb)

		It("should get them by id", func() {
			for _, user := range testUsers {
				gotUser, err := db.GetUser(user.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(gotUser.ID).To(Equal(user.ID))
				Expect(gotUser.Login).To(Equal(user.Login))
				Expect(gotUser.Token).To(Equal(user.Token))
				Expect(gotUser.Expiry.Equal(user.Expiry)).To(BeTrue())
				Expect(gotUser.Refresh).To(Equal(user.Refresh))
				Expect(gotUser.Email).To(Equal(user.Email))
				Expect(gotUser.Hash).To(Equal(user.Hash))
			}
		})

		It("should get the list of them", func() {
			users, err := db.GetAllUsers()
			Expect(err).NotTo(HaveOccurred())
			for index, user := range users {
				Expect(user.ID).To(Equal(testUsers[index].ID))
				Expect(user.Login).To(Equal(testUsers[index].Login))
				Expect(user.Token).To(Equal(testUsers[index].Token))
				Expect(user.Expiry.Equal(testUsers[index].Expiry)).To(BeTrue())
				Expect(user.Refresh).To(Equal(testUsers[index].Refresh))
				Expect(user.Email).To(Equal(testUsers[index].Email))
				Expect(user.Hash).To(Equal(testUsers[index].Hash))
			}
		})

		It("should get the count of users", func() {
			count, err := db.GetUsersCount()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(len(testUsers)))
		})

		It("should update the users' text", func() {
			for _, user := range testUsers {
				err := db.UpdateUser(&model.User{
					ID:    user.ID,
					Login: user.Login + "abc",
				})
				updatedTodo, _ := db.GetUser(user.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(updatedTodo.Login).To(Equal(user.Login + "abc"))
			}
		})

		It("should delete all the users", func() {
			for _, user := range testUsers {
				err := db.DeleteUser(user)
				_, err1 := db.GetUser(user.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(err1).To(HaveOccurred())
			}
		})
	})
})
