/*
  Warnings:

  - Added the required column `usersId` to the `LineUsers` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_LineUsers" (
    "lineUid" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "profilePic" TEXT,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "usersId" TEXT NOT NULL,
    CONSTRAINT "LineUsers_usersId_fkey" FOREIGN KEY ("usersId") REFERENCES "Users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
INSERT INTO "new_LineUsers" ("createdAt", "lineUid", "name", "profilePic", "updatedAt") SELECT "createdAt", "lineUid", "name", "profilePic", "updatedAt" FROM "LineUsers";
DROP TABLE "LineUsers";
ALTER TABLE "new_LineUsers" RENAME TO "LineUsers";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
