/*
  Warnings:

  - You are about to drop the column `usersId` on the `LineUsers` table. All the data in the column will be lost.
  - Added the required column `userId` to the `LineUsers` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_LineUsers" (
    "lineUid" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "profilePic" TEXT,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "userId" TEXT NOT NULL,
    CONSTRAINT "LineUsers_userId_fkey" FOREIGN KEY ("userId") REFERENCES "Users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_LineUsers" ("createdAt", "lineUid", "name", "profilePic", "updatedAt") SELECT "createdAt", "lineUid", "name", "profilePic", "updatedAt" FROM "LineUsers";
DROP TABLE "LineUsers";
ALTER TABLE "new_LineUsers" RENAME TO "LineUsers";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
