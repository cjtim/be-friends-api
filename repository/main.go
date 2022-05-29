package repository

var (
	DB *PrismaClient
)

func CreateConnection() error {
	c := NewClient()
	if err := c.Prisma.Connect(); err != nil {
		return err
	}
	DB = c
	return nil
}
