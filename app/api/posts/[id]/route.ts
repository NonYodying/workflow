import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

export async function GET(
  req: Request,
  { params }: { params: { id: string } },
) {
  return Response.json(await prisma.post.findUnique({
    where: { id: Number(params.id) },
  }))
}

export async function PUT(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    const { title, amount, quantity } = await req.json()
    return Response.json(await prisma.post.update({
      where: { id: Number(params.id) },
      data: { 
        title,
        amount:Number(amount),
        quantity:Number(quantity) },
    }))
  } catch (error) {
    return new Response(error as BodyInit, {
      status: 500,
    })
  }
}

export async function PATCH(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    const { status } = await req.json()
    return Response.json(await prisma.post.update({
      where: { id: Number(params.id) },
      data: { 
       status },
    }))
  } catch (error) {
    return new Response(error as BodyInit, {
      status: 500,
    })
  }
}

export async function DELETE(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    return Response.json(await prisma.post.delete({
      where: { id: Number(params.id) },
    }))
  } catch (error) {
    return new Response(error as BodyInit, {
      status: 500,
    })
  }
}