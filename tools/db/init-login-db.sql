-- CreateTable
CREATE TABLE "Users" (
	"id" TEXT NOT NULL,
	"name" TEXT NOT NULL,
	"loginMethodId" INTEGER NOT NULL,
	"createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updatedAt" TIMESTAMP(3) NOT NULL,

	CONSTRAINT "Users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "LoginMethods" (
	"id" SERIAL NOT NULL,
	"name" TEXT NOT NULL,

	CONSTRAINT "LoginMethods_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Users" ADD CONSTRAINT "Users_loginMethodId_fkey" FOREIGN KEY ("loginMethodId") REFERENCES "LoginMethods"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- CreateTable
CREATE TABLE "LineUsers" (
	"lineUid" TEXT NOT NULL,
	"name" TEXT NOT NULL,
	"profilePic" TEXT,
	"createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updatedAt" TIMESTAMP(3) NOT NULL,
	"userId" TEXT NOT NULL,

	CONSTRAINT "LineUsers_pkey" PRIMARY KEY ("lineUid")
);

-- CreateIndex
CREATE UNIQUE INDEX "LineUsers_lineUid_key" ON "LineUsers"("lineUid");

-- CreateIndex
CREATE UNIQUE INDEX "LineUsers_userId_key" ON "LineUsers"("userId");

-- AddForeignKey
ALTER TABLE "LineUsers" ADD CONSTRAINT "LineUsers_userId_fkey" FOREIGN KEY ("userId") REFERENCES "Users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- CreateTable
CREATE TABLE "JwtUsers" (
    "id" SERIAL NOT NULL,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "JwtUsers_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "JwtUsers" ADD CONSTRAINT "JwtUsers_userId_fkey" FOREIGN KEY ("userId") REFERENCES "Users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
