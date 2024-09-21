'use client'

import React, { useEffect, useState } from 'react'
import axios from 'axios'


import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useRouter } from 'next/navigation'

const Edit = ({ params }: { params: { id: string } }) => {
    const [title, setTitle] = useState('')
    const [amount, setAmount] = useState('')
    const [quantity, setQuantity] = useState('')
    const router = useRouter()
    const { id } = params

    const fetchPost = async (id: Number) => {
        try {
            const res = await axios.get(`/api/posts/${id}`)
            setTitle(res.data.title)
            setAmount(res.data.amount)
            setQuantity(res.data.quantity)
        } catch (error) {
            console.error(error)
        }
    }

    useEffect(() => {
        if (id) {
            fetchPost(parseInt(id))
        }
    }, [id])

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()

        try {
            await axios.put(`/api/posts/${id}`, {
                title,
                amount:Number(amount),
                quantity:Number(quantity)
            })
            router.push('/')
        } catch (error) {
            console.error(error)
        }
    }

    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Edit Requirement {id}</CardTitle>

            </CardHeader>
            <CardContent>
                <form>
                    <div className="grid w-full items-center gap-4">
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="title">Title</Label>
                            <Input
                                id="title"
                                placeholder="Name of your budget request"
                                type="text"
                                name="title"
                                required
                                value={title}
                                onChange={(e) => setTitle(e.target.value)} />
                        </div>
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="amount">Amount</Label>
                            <Input
                                id="amount"
                                type="text"
                                name="amount"
                                required
                                value={amount}
                                onChange={(e) => setAmount(e.target.value)} />
                        </div>
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="quantity">Quantity</Label>
                            <Input
                                id="quantity"
                                type="text"
                                name="quantity"
                                required
                                value={quantity}
                                onChange={(e) => setQuantity(e.target.value)} />
                        </div>

                    </div>
                </form>
            </CardContent>
            <CardFooter className="flex justify-between">
                <Button variant="outline">Cancel</Button>
                <Button onClick={handleSubmit}>Edit</Button>
            </CardFooter>
        </Card>
    )
}

export default Edit