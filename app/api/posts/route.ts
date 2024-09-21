import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

export async function GET() {
  return Response.json(await prisma.post.findMany())
}

export async function POST(req: Request) {
  try {
    const { title, amount, quantity } = await req.json()
    const newPost = await prisma.post.create({
      data: {
        title,
        amount:Number(amount),
        quantity:Number(quantity)
      },
    })
    return Response.json(newPost)
  } catch (error) {
    return new Response(error as BodyInit, {
      status: 500,
    })
  }
}

