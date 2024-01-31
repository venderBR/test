package unit

import(
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/venderBR/test/entity"
)

func TestAll(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_all`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
func TestFirstname(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Firstname is required"))
	})
}


func TestLastname(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Lastname is required"))
	})
}

func TestId(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})
	t.Run(`test_format`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is invalid"))
	})
}

func TestPhone(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))
	})
	t.Run(`test_format`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is invalid"))
	})
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})
	t.Run(`test_format`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor",
			LinkedIn: "https://mangpor.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}

func TestURL(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`test_required`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("LinkedIn is required"))
	})
	t.Run(`test_format`, func(t *testing.T) {
		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "064572xxxx",
			Email: "mangpor@gmail.com",
			LinkedIn: "mangpor",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Url LinkedIn is invalid"))
	})
}
