/*
  Warnings:

  - A unique constraint covering the columns `[lineUid]` on the table `LineUsers` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[userId]` on the table `LineUsers` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "LineUsers_lineUid_key" ON "LineUsers"("lineUid");

-- CreateIndex
CREATE UNIQUE INDEX "LineUsers_userId_key" ON "LineUsers"("userId");
