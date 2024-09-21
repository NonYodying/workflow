/*
  Warnings:

  - You are about to drop the column `content` on the `Post` table. All the data in the column will be lost.

*/
-- CreateEnum
CREATE TYPE "Status" AS ENUM ('PENDING', 'APPROVED', 'REJECTED');

-- AlterTable
ALTER TABLE "Post" DROP COLUMN "content",
ADD COLUMN     "amount" INTEGER,
ADD COLUMN     "quantity" INTEGER,
ADD COLUMN     "status" "Status" NOT NULL DEFAULT 'PENDING';
